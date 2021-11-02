package main

import "fmt"

func Swap(n int, m int) (n2, m2 int) {
	n2, m2 = m, n
	return n2, m2
}

func main() {
	chk1, chk2 := Swap(10, 20)
	fmt.Println(chk1, chk2)
}
