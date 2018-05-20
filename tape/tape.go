package tape

// Tape struct contains buffer of symbols,
// empty symbol, current position of head, and
// shift of zero position of the tape.
type Tape struct {
	buffer      []rune
	zeroPos     uint64
	currentPos  uint64
	emptySymbol rune
}

// Tape constructor.
func Create(content string, emptySymbol rune) Tape {
	if len(content) == 0 {
		content = string(emptySymbol)
	}
	return Tape{
		buffer:      []rune(content),
		zeroPos:     0,
		currentPos:  0,
		emptySymbol: emptySymbol,
	}
}

// "Enum" for enlarge to understand
// what side of buffer must be enlarged.
type side bool

// "Enum"'s aliases.
const (
	left  side = true
	right side = false
)

// Constant for buffer enlarging.
const enlargeLength uint64 = 16

// enlarge method appends array of
// empty symbols of `enlargeLength` length to
// the side of the tape buffer.
func (tape *Tape) enlarge(lr side) {
	switch lr {
	case left:
		leftBuf := make([]rune, enlargeLength)
		for i := range leftBuf {
			leftBuf[i] = tape.emptySymbol
		}
		*tape = Tape{
			buffer:      append(leftBuf, tape.buffer...),
			zeroPos:     enlargeLength + tape.zeroPos,
			currentPos:  enlargeLength + tape.currentPos,
			emptySymbol: tape.emptySymbol,
		}
	case right:
		rightBuf := make([]rune, enlargeLength)
		for i := range rightBuf {
			rightBuf[i] = tape.emptySymbol
		}
		tape.buffer = append(tape.buffer, rightBuf...)
	}
}

// Returns symbol to which the head
// of the tape are looking.
func (tape Tape) HeadSymbol() rune {
	return tape.buffer[tape.currentPos]
}

// Shifts head to left.
// Adds additional space to buffer
// if it exceeds the buffer's
// limit to its left side.
func (tape *Tape) HeadToLeft() {
	if tape.currentPos == 0 {
		tape.enlarge(left)
	}
	tape.currentPos--
}

// Shifts head to right.
// Adds additional space to buffer
// if it exceeds the buffer's
// limit to its right side.
func (tape *Tape) HeadToRight() {
	tape.currentPos++
	if tape.currentPos == uint64(len(tape.buffer)) {
		tape.enlarge(right)
	}
}

// Sets symbol to head position.
func (tape *Tape) Set(symbol rune) {
	tape.buffer[tape.currentPos] = symbol
}

// Returns debug representation of tape
// in the form of string of the tape's
// content and head position.
func (tape Tape) String() string {
	top := string(tape.buffer)
	pos := make([]rune, tape.currentPos+1, len(tape.buffer)+1)
	for i := range pos {
		if uint64(i) == tape.currentPos {
			pos[i] = '^'
		} else {
			pos[i] = ' '
		}
	}
	return top + "\n" + string(pos)
}
