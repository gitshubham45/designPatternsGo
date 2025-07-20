package factory

type AK47 struct{
	Gun
}

func NewAK47() IGun {
	return &AK47{
		Gun:Gun{
			name : "AK47 gun",
			power: 4,
		},
	}
}