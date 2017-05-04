package main

import (
	"path/filepath"
	"fmt"
)

var (
	kRoot = "F:\\go-dev\\src\\cnat_manager"
	kScript = "scripts\\templates"
	kVersion = "v0.2.0-igw"
)

func main() {
	tmplPath := filepath.Join(kRoot, kScript)
	fmt.Println(tmplPath)
	absolutePath := filepath.Join(tmplPath, kVersion)
	fmt.Println(absolutePath)

	pattern := filepath.Join(absolutePath, "*.tmpl")
	fmt.Println(pattern)
}
