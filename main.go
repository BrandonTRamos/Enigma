package main

import (
	"Enigma/machine"
	"Enigma/machine/attack"
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	fmt.Println("Running...")
	for _, order := range machine.PossibleRotarOrders {
		wg.Add(1)
		go attack.Attack(order,wg)
	}
	wg.Wait()
}
