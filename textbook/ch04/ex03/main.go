package main

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting"
	mygreeting "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do())
	fmt.Println(mygreeting.Do(time.Now()))
}
