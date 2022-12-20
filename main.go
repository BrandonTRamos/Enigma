package main

import (
	"Enigma/machine"
	"fmt"
)

func main() {
	enigma := machine.NewEnigmaMachine()
	original := "HELLOWORLD"
	encoded := enigma.EncodeDecodeText(original)
	decoded := enigma.EncodeDecodeText(encoded)
	fmt.Println(original, "->", encoded, "->", decoded)
}
