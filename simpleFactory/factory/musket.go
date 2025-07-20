package factory

type musket struct {
	Gun
}

func NewMusket() IGun{
	return &musket{
		Gun : Gun{
			name : "Musket Gun",
			power : 1,
		},
	}
}