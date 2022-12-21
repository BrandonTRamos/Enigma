package attack

import (
	"Enigma/machine"
	"sort"
	"sync"
)

type RotorSummary struct {
	Name        string
	Position    int32
	RingSetting int32
}

type AttackPermutationResult struct {
	IOC         float32
	Rotors      [3]*RotorSummary
	DecodedText string
}

func Attack(text string, order string, wg *sync.WaitGroup) *AttackPermutationResult {
	defer wg.Done()
	results := intializeResultArray()
	enigma := machine.NewEnigmaMachineRotorOrder(order)
	for i := 0; i < 26; i++ {
		//left rotor
		for j := 0; j < 26; j++ {
			//middle rotor
			for k := 0; k < 26; k++ {
				//right rotor
				enigma.Left.RotorPosition = (int32(i))
				enigma.Middle.RotorPosition = (int32(j))
				enigma.Right.RotorPosition = (int32(k))
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
					results[9] = &AttackPermutationResult{IOC: ioc, Rotors: rotors, DecodedText: decodedText}
					SortArrayDesc(results[:])
				}
			}
		}
	}
	return results[0]
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
