package main

import "fmt"

func main() {
	var sum int = 5 + 6 + 3
	avg := float64(sum) / 3
	if avg > 4.5 {
		fmt.Println("good")
	}
}
