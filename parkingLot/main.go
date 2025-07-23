package main

// type ParkingSlot struct{
// 	Position inti
// }

type Vehicle interface{
	Park([]int) bool
	UnPark([]int) bool
}

type TwoWheeler struct{
	RegistrationNumber string
	OwnerName string
	OwnerNumber string
	ParkingSlotNumber int
}

func (tw *TwoWheeler) Park(slot map[int]int){
	slot[]
}

func (tw *TwoWheeler) UnPark(slot []int){
	
}

func main(){

}


