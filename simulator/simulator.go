package simulator

import (
	"tsims/tape"
	"errors"
	"fmt"
)

// "Enum" for input tape move control.
type StillRight bool

// "Enum"'s alias values.
const (
	InStill StillRight = true
	InRight StillRight = false
)

// "Enum" for memory tape move control.
type LeftStillRight int8

// "Enum"'s values.
const (
	MemLeft  LeftStillRight = iota
	MemStill
	MemRight
)

// Operation struct provides
// changing of the machine state.
type Operation struct {
	MoveInput StillRight
	MoveMem   LeftStillRight
	Symbol    rune
	State     uint64
}

// Snapshot struct provides information
// about current state of the machine.
type Snapshot struct {
	InputSymbol  rune
	MemorySymbol rune
	State        uint64
}

type Program struct {
	code        map[Snapshot]Operation
	emptySymbol rune
	finalStates map[uint64]string
}

// Program constructor.
func NewProgram(emptySymbol rune) Program {
	return Program{
		code:        make(map[Snapshot]Operation),
		emptySymbol: emptySymbol,
		finalStates: make(map[uint64]string),
	}
}

// Set what operation will be performed
// when the machine will be in the snapshot state.
func (p Program) AddOperation(snapshot Snapshot, operation Operation) {
	p.code[snapshot] = operation
}

// Adds finals states with their descriptions.
func (p Program) AddFinalState(state uint64, description string) {
	p.finalStates[state] = description
}

// Returns interpretation of the state or error if it is not.
func (p Program) Interpret(state uint64) (string, error) {
	if desc, ok := p.finalStates[state]; ok {
		return desc, nil
	} else {
		return "", errors.New("there is no such final state")
	}
}

const DEBUG = false

// Simulate program with the input tape.
// Returns final state in which machine would stopped or error.
func Simulate(program Program, input string) (state uint64, err error) {
	// Tapes initialization.
	inputTape := tape.Create(input, program.emptySymbol)
	memory := tape.Create("", program.emptySymbol)
	state = uint64(0)

	for {
		// For debug purposes.
		if (DEBUG) {
			fmt.Println("state", state)
			fmt.Println(inputTape)
			fmt.Println(memory)
		}
		char := inputTape.HeadSymbol()
		// If input tape reaches empty symbol,
		// it means that there is no input data,
		// so the machine must end its work in the moment
		// and have answer.
		if char == program.emptySymbol {
			if _, ok := program.finalStates[state]; ok {
				// Program have done its work perfect.
				return state, nil
			} else {
				// Something wrong, check your program.
				return 0, errors.New(
					fmt.Sprintf(
						"Interpreting error in state %v: the state is not final.\n"+
							"Input tape:\n%s\n"+
							"Memory tape:\n%s\n",
						state, inputTape.String(), memory.String()))
			}
		}

		// Snapshot construction.
		snapshot := Snapshot{
			InputSymbol:  char,
			MemorySymbol: memory.HeadSymbol(),
			State:        state,
		}
		// Try to get operation to perform.
		if op, ok := program.code[snapshot]; ok {
			// Operation found, performing.
			// Writing new symbol to the memory.
			memory.Set(op.Symbol)

			// Move head of input tape.
			switch op.MoveInput {
			case InStill:
			case InRight:
				inputTape.HeadToRight()
			}
			// Move head of memory tape.
			switch op.MoveMem {
			case MemLeft:
				memory.HeadToLeft()
			case MemStill:
			case MemRight:
				memory.HeadToRight()
			}
			// Change the state to new state.
			state = op.State
		} else {
			// Operation wasn't found,
			// return error.
			return 0, errors.New(
				fmt.Sprintf("Interpreting error in state %v: there is no operation for the snapshot.\n"+
					"Input tape:\n%s\n"+
					"Memory tape:\n%s\n",
					snapshot, inputTape.String(), memory.String()))
		}
	}
}
