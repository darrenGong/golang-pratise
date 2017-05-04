package main

import (
	"net/http"
	"fmt"
	"hello/web"
	"github.com/elazarl/go-bindata-assetfs"
)

/*
 * @ track all request and replay it easily
 */

func main() {
	http.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to the home page ! %s", "hello web")
	}))

	http.Handle("/", http.FileServer(
    &assetfs.AssetFS{Asset: web.Asset, AssetDir: web.AssetDir, AssetInfo: web.AssetInfo, Prefix: "data"}))


	http.ListenAndServe(":8080", nil)

	//server := httptest.NewServer(nil)
	//addr, debugAddr := ":8080", ":8089"
	//boast.Serve(server, addr, debugAddr)
}
