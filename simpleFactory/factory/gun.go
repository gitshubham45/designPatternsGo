package factory

type Gun struct{
	name string
	power int
}

func (g *Gun) SetName(name string){
	g.name = name
}

func (g *Gun) SetPower(power int){
	g.power = power
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) GetPower() int{
	return g.power
}

func NewGun(name string , power int) *Gun{
	return &Gun{
		name : name,
		power: power,
	}
}