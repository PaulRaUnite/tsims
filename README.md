# Tsims
<b>T</b>uring  <b>Sim</b>ulator

## Installation

Clone repository or `go get gihub.com/PaulRaUnite/tsims` or download archive and perform `go run main.go` command in the project directory.

## The problem

Need to create a turing machine simulator and program which should answer the following question: amount of which symbols is bigger, 0 or 1 or they are equal?

## Turing machine
The machine has 2 tapes:
- input tape &mdash; read-only, its head can be moved to right or stay still;
- memory tape &mdash; read-write, its head can be moved to right or left or stay still, program can write symbol to current head position; can be increased indefinitely.

The machine stops when its head meets empty symbol in input tape; can have multiple finite states.

Program, which is performed by the Turing machine, is actually the following mapping: `(state, symbol of input head, symbol of memory head) -> (input move, memory move, symbol to write to memory, new state)`.

## Implementation
`tape` package implements tape primitive: buffer with head position, it keeps only symbols that were accessed by the head.

`simulator` package implements `Program` creation and its simulation via `Simulate`.

`main.go` defines program code and tests it.

## Program diagram

![Diagram](./diagram.svg)

Diagram legend:
- nodes are the states;
- signatures under arrows mean the following: if machine in state _where arrow starts_, input head looks at _first symbol of signature_ and memory head looks at _second symbol_, then move(`R`) or not(`S`) input head according to _third symbol_, move(`R`,`L`) or not(`S`) memory head according to _fourth symbol_ and write _fifth symbol_ in cell of current memory head position.

States:
- 0 &mdash; state of equality &mdash; amounts of 1 and 0 are the same;
- 1 &mdash; state of "1 > 0" &mdash; amount of 1 is bigger than amount of 0;
- 2 &mdash; state of choice &mdash; after convolution of 1 and 0, here is decided next state, based on memory;
- 3 &mdash; state of "1 < 0" &mdash; amount of 0 is bigger than amount of 1.

### Main principle
1. program starts in 0
3. if input is empty, program finishes with `state 0`;
2. if first symbol is 1 goto `state 1`, if 0 &mdash; `state 3`;
4. if input symbol in `state 1` is 1 &mdash; write `1` to memory;
5. if next symbol in `state 1` is 0 &mdash; move memory head to left and goto `state 2`;
6. if in `state 2` memory symbol is 1, it means that amount of 1 is bigger, so delete is and goto `state 1`;
7. the algorithm is symmetric, so for 0 it needs to replace 1 by 0 in steps `4-5`.