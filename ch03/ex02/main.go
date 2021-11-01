package main

import "fmt"

func main() {
	num := []int{19, 86, 1, 12}
	var sum int
	for _, i := range num {
		sum += i
	}
	fmt.Println(sum)
}
