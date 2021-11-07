package main

import (
	"flag"
	"fmt"
	"gopher-dojo/work/ex02/convert"
	"log"
)

var SrcPath string
var PreFormat string
var AfterFormat string

func init() {
	flag.StringVar(&SrcPath, "path", ".", "the source directory")
	flag.StringVar(&PreFormat, "p", "jpg", "pre image format")
	flag.StringVar(&AfterFormat, "a", "png", "after image format")
}

func main() {

	flag.Parse()
	chk := CheckFormat(SrcPath, PreFormat, AfterFormat)
	if chk {
		i := convert.CreateImageFormat(SrcPath, PreFormat, AfterFormat)
		err := i.Convert()
		if err != nil {
			fmt.Printf("%T, %v\n", err, &err)
			log.Fatal(err)
		}
	}
}

func CheckFormat(path string, pre string, after string) bool {
	var prechk bool = false
	var afterchk bool = false
	var chk bool = true

	extensions := []string{"png", "jpg", "jpeg", "gif"}
	for _, extension := range extensions {
		if PreFormat == extension {
			prechk = true
		}
		if AfterFormat == extension {
			afterchk = true
		}
	}
	if !prechk {
		chk = false
		fmt.Println("The pre extension is incorrect")
	}
	if !afterchk {
		chk = false
		fmt.Println("The after extension is incorrect")
	}
	return chk
}
