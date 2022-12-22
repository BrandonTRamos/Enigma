package attack

import (
	"Enigma/machine"
	"fmt"
)


func ringSettingAttack(encodedText string, topResult *AttackPermutationResult) {
	enigma := machine.NewEnigmaMachineRotorOrder(topResult.Order)
	topIoc:=topResult.IOC
	var bestMiddleRingSetting int32
	var bestRightRingSetting int32
	var bestLeftRingSetting int32
	for i:=0;i<26;i++{
		for j:=0;j<26;j++{
			for k:=0; k<26;k++{
				leftRingSetting:=int32(k)
				middleRingSetting:=int32(j)
				rightRingSetting:=int32(i)
				enigma.SetRotorPositions(topResult.Rotors[0].Position, topResult.Rotors[1].Position, topResult.Rotors[2].Position)
				enigma.SetRingSettings(leftRingSetting,middleRingSetting,rightRingSetting)
				decodedText:=enigma.EncodeDecodeText(encodedText)
				ioc:=CalcIndexOfCooincidence(decodedText)
				if i==12 &&j==12&&k==12{
					fmt.Println("alll 12s ioc",ioc)
				}
				if i==0 &&j==0&&k==0{
					fmt.Println("alll 0s ioc",ioc)
				}
				if ioc>topIoc{
					topIoc=ioc
					bestMiddleRingSetting=middleRingSetting
					bestRightRingSetting=rightRingSetting
					bestLeftRingSetting=leftRingSetting
				}
			}
		}
	}
	fmt.Println("Best Ioc: ",topIoc,"bestLeftRingSetting",bestLeftRingSetting,"bestMiddleRingSetting",bestMiddleRingSetting,"bestRightRingSetting",bestRightRingSetting)
}
