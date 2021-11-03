package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (n I) String() string {
	return "int"
}

type S string

func (s S) String() string {
	return "string"
}

type B bool

func (b B) String() string {
	return "bool"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case S:
		fmt.Println(string(v), "S")
	case B:
		fmt.Println(bool(v), "B")
	}
}

func main() {
	var n I = I(100)
	F(n)
	F(S("hoge"))
	F(B(true))
}
