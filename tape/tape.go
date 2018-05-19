package tape

type Tape struct {
	buffer        []rune
	zeroPos       uint64
	currentPos    uint64
	defaultSymbol rune
}

func Create(content string, defaultSymbol rune) Tape {
	return Tape{
		buffer:        []rune(content),
		zeroPos:       0,
		currentPos:    0,
		defaultSymbol: defaultSymbol,
	}
}

type leftRight bool

const (
	left  leftRight = true
	right leftRight = false
)

const rellocationLenght uint64 = 16

func (tape *Tape) reallocate(lr leftRight) {
	switch lr {
	case left:
		leftBuf := make([]rune, rellocationLenght)
		for i := range leftBuf {
			leftBuf[i] = tape.defaultSymbol
		}
		*tape = Tape{
			buffer:        append(leftBuf, tape.buffer...),
			zeroPos:       rellocationLenght + tape.zeroPos,
			currentPos:    rellocationLenght + tape.currentPos,
			defaultSymbol: tape.defaultSymbol,
		}
	case right:
		rightBuf := make([]rune, rellocationLenght)
		for i := range rightBuf {
			rightBuf[i] = tape.defaultSymbol
		}
		tape.buffer = append(tape.buffer, rightBuf...)
	}
}

func (tape Tape) Now() rune {
	return tape.buffer[tape.currentPos]
}

func (tape *Tape) LeftShifting() rune {
	if tape.currentPos == 0 {
		tape.reallocate(left)
	}
	tape.currentPos--
	return tape.Now()
}

func (tape *Tape) RightShifting() rune {
	tape.currentPos++
	if tape.currentPos == uint64(len(tape.buffer)) {
		tape.reallocate(right)
	}
	return tape.Now()
}

func (tape Tape) String() string {
	top := string(tape.buffer)
	pos := make([]rune, tape.currentPos+1, len(tape.buffer))
	for i := range pos {
		if uint64(i) == tape.currentPos {
			pos[i] = '^'
		} else {
			pos[i] = ' '
		}
	}
	return top + "\n" + string(pos)
}
