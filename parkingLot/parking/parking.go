package parking

import (
	"fmt"
	"sync"
)

type Parking struct {
	Name    string
	Capcity int
	Space   []Vehicle
}

func (p *Parking) Park(registrationNumber string) (int , bool) {
	fmt.Println(p.Space)
	for i := range p.Space {
		if p.Space[i].RegistrationNumber == "" {
			p.Space[i].ParkingNumber = i + 1
			p.Space[i].RegistrationNumber = registrationNumber
			return p.Space[i].ParkingNumber, true
		}
	}
	return 0, false
}

func (p *Parking) Unpark(parkingNumber int, registrationNumber string) bool {
	index := parkingNumber - 1
	if index < 0 || index >= len(p.Space) {
		return false
	}

	if p.Space[index].RegistrationNumber == registrationNumber {
		p.Space[index].ParkingNumber = 0
		p.Space[index].RegistrationNumber = ""
		return true
	}
	return false
}

var (
	parkingInstance *Parking
	once            sync.Once
)

func NewParking(capacity int, name string) *Parking {
	once.Do(func() {
		parkingInstance = &Parking{
			Name:    name,
			Capcity: capacity,
			Space:   make([]Vehicle, capacity),
		}
	})
	return parkingInstance
}
