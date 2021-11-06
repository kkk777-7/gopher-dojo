package main

import (
	"flag"
	"gopher-dojo/work/ex01/convert"
	"log"
)

var PreFormat string
var AfterFormat string

func init() {
	flag.StringVar(&PreFormat, "p", "jpg", "pre image format")
	flag.StringVar(&AfterFormat, "a", "png", "after image format")
	flag.Parse()
}

func main() {
	i := convert.CreateImageFormat(flag.Args()[0], PreFormat, AfterFormat)
	err := i.Convert()
	if err != nil {
		log.Fatal(err)
	}
}
