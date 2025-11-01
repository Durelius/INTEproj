package class

import "math"

type Paladin struct {
	baseDmg  int
	holiness int
}

const PALADIN_STR ClassName = "Paladin"

func NewPaladin() *Paladin {
	return &Paladin{
		baseDmg:  5,
		holiness: 100,
	}
}
func LoadPaladin(baseDmg, holiness int) *Paladin {
	return &Paladin{
		baseDmg:  baseDmg,
		holiness: holiness,
	}
}

func (p *Paladin) Name() ClassName {
	return PALADIN_STR
}

func (p *Paladin) GetDescription() string {
	return "Magic tank."
}

func (p *Paladin) GetBaseDmg() int {
	return p.baseDmg
}
func (p *Paladin) GetEnergy() int {
	return p.holiness
}

func (p *Paladin) IncreaseStats(level int) {
	p.baseDmg += int(math.Round(float64(level) * 1.5))
	p.holiness += level * 10
}
