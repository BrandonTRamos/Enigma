package attack

import (
	"Enigma/machine"
	"fmt"
	"sort"
)

type RotorSummary struct {
	Name        string
	Position    int32
	RingSetting int32
}

type AttackPermutationResult struct {
	IOC         float32
	Order string
	Rotors      [3]*RotorSummary
	PlugSummary string
	DecodedText string
}

func Attack(encodedText string) {
	fmt.Println("===========")
	fmt.Println("Initiating order and position attack")
	topResult:=orderAndPositionAttack(encodedText)
	ringSettingAttack(encodedText,topResult)
	fmt.Println("===========")
	fmt.Println("Initiating plugboard attack")
	plugboardAttack(encodedText,topResult)
}

func ringSettingAttack(encodedText string, topResult *AttackPermutationResult){

}

func plugboardAttack(encodedText string, topResult *AttackPermutationResult){
	wiringCombos:=generatePossibleWiringCombos()
	enigma := machine.NewEnigmaMachineRotorOrder(topResult.Order)
	baseIoc:=topResult.IOC
	for _,pair:=range wiringCombos{
		enigma.SetRotorPositions(topResult.Rotors[0].Position,topResult.Rotors[1].Position,topResult.Rotors[2].Position)
		enigma.GenerateNewPlugboardFromSinglePair(pair)
		decodedText:=enigma.EncodeDecodeText(encodedText)
		ioc:=CalcIndexOfCooincidence(decodedText)

		if ioc>baseIoc{
			fmt.Println("Pair found: ",pair,"Base IOC:",baseIoc,"New IOC:",ioc)
		}

	}
}

func generatePossibleWiringCombos()[] string{
	combos:= make([]string,0,325)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			pair:=string(rune(alphabet[i]))+string(rune(alphabet[j]))
			combos=append(combos, pair)
		}
	}

	return combos;
}

func orderAndPositionAttack(encodedText string) *AttackPermutationResult{
	resultChannel := make(chan *AttackPermutationResult)
	var bestResults []*AttackPermutationResult
	for _, order := range machine.PossibleRotarOrders {
		go bruteForceOrderAndPositions(encodedText, order, resultChannel)

	}
	for i := 0; i < len(machine.PossibleRotarOrders); i++ {
		result := <-resultChannel
		bestResults = append(bestResults, result)
	}
	close(resultChannel)

	SortArrayDesc(bestResults)
	printBestResult(bestResults)
	return bestResults[0]
}

func bruteForceOrderAndPositions(text string, order string, resultChannel chan *AttackPermutationResult){
	results := intializeResultArray()
	enigma := machine.NewEnigmaMachineRotorOrder(order)
	for i := 0; i < 26; i++ {
		//left rotor
		for j := 0; j < 26; j++ {
			//middle rotor
			for k := 0; k < 26; k++ {
				//right rotor
				enigma.SetRotorPositions(int32(i),int32(j),int32(k))
				decodedText := enigma.EncodeDecodeText(text)
				ioc := CalcIndexOfCooincidence(decodedText)
				currentWorstIoc := results[9].IOC
				if ioc > currentWorstIoc {
					var rotors [3]*RotorSummary
					leftRotor := &RotorSummary{Name: enigma.Left.Name, Position: int32(i), RingSetting: enigma.Left.GetRingSetting()}
					middleRotor := &RotorSummary{Name: enigma.Middle.Name, Position: int32(j), RingSetting: enigma.Middle.GetRingSetting()}
					rightRotor := &RotorSummary{Name: enigma.Right.Name, Position: int32(k), RingSetting: enigma.Right.GetRingSetting()}
					rotors[0] = leftRotor
					rotors[1] = middleRotor
					rotors[2] = rightRotor
					results[9] = &AttackPermutationResult{IOC: ioc, Order: order,Rotors: rotors, DecodedText: decodedText}
					SortArrayDesc(results[:])
				}
			}
		}
	}
	resultChannel <- results[0]
}

func intializeResultArray() [10]*AttackPermutationResult {
	var results [10]*AttackPermutationResult
	for i := 0; i < 10; i++ {
		results[i] = &AttackPermutationResult{IOC: 0.0}
	}
	return results
}

func SortArrayDesc(results []*AttackPermutationResult) {
	sort.Slice(results, func(i, j int) bool {
		return results[j].IOC < results[i].IOC
	})
}


func printBestResult(bestResults []*AttackPermutationResult){
	fmt.Println("Top Result:")
	fmt.Printf("IOC: %#v, ", bestResults[0].IOC)
	fmt.Printf("%+v, ", bestResults[0].Rotors[0])
	fmt.Printf("%+v, ", bestResults[0].Rotors[1])
	fmt.Printf("%+v\n\n", bestResults[0].Rotors[2])
	fmt.Println("Decoded Text:")
	fmt.Println(bestResults[0].DecodedText)
}