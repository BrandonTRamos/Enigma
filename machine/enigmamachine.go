package machine

const (
	firstRotorEncoding  = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	secondRotorEncoding = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	thirdRotorEncoding  = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
)

type EngimaMachine struct {
	Right     *Rotor
	Middle    *Rotor
	Left      *Rotor
	Reflector *Reflector
}

func (e *EngimaMachine) EncodeDecodeText(text string) string {
	encodedDecodedRuneSlice := []rune{}
	for _, letter := range text {
		encodedDecodedRuneSlice = append(encodedDecodedRuneSlice, e.encodeDecodeLetter(letter))
	}

	return string(encodedDecodedRuneSlice)
}

func (e *EngimaMachine) encodeDecodeLetter(letter rune) rune {
	e.rotateRotars()
	char := letter - 65
	char = e.Right.encodeForward(char)
	char = e.Middle.encodeForward(char)
	char = e.Left.encodeForward(char)
	char = e.Reflector.Reflect(char)
	char = e.Left.encodeReverse(char)
	char = e.Middle.encodeReverse(char)
	char = e.Right.encodeReverse(char)
	return char + 65
}

func (e *EngimaMachine) rotateRotars() {
	e.Right.rotate()
}

func NewEnigmaMachine() *EngimaMachine {
	enigma := EngimaMachine{Right: NewRotor("I", firstRotorEncoding,2), Middle: NewRotor("II", secondRotorEncoding,25), Left: NewRotor("III", thirdRotorEncoding,24), Reflector: NewReflector()}
	return &enigma
}
