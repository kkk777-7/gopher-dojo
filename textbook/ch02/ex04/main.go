package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)

	switch s {
	case 0:
		fmt.Println("凶")
	case 1, 2:
		fmt.Println("吉")
	case 3, 4:
		fmt.Println("中吉")
	case 5:
		fmt.Println("大吉")
	}
}
