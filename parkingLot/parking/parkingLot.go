package parking

type ParkingLot interface {
	Park(string) (int , bool)
	Unpark(int,string) bool
}

