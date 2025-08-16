package parking

import (
	"fmt"
	"time"
)

type ParkingSlot struct {
	ID      int
	Vehicle *Vehicle
}

func (p *ParkingLot) Park(registrationNumber string, vt VehicleType) (*Ticket, bool) {
	for i := range p.Slots {
		if p.Slots[i] == nil {

			ps := &ParkingSlot{ID: i + 1}
			v := &Vehicle{
				RegistrationNumber: registrationNumber,
				Type:               vt,
			}
			ps.Vehicle = v
			p.Slots[i] = ps
			t := &Ticket{
				SlotID:    ps.ID,
				Vehicle:   v,
				EntryTime: time.Now(),
			}
			p.Tickets[ps.ID] = *t
			return t, true
		}
	}
	return nil, false
}

func (p *ParkingLot) Unpark(parkingSlot int) (*Ticket, bool) {
	ticket := p.Tickets[parkingSlot]
	index := ticket.SlotID - 1
	if index < 0 || index >= len(p.Slots) {
		return nil, false
	}

	registrationNumber := ticket.Vehicle.RegistrationNumber
	fmt.Println("registrationNumber", registrationNumber)

	if p.Slots[index].Vehicle.RegistrationNumber == registrationNumber {
		p.Slots[index].ID = 0
		p.Slots[index].Vehicle.RegistrationNumber = ""
		return &ticket, true
	}
	return nil, false
}

func (p *ParkingLot) CalculteBill(ticket *Ticket) float64 {
	// parkingDuration := time.Now().Sub(ticket.EntryTime)
	parkingDuration := time.Since(ticket.EntryTime)

	hours := parkingDuration.Hours()
	billedHours := int(hours)
	if hours > float64(billedHours) {
		billedHours++
	}
	if billedHours == 0 {
		billedHours = 1
	}

	rate := p.Rates[ticket.Vehicle.Type]
	bill := float64(billedHours) * rate
	return bill
}
