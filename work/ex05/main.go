package main

import (
	"gopher-dojo/work/ex05/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":8080", nil)
}
