package factory

type Musket struct {
	Gun 
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket Gun",
			power: 1,
		},
	}
}
