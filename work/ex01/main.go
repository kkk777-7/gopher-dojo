package main

import (
	"flag"
	"fmt"
	"gopher-dojo/work/ex01/convert"
	"log"
)

var PreFormat string
var AfterFormat string

func init() {
	flag.StringVar(&PreFormat, "pre", "jpeg", "pre image format")
	flag.StringVar(&AfterFormat, "after", "png", "after image format")
	flag.Parse()
}

func main() {
	fmt.Println(PreFormat, AfterFormat, flag.Args()[0], flag.Args()[1])
	err := convert.Convert(flag.Args()[0], flag.Args()[1], PreFormat, AfterFormat)
	if err != nil {
		log.Fatal(err)
	}
}
