package machine

import (
	"fmt"
	"strings"
)

const (
	FirstRotorEncoding  = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	SecondRotorEncoding = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	ThirdRotorEncoding  = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
)

var PossibleRotarOrders []string = []string{"III,II,I", "III,I,II", "II,III,I", "II,I,III", "I,III,II", "I,II,III"}
var RotorNumberEncodingMap map[string]string = map[string]string{"I": FirstRotorEncoding, "II": SecondRotorEncoding, "III": ThirdRotorEncoding}

type EngimaMachine struct {
	Right     *Rotor
	Middle    *Rotor
	Left      *Rotor
	Reflector *Reflector
	PlugBoard *PlugBoard
}

func NewEnigmaMachineRotorOrder(rotorOrder string) *EngimaMachine {
	split := strings.Split(rotorOrder, ",")
	enigma := EngimaMachine{Right: NewRotorFromName(split[2]), Middle: NewRotorFromName(split[1]), Left: NewRotorFromName(split[0]), Reflector: NewReflector()}
	enigma.PlugBoard = NewPlugBoard()
	return &enigma
}

func NewEnigmaMachineTest() *EngimaMachine {
	enigma := EngimaMachine{Right: NewRotor("I", FirstRotorEncoding, 16, 9, 16), Middle: NewRotor("II", SecondRotorEncoding, 16, 9, 4), Left: NewRotor("III", ThirdRotorEncoding, 16, 9, 21), Reflector: NewReflector()}
	enigma.PlugBoard = NewPlugBoard()
	//enigma.PlugBoard.addMappingPair('A', 'Z')
	// enigma.PlugBoard.addMappingPair('B', 'E')
	// enigma.PlugBoard.addMappingPair('C', 'J')
	// enigma.PlugBoard.addMappingPair('D', 'X')
	// enigma.PlugBoard.addMappingPair('F', 'Q')
	return &enigma
}
func NewEnigmaMachineTest2() *EngimaMachine {
	enigma := EngimaMachine{Right: NewRotor("I", FirstRotorEncoding, 24, 0, 16), Middle: NewRotor("II", SecondRotorEncoding, 24, 0, 4), Left: NewRotor("III", ThirdRotorEncoding, 23, 0, 21), Reflector: NewReflector()}
	enigma.PlugBoard = NewPlugBoard()
	//enigma.PlugBoard.addMappingPair('A', 'Z')
	// enigma.PlugBoard.addMappingPair('B', 'E')
	// enigma.PlugBoard.addMappingPair('C', 'J')
	// enigma.PlugBoard.addMappingPair('D', 'X')
	// enigma.PlugBoard.addMappingPair('F', 'Q')
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
	if e.Right.RotorPosition == e.Right.RotationPoint {
		e.Middle.rotate()
		if e.Middle.RotorPosition == e.Middle.RotationPoint {
			e.Left.rotate()
		}
	}
}
func (e *EngimaMachine) SetRotorPositions(leftPosition int32, middlePosition int32, rightPosition int32) {
	e.Left.RotorPosition = (leftPosition)
	e.Middle.RotorPosition = (middlePosition)
	e.Right.RotorPosition = (rightPosition)
}

func (e *EngimaMachine) SetRingSettings(left int32,middle int32, right int32) {
	e.Left.RingSetting = left
	e.Middle.RingSetting = middle
	e.Right.RingSetting = right
}

func (e *EngimaMachine) AddPlugboardPair(pair string) {
	e.PlugBoard.addMappingPair(rune(pair[0]), rune(pair[1]))
}

func (e *EngimaMachine) RemovePlugboardPair(pair string) {
	e.PlugBoard.removeMappingPair(rune(pair[0]), rune(pair[1]))
}

func (e *EngimaMachine) GenerateNewPlugboardFromSinglePair(pair string) {
	plugboard := NewPlugBoard()
	plugboard.addMappingPair(rune(pair[0]), rune(pair[1]))
	e.PlugBoard = plugboard
}

func (e *EngimaMachine) GeneratePlugBoardFromInitalPairs(pairs string) {
	fmt.Println("Pairs:", pairs)
	plugBoard := NewPlugBoardFromPairString(pairs)
	e.PlugBoard = plugBoard
}
