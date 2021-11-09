package main

import (
	"flag"
	"gopher-dojo/work/ex03/typing"
	"time"
)

func main() {
	n := flag.Duration("t", 30*time.Second, "Time of typing game")
	flag.Parse()
	typing.Typing(n)

}
