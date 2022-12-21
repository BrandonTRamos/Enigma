package attack

import (
	"Enigma/machine"
	"fmt"
	"sync"
)

func Attack(order string,wg *sync.WaitGroup){
	defer wg.Done()
	encoderEnigma := machine.NewEnigmaMachineRotorOrder(order)
	decoderEnigma := machine.NewEnigmaMachineRotorOrder(order)
	original := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS"
	encoded := encoderEnigma.EncodeDecodeText(original)
	decoded := decoderEnigma.EncodeDecodeText(encoded)
	fmt.Println("Order:",order,"Original Message:", original, "->", "Encoded Message:", encoded, "->", "Decoded Message:", decoded)
}