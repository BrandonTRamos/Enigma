package attack

import (
	"fmt"
)

func Attack(encodedText string) {
	fmt.Println("===========")
	fmt.Println("Initiating order and position attack")
	topResult := orderAndPositionAttack(encodedText)
	ringSettingAttack(encodedText, topResult)
	fmt.Println("===========")
	fmt.Println("Initiating plugboard attack")
	plugboardAttack(encodedText, topResult)
	fmt.Println(topResult.PlugBoardDecodedText)
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
