package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/parkingLot/parking"
)

func main() {
	// var myParking *parking.ParkingLot
	myParking := parking.NewParkingLot(20, "Central Parking")

	fmt.Printf("Welcome to %s\n", myParking.Name)
	// fmt.Println("Type vehicle registration number to park, or 'exit' to quit.")

	for {
		var requestType string
		fmt.Println("Enter Type of Request:[park/unpark/exit]")
		fmt.Scanln(&requestType)

		switch requestType {
		case "park":
			var registrationNumber string
			var vehicleType parking.VehicleType
			fmt.Print("Enter registration number: ")
			fmt.Scanln(&registrationNumber)
			fmt.Println("Enter Vehicle Type:[Bike,Bus,Truck,Car]")
			fmt.Scanln(&vehicleType)

			ticket, ok := myParking.Park(registrationNumber, vehicleType)
			if !ok {
				fmt.Println("Error: Parking is full or invalid input.")
				continue
			}
			fmt.Printf("Parked -> parking Details: %+v\n", ticket)
		case "unpark":
			var parkingNumber int
			// fmt.Print("Enter registration number and Parking Number: ")
			// fmt.Scanln(&registrationNumber,&parkingNumber)

			fmt.Println("Enter parking number :[SlotID]")
			fmt.Scanln(&parkingNumber)

			ticket,ok := myParking.Unpark(parkingNumber)
			if !ok {
				fmt.Println("Error: Registration Number or Parking Number is not correct.")
				continue
			}

			billAmount := myParking.CalculteBill(ticket)
			fmt.Println("registration" , ticket.Vehicle.RegistrationNumber)
			fmt.Printf("Unparked -> Registration Number: %s , Please pay amount : Rs. %.2f\n",
	ticket.Vehicle.RegistrationNumber, billAmount)

		case "exit":
			fmt.Println("Exiting... Goodbye!")
			return
		}
	}
}
