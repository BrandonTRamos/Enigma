package attack

import (
	"Enigma/machine"
	"fmt"
)

func Attack(encodedText string) {
	fmt.Println("===========")
	fmt.Println("Initiating order and position attack")
	topResult := orderAndPositionAttack(encodedText)
	fmt.Println("===========")
	fmt.Println("Initiating Ring setting attack")
	ringSettingAttack(encodedText, topResult)
	fmt.Println("===========")
	fmt.Println("Initiating plugboard attack")
	// plugboardAttack(encodedText, topResult)
	// fmt.Println(topResult.PlugBoardDecodedText)
	engima2:=machine.NewEnigmaMachineTest2()
	decoded:=engima2.EncodeDecodeText(encodedText)
	ioc:=CalcIndexOfCooincidence(decoded)
	fmt.Println("enigma2 ioc:",ioc)
	fmt.Println(decoded)
}

func printBestResult(bestResults []*AttackPermutationResult) {
	fmt.Println("Top Result:")
	fmt.Printf("IOC: %#v, ", bestResults[0].IOC)
	fmt.Printf("%+v, ", bestResults[0].Rotors[0])
	fmt.Printf("%+v, ", bestResults[0].Rotors[1])
	fmt.Printf("%+v\n\n", bestResults[0].Rotors[2])
	fmt.Println("Decoded Text:")
	fmt.Println(bestResults[0].OrderAndPositionDecodedText)
}
