package main

import (
	"text/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type T struct {
	Material	string
	Count		int
	Protocol 	int32
	ArrayData	[]int32
}

var (
	kRoot = "F:\\go-dev\\src\\hello"
	kTemplates = "templates"
)

func Protocol(protoType int32) string {
	switch protoType {
	case 1:
		return "tcp"
	case 2:
		return "udp"
	case 3:
		return "gre"
	default:
		return strconv.FormatInt(int64(protoType), 10)
	}
}

func callTemplate(tmpl *template.Template, name string, data interface{}, dir string) {
	var result *os.File
	if "" == dir {
		result = os.Stdout
	} else {
		path := filepath.Join(dir, name)
		result, _ = os.Create(path)
	}
	defer result.Close()

	if err := tmpl.ExecuteTemplate(result, name, data); err != nil {
		log.Fatalf("Failed to execute tmpl[name:%s, err:%v]", name, err)
	}

	log.Printf("generate tmpl %s at %s", name, dir)
}

func main() {
	arrayData := []int32{3,4,5,6,7,8}
	sweaters := T{"wool", 100, 1, arrayData}
	funcs := template.FuncMap{
		"protocol": Protocol,
	}

	tmplPathDir := filepath.Join(kRoot, kTemplates)
	tmplPath := filepath.Join(tmplPathDir, "*.tmpl")
	tmpl, err := template.New("test").Funcs(funcs).ParseGlob(tmplPath)
	if err != nil {
		log.Panicf("Failed to parse template[err:%v]", err)
	}

	callTemplate(tmpl, "A.txt", &sweaters, tmplPathDir)
}
