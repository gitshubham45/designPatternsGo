package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/parkingLot/parking"
)

func main() {
	parking := parking.NewParking(10, "Central Parking")

	fmt.Printf("Welcome to %s\n", parking.Name)
	// fmt.Println("Type vehicle registration number to park, or 'exit' to quit.")

	for {
		var requestType string
		fmt.Println("Enter Type of Request:[park/unpark/exit]")
		fmt.Scanln(&requestType)

		switch requestType {
		case "park":
			var registrationNumber string
			fmt.Print("Enter registration number: ")
			fmt.Scanln(&registrationNumber)

			slot, ok := parking.Park(registrationNumber)
			if !ok {
				fmt.Println("Error: Parking is full or invalid input.")
				continue
			}
			fmt.Printf("Parked -> parking slot: %d\n", slot)
		case "unpark" :
			var registrationNumber string
			var parkingNumber int
			// fmt.Print("Enter registration number and Parking Number: ")
			// fmt.Scanln(&registrationNumber,&parkingNumber)

			fmt.Print("Enter registration number : ")
			fmt.Scanln(&registrationNumber)
			fmt.Println("Enter parking number : ")
			fmt.Scanln(&parkingNumber)

			ok := parking.Unpark(parkingNumber,registrationNumber)
			if !ok {
				fmt.Println("Error: Registration Number or Parking Number is not correct.")
				continue
			}
			fmt.Printf("Unparked -> Registration Number: %s\n", registrationNumber)
		case "exit" :
			fmt.Println("Exiting... Goodbye!")
			return
		}

		
	}
}
