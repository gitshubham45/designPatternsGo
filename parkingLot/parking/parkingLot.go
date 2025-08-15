package parking

import (
	"sync"
	"time"
)

type Parking interface {
	Park(string) (int, bool)
	Unpark(int, string) bool
}

type RateCard map[VehicleType]float64

type Ticket struct {
	SlotID    int
	Vehicle   Vehicle
	EntryTime time.Time
}

type ParkingLot struct {
	Name     string
	Capacity int
	Slots    []*ParkingSlot
	Tickets  map[string]Ticket
	Rates    RateCard
}

var (
	parkingLotInstance *ParkingLot
	once            sync.Once
)

func NewParkingLot(capacity int, name string) *ParkingLot {
	once.Do(func() {
		parkingLotInstance = &ParkingLot{
			Name:     name,
			Capacity: capacity,
			Slots:    make([]*ParkingSlot, capacity),
			Tickets:  make(map[string]Ticket),
			Rates: RateCard{
				Car:   20.0,
				Bike:  10.0,
				Truck: 50.0,
			},
		}
	})
	return parkingLotInstance
}
