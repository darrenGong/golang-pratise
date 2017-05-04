package main

import (
	"net/http"
	"context"
	"time"
	"fmt"
	"common-lib/userip"
	"encoding/json"
	"html/template"
	"log"
)

func main() {
	http.HandleFunc("/search", handleSearch)

	log.Fatal(http.ListenAndServe(":8089", nil))
}

var resultTemplate = template.Must(template.New("result").Parse(`
<html>
<head/>
<body>
	<ol>
	{ <br/>
	<p>Recode: {{.Result.RetCode}}</p>
	<p>Message: {{.Result.Message}}</p>
	} <br/>
	</ol>

	<p>result in {{.Elapsed}}, timeout {{.Timeout}} </p>
</body>
</html>
`))

func handleSearch(w http.ResponseWriter, req *http.Request) {
	var (
		ctx context.Context
		cancel context.CancelFunc
	)
	log.Printf("Request: %v", req.RequestURI)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	start := time.Now()
	result, err := search(ctx, "13042")
	elapsed := time.Since(start)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resultTemplate.Execute(w, struct {
		Result			 	Result
		Timeout, Elapsed	time.Duration
	}{
		Result: 	*result,
		Timeout: 	timeout,
		Elapsed: 	elapsed,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
}

type Result struct {
	RetCode int
	Message string
}

func search(ctx context.Context, query string) (*Result, error) {
	req, err := http.NewRequest("GET", "http://172.17.144.42:16001/", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to NewRequest[%v]", err)
	}

	q := req.URL.Query()
	q.Set("Action", query)

	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("OrganizationIds[]", string(userIP))
	}
	req.URL.RawQuery = q.Encode()

	result := Result{}
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("Failed to decode response body[err: %v]", err)
		}

		return nil
	})

	return &result, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	c := make(chan error, 1)
	go func() {
		c <- f(client.Do(req))
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c

		return ctx.Err()
	case err := <-c:
		return err
	}
}
