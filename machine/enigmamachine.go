package machine

import "strings"

const (
	firstRotorEncoding  = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	secondRotorEncoding = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	thirdRotorEncoding  = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
)

var PossibleRotarOrders []string = []string{"III,II,I", "III,I,II", "II,III,I", "II,I,III", "I,III,II", "I,II,III"}
var RotorNumberEncodingMap map[string]string = map[string]string{"I": firstRotorEncoding, "II": secondRotorEncoding, "III": thirdRotorEncoding}

type EngimaMachine struct {
	Right     *Rotor
	Middle    *Rotor
	Left      *Rotor
	Reflector *Reflector
	PlugBoard *PlugBoard
}

func NewEnigmaMachineRotorOrder(rotorOrder string) *EngimaMachine {
	split := strings.Split(rotorOrder, ",")
	enigma := EngimaMachine{Right: NewRotorFromName(split[2]), Middle: NewRotorFromName(split[1]), Left: NewRotorFromName(split[0]),Reflector: NewReflector()}
	enigma.PlugBoard = NewPlugBoard()
	return &enigma
}

func NewEnigmaMachineTest() *EngimaMachine {
	enigma := EngimaMachine{Right: NewRotor("I", firstRotorEncoding, 25, 1, 16), Middle: NewRotor("II", secondRotorEncoding, 1, 25, 4), Left: NewRotor("III", thirdRotorEncoding, 1, 13, 21), Reflector: NewReflector()}
	enigma.PlugBoard = NewPlugBoard()
	enigma.PlugBoard.addMappingPair('A', 'Z')
	enigma.PlugBoard.addMappingPair('B', 'E')
	enigma.PlugBoard.addMappingPair('C', 'J')
	enigma.PlugBoard.addMappingPair('D', 'X')
	enigma.PlugBoard.addMappingPair('F', 'Q')
	return &enigma
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
	swappedLetter := e.PlugBoard.SwapLetter(letter)
	char := swappedLetter - 65
	char = e.Right.encodeForward(char)
	char = e.Middle.encodeForward(char)
	char = e.Left.encodeForward(char)
	char = e.Reflector.Reflect(char)
	char = e.Left.encodeReverse(char)
	char = e.Middle.encodeReverse(char)
	char = e.Right.encodeReverse(char)
	return e.PlugBoard.SwapLetter(char + 65)
}

func (e *EngimaMachine) rotateRotars() {
	e.Right.rotate()
	if e.Right.rotorPosition == e.Right.rotationPoint {
		e.Middle.rotate()
		// if(e.Middle.rotorPosition==e.Middle.rotationPoint){
		// 	e.Left.rotate()
		// }
	}
}
