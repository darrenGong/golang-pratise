package main

import (
	"regexp"
	"fmt"
)

func main() {
	b, err := regexp.MatchString("^I", "IModifyEIPBandwidth")
	if err != nil {
		fmt.Println(err)
		return
	}

	if b {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}
}
