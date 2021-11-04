package main

import "fmt"

type Myint int

func (n *Myint) Inc() { *n++ }
func main() {
	var n Myint
	fmt.Println(n)
	n.Inc()
	fmt.Println(n)
}
