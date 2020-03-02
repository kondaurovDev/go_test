package main

import (
	"fmt"

	"./functions"
)

func main() {
	fmt.Printf("Hello!!!\n")
	functions.SayHello("arr")
	functions.StartHttp()
}
