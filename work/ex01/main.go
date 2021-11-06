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
	flag.StringVar(&PreFormat, "p", "jpg", "pre image format")
	flag.StringVar(&AfterFormat, "a", "png", "after image format")
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		log.Fatal("Please input the target path!")
	}

	chk := checkFormat(flag.Args()[0], PreFormat, AfterFormat)
	if chk {
		i := convert.CreateImageFormat(flag.Args()[0], PreFormat, AfterFormat)
		err := i.Convert()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func checkFormat(path string, pre string, after string) bool {
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
