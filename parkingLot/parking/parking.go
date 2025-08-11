package parking

import "sync"

type Parking struct {
	Name    string
	Capcity int
	Space   []Vehicle
}

func (p *Parking) Park(registrationNumber string) (int , bool) {
	var parkingSlot int
	for i, vehicle := range p.Space {
		if vehicle.ParkingSlot == 0 {
			vehicle.ParkingSlot = i + 1
			vehicle.RegistrationNumber = registrationNumber
			parkingSlot = i
		}
	}
	return parkingSlot , true
}

func (p *Parking) Unpark(parkingSlot int, registrationNumber string) bool {
	vehicle := p.Space[parkingSlot]

	if vehicle.RegistrationNumber == registrationNumber {
		vehicle.ParkingSlot = 0
		vehicle.RegistrationNumber = ""
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
