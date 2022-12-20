package machine

const (
	firstRotorEncoding  = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	secondRotorEncoding = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	thirdRotorEncoding  = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
)

type EngimaMachine struct {
	I         *Rotor
	II        *Rotor
	III       *Rotor
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
	char := letter - 65
	char = e.I.encodeForward(char)
	char = e.II.encodeForward(char)
	char = e.III.encodeForward(char)
	char = e.Reflector.Reflect(char)
	char = e.III.encodeReverse(char)
	char = e.II.encodeReverse(char)
	char = e.I.encodeReverse(char)
	return char + 65
}

func NewEnigmaMachine() *EngimaMachine {
	enigma := EngimaMachine{I: NewRotor("I", firstRotorEncoding), II: NewRotor("II", secondRotorEncoding), III: NewRotor("III", thirdRotorEncoding), Reflector: NewReflector()}
	return &enigma
}
