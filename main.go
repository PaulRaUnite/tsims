package main

import (
	sim "tsims/simulator"
	"fmt"
	"log"
)

// addOp is a utility function for adding operations to program code
func addOp(p sim.Program, nowState uint64, inputSymbol, memorySymbol,
inputMoveSymbol, memoryMoveSymbol, writeSymbol rune, next uint64) {
	var moveIn sim.StillRight
	switch inputMoveSymbol {
	case 'S':
		moveIn = sim.InStill
	case 'R':
		moveIn = sim.InRight
	default:
		panic(
			"Input tape's head can be moved to right or stay still.\n" +
				"Use S to stay still or R to move to right.",
		)
	}

	var moveMem sim.LeftStillRight
	switch memoryMoveSymbol {
	case 'L':
		moveMem = sim.MemLeft
	case 'S':
		moveMem = sim.MemStill
	case 'R':
		moveMem = sim.MemRight
	default:
		panic(
			"Memory tape's head can be moved to right, to left or stay still.\n" +
				"Use S to stay still, R to move to right or L to move to left.",
		)
	}
	p.AddOperation(
		sim.Snapshot{
			InputSymbol:  inputSymbol,
			MemorySymbol: memorySymbol,
			State:        nowState,
		},
		sim.Operation{
			MoveInput: moveIn,
			MoveMem:   moveMem,
			Symbol:    writeSymbol,
			State:     next,
		},
	)
}

func main() {
	p := sim.NewProgram('#')
	p.AddFinalState(0, "equal")
	p.AddFinalState(1, "1 more")
	p.AddFinalState(3, "0 more")

	// Program code that finds what symbols are more: 0 or 1 or their amount is equal
	// See documentation for more info and state diagram.
	addOp(p, 0, '1', '#', 'R', 'S', '#', 1)
	addOp(p, 1, '1', '#', 'R', 'R', '1', 1)
	addOp(p, 1, '0', '#', 'S', 'L', '#', 2)
	addOp(p, 2, '0', '1', 'R', 'S', '#', 1)
	addOp(p, 2, '0', '#', 'R', 'S', '#', 0)

	addOp(p, 0, '0', '#', 'R', 'S', '#', 3)
	addOp(p, 3, '0', '#', 'R', 'R', '0', 3)
	addOp(p, 3, '1', '#', 'S', 'L', '#', 2)
	addOp(p, 2, '1', '0', 'R', 'S', '#', 3)
	addOp(p, 2, '1', '#', 'R', 'S', '#', 0)

	// simple testing
	input := "101010010101011000110"
	expect := 0
	for _, symbol := range input {
		switch symbol {
		case '0':
			expect--
		case '1':
			expect++
		}
	}
	var expectation string
	if expect == 0 {
		expectation = "equal"
	} else if expect > 0 {
		expectation = "1 more"
	} else {
		expectation = "0 more"
	}
	result, err := sim.Simulate(p, input)
	if err != nil {
		log.Fatalln(err)
	}
	description, err := p.Interpret(result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(description, expectation)
}
