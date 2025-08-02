package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/simpleFactory/factory"
)

func main() {
	ak47, _ := factory.GetGun("AK47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
