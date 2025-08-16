package parking

type VehicleType string

const (
	Car   VehicleType = "Car"
	Bike  VehicleType = "Bike"
	Truck VehicleType = "Truck"
	Bus   VehicleType = "Bus"
)

type Vehicle struct {
	RegistrationNumber string
	Type VehicleType
}


