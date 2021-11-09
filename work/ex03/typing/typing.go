package typing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func Typing(t *time.Duration) {
	var cnt int

	fmt.Println("Start!!!")
	fmt.Println("--------------")
	ch := input(os.Stdin)
	for {
		fmt.Print("test->")
		select {
		case res := <-ch:
			chk := checkInput("test", res)
			if chk {
				cnt++
			}
		case <-time.After(*t):
			fmt.Println()
			fmt.Println("--------------")
			fmt.Println("Finish!!")
			fmt.Printf("Your Score: %v\n", cnt)
			return
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func checkInput(w string, input string) bool {
	var chk bool
	if w == input {
		fmt.Println("Nice!")
		chk = true
	} else {
		fmt.Println("Bad!")
	}
	return chk
}
