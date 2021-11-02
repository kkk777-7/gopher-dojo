package main

import "fmt"

func main() {
	n, m := 10, 200
	swap2(&n, &m)
	fmt.Println(n, m)
}

func swap2(n *int, m *int) {
	*n, *m = *m, *n
}
