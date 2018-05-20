package simulator

import (
	"tsims/tape"
	"errors"
	"fmt"
)

type StillRight bool

const (
	InStill StillRight = true
	InRight StillRight = false
)

type LeftStillRight int8

const (
	MemLeft  LeftStillRight = iota
	MemStill
	MemRight
)

type Operation struct {
	MoveInput StillRight
	MoveMem   LeftStillRight
	Symbol    rune
	State     uint64
}

type Program struct {
	code        map[uint64]map[rune]Operation
	emptySymbol rune
	finalStates map[uint64]string
}

func NewProgram(emptySymbol rune) Program {
	return Program{
		code:        make(map[uint64]map[rune]Operation),
		emptySymbol: emptySymbol,
		finalStates: make(map[uint64]string),
	}
}

func (p Program) AddOperation(state uint64, char rune, operation Operation) {
	operations, ok := p.code[state]
	if !ok {
		operations = make(map[rune]Operation)
		p.code[state] = operations
	}
	operations[char] = operation
}

func (p Program) AddFinalState(state uint64, description string) {
	p.finalStates[state] = description
}

func Interpret(program Program, input string) (description string, err error) {
	inputTape := tape.Create(input, program.emptySymbol)
	memory := tape.Create(string(program.emptySymbol), program.emptySymbol)
	state := uint64(0)
	for {
		char := inputTape.Now()
		if char == program.emptySymbol {
			if desc, ok := program.finalStates[state]; ok {
				return desc, nil
			} else {
				return "", errors.New(fmt.Sprintf("Interpreting error in state %v: the state is not final.\n%s\n", state, memory.String()))
			}
		}
		if operations, ok := program.code[state]; ok {
			if op, ok := operations[char]; ok {
				memory.Set(op.Symbol)
				switch op.MoveInput {
				case InStill:
				case InRight:
					inputTape.RightShift()
				}
				switch op.MoveMem {
				case MemLeft:
					memory.LeftShift()
				case MemStill:
				case MemRight:
					memory.RightShift()
				}
				state = op.State
			} else {
				return "", errors.New(fmt.Sprintf("Interpreting error in state %v: there is no operations for the character `%s`.\n%s\n", string(char), state, memory.String()))
			}
		} else {
			return "", errors.New(fmt.Sprintf("Interpreting error in state %v: there is no operations for the state.\n%s\n", state, memory.String()))
		}
	}
}
