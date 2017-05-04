package main

import (
	"flag"
	"fmt"
)

type Interface struct {
	Data interface{}
}

var (
	KeyData = make(map[string]interface{})
)

var (
	config = flag.String("c", "./config/config.json", "application config file, json format")
)

func main() {
	flag.Parse()
	fmt.Println(*config)

	KeyData["hello"] = "world"
	interfaceData := Interface{Data: KeyData}
	data, ok := interfaceData.Data.(map[string]interface{})
	if ok {
		fmt.Println(data)
	}
}