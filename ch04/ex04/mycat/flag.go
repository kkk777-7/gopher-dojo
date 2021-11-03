package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var op bool
var n int = 1
var num int = 1

func init() {
	flag.BoolVar(&op, "n", false, "Show the number of lines")
}

func main() {
	flag.Parse()
	if op {
		num = 2
	}
	files := os.Args[num:]
	for _, fn := range files {
		readFile(fn, op)
	}
}

func readFile(fn string, op bool) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if !op {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer f.Close()
	} else {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(n, ":", scanner.Text())
			n++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer f.Close()
	}
}
