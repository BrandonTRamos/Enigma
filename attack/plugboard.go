package attack

import (
	"Enigma/machine"
	"fmt"
)

type PlugboardResult struct {
	IOC   float32
	pairs string
}

func plugboardAttack(encodedText string, topResult *AttackPermutationResult) *PlugboardResult {
	wiringCombos := generatePossibleWiringCombos()
	intialPlugBoardResult := &PlugboardResult{IOC: topResult.IOC, pairs: ""}
	enigma := machine.NewEnigmaMachineRotorOrder(topResult.Order)
	runPlugBoardAttack(encodedText, intialPlugBoardResult, wiringCombos, enigma, topResult, 0)
	fmt.Printf("%+v\n", intialPlugBoardResult)
	return nil
}

func runPlugBoardAttack(encodedText string, plugboardResult *PlugboardResult, wiringCombos []string, enigma *machine.EngimaMachine, topResult *AttackPermutationResult, depth int) {
	if depth > 1 {
		return
	}
	intialPairs := plugboardResult.pairs
	enigma.GeneratePlugBoardFromInitalPairs(intialPairs)
	for _, pair := range wiringCombos {
		enigma.SetRotorPositions(topResult.Rotors[0].Position, topResult.Rotors[1].Position, topResult.Rotors[2].Position)
		enigma.AddPlugboardPair(pair)
		decodedText := enigma.EncodeDecodeText(encodedText)
		enigma.RemovePlugboardPair(pair)
		ioc := CalcIndexOfCooincidence(decodedText)
		if pair == "AZ" {
			fmt.Println("AZ IOC: ", ioc)
			fmt.Println("AZ decode")
			fmt.Println(decodedText)
			fmt.Println("===============")
		}
		if (ioc > plugboardResult.IOC) && (ioc < .065) {
			fmt.Println("Pair found: ", pair, "Base IOC:", plugboardResult.IOC, "New IOC:", ioc)
			plugboardResult.IOC = ioc
			topResult.PlugBoardDecodedText = decodedText
			if intialPairs == "" {
				plugboardResult.pairs = pair
			} else {
				plugboardResult.pairs = intialPairs + "," + pair
			}

		}

	}
	runPlugBoardAttack(encodedText, plugboardResult, wiringCombos, enigma, topResult, depth+1)
}

func generatePossibleWiringCombos() []string {
	combos := make([]string, 0, 325)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			pair := string(rune(alphabet[i])) + string(rune(alphabet[j]))
			combos = append(combos, pair)
		}
	}

	return combos
}
