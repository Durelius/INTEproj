package class

type Rogue struct {
	baseDmg int
	stealth int
}

func NewRogue() *Rogue {
	return &Rogue{
		baseDmg: 7,
		stealth: 100,
	}
}

func (r *Rogue) Name() string {
	return "Rogue"
}

func (r *Rogue) GetDescription() string {
	return "Stealth assassin."
}

func (r *Rogue) GetBaseDmg() int {
	return r.baseDmg
}

func (r *Rogue) IncreaseStats(level int) {
	r.baseDmg += level * 2
	r.stealth += level * 10
}
