package main

import (
	"tsims/simulator"
	"fmt"
)

func main() {
	p := simulator.NewProgram('#')
	p.AddOperation(0, '1', simulator.Operation{
		MoveInput: simulator.InStill,
		MoveMem:   simulator.MemStill,
		Symbol:    '1',
		State:     1,
	})
	fmt.Println(simulator.Interpret(p, "1"))
}
