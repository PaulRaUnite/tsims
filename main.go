package main

import (
	"fmt"
	"tsims/tape"
)

func main() {
	t := tape.Create("a", '#')
	t.RightShifting()
	t.RightShifting()
	fmt.Println(t)
}