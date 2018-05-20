package tape

type Tape struct {
	buffer      []rune
	zeroPos     uint64
	currentPos  uint64
	emptySymbol rune
}

func Create(content string, emptySymbol rune) Tape {
	return Tape{
		buffer:      []rune(content),
		zeroPos:     0,
		currentPos:  0,
		emptySymbol: emptySymbol,
	}
}

type leftRight bool

const (
	left  leftRight = true
	right leftRight = false
)

const reallocationLength uint64 = 16

func (tape *Tape) reallocate(lr leftRight) {
	switch lr {
	case left:
		leftBuf := make([]rune, reallocationLength)
		for i := range leftBuf {
			leftBuf[i] = tape.emptySymbol
		}
		*tape = Tape{
			buffer:      append(leftBuf, tape.buffer...),
			zeroPos:     reallocationLength + tape.zeroPos,
			currentPos:  reallocationLength + tape.currentPos,
			emptySymbol: tape.emptySymbol,
		}
	case right:
		rightBuf := make([]rune, reallocationLength)
		for i := range rightBuf {
			rightBuf[i] = tape.emptySymbol
		}
		tape.buffer = append(tape.buffer, rightBuf...)
	}
}

func (tape Tape) Now() rune {
	if tape.currentPos == 0 && len(tape.buffer) == 0 {
		return tape.emptySymbol
	}
	return tape.buffer[tape.currentPos]
}

func (tape *Tape) LeftShift() {
	if tape.currentPos == 0 {
		tape.reallocate(left)
	}
	tape.currentPos--
}

func (tape *Tape) RightShift() {
	tape.currentPos++
	if tape.currentPos == uint64(len(tape.buffer)) {
		tape.reallocate(right)
	}
}

func (tape Tape) Set(symbol rune) {
	tape.buffer[tape.currentPos] = symbol
}

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
