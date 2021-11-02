package main

import "fmt"

func Check(i int) {
	if i%2 == 0 {
		fmt.Printf("%d-偶数\n", i)
	} else {
		fmt.Printf("%d-奇数\n", i)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		Check(i)
	}
}
