package playerclasses

type Rogue struct {
	Level      int
	Experience int
	BaseDmg    int
	stealth    int
}

func NewRogue() *Rogue {
	return &Rogue{
		Level:      1,
		Experience: 0,
		BaseDmg:    7,
		stealth:    100,
	}
}

func (r *Rogue) Name() string {
	return "Rogue"
}

func (r *Rogue) GetDescription() string {
	return "Stealth assassin."
}

func (r *Rogue) GetBaseDmg() int {
	return r.BaseDmg
}
