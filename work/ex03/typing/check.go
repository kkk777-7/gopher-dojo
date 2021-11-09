package typing

import "fmt"

func CheckInput(w1 string, w2 string) bool {
	var chk bool
	if w1 == w2 {
		fmt.Println("Nice!")
		chk = true
	} else {
		fmt.Println("Bad!")
		chk = false
	}
	return chk
}
