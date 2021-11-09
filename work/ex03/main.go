package main

import (
	"fmt"
	"gopher-dojo/work/ex03/typing"
	"os"
	"time"
)

func main() {
	var cnt int = 0
	fmt.Println("Start!!!")
	ch := typing.Input(os.Stdin)
	for {
		fmt.Print("test->")
		select {
		case res := <-ch:
			chk := typing.CheckInput("test", res)
			if chk {
				cnt++
			}
		case <-time.After(5 * time.Second):
			fmt.Println()
			fmt.Println("Finish!!")
			fmt.Printf("Your Score: %v\n", cnt)
			return
		}
	}
}
