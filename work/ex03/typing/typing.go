package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func Typing(t *time.Duration) {
	var cnt int

	limit := time.After(*t)
	fmt.Println("Start!!!")
	fmt.Println("--------------")
	ch := input(os.Stdin)
	for {
		word := words()
		fmt.Print(word + "->")
		select {
		case res := <-ch:
			chk := checkInput(word, res)
			if chk {
				cnt++
			}
		case <-limit:
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

func words() string {
	words := []string{
		"kota",
		"kimura",
		"ririko",
		"nakaoka",
		"tabetai",
		"yakiniku",
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))

	word := words[i]
	return word
}
