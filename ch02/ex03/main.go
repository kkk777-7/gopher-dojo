package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 101; i++ {
		if i%2 == 1 {
			fmt.Printf("%d-奇数\n", i)
		} else {
			fmt.Printf("%d-偶数\n", i)
		}
	}
}
