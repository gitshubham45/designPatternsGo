package parking

type ParkingSlot struct {
	ID      int
	Vehicle *Vehicle
}

func (p *ParkingLot) Park(registrationNumber string,vt VehicleType) (int, bool) {
	for i := range p.Slots {
		if p.Slots[i] == nil {
			p.Slots[i].ID = i + 1
			p.Slots[i].Vehicle.RegistrationNumber = registrationNumber
			p.Slots[i].Vehicle.Type = vt
			return i + 1, true
		}
	}
	return 0, false
}

func (p *ParkingLot) Unpark(parkingNumber int, registrationNumber string) bool {
	index := parkingNumber - 1
	if index < 0 || index >= len(p.Slots) {
		return false
	}

	if p.Slots[index].Vehicle.RegistrationNumber == registrationNumber {
		p.Slots[index].ID = 0
		p.Slots[index].Vehicle.RegistrationNumber = ""
		return true
	}
	return false
}
