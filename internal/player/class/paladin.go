package class

import "math"

type Paladin struct {
	baseDmg  int
	holiness int
}

func NewPaladin() *Paladin {
	return &Paladin{
		baseDmg:  5,
		holiness: 100,
	}
}

func (p *Paladin) Name() string {
	return "Paladin"
}

func (p *Paladin) GetDescription() string {
	return "Magic tank."
}

func (p *Paladin) GetBaseDmg() int {
	return p.baseDmg
}

func (p *Paladin) IncreaseStats(level int) {
	p.baseDmg += int(math.Round(float64(level) * 1.5))
	p.holiness += level * 10
}