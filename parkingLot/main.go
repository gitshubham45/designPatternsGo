package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/parkingLot/parking"
)

func main() {
	parking := parking.NewParking(10, "Central Parking")

	var registrationNumber string
	fmt.Printf("Welcome to %s : \n Please enter the registration number of vehicle to park \n", parking.Name)
	fmt.Scanln(&registrationNumber)

	slot, ok := parking.Park(registrationNumber)
	if !ok {
		fmt.Println("Error parking the vehicle")
	}

	fmt.Printf("Parked -> parkign slot : %d\n", slot)
}
