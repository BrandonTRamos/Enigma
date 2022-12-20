package main

import (
	"Enigma/machine"
	"fmt"
)

func main() {
	encoderEnigma := machine.NewEnigmaMachine()
	decoderEnigma := machine.NewEnigmaMachine()
	original := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS"
	encoded := encoderEnigma.EncodeDecodeText(original)
	decoded := decoderEnigma.EncodeDecodeText(encoded)
	fmt.Println(original, "->", encoded, "->", decoded)
}
