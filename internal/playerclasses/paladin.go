package playerclasses

type Paladin struct {
	Level      int
	Experience int
	BaseDmg    int
	Mana       int
	holiness   int
}

func NewPaladin() *Paladin {
	return &Paladin{
		Level:      1,
		Experience: 0,
		BaseDmg:    7,
		Mana:       50,
		holiness:   100,
	}
}

func (p *Paladin) Name() string {
	return "Paladin"
}

func (p *Paladin) GetDescription() string {
	return "Magic tank."
}

func (p *Paladin) GetBaseDmg() int {
	return p.BaseDmg
}