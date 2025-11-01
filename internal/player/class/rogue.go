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
func LoadRogue(baseDmg, stealth int) *Rogue {
	return &Rogue{
		baseDmg: baseDmg,
		stealth: stealth,
	}
}

const ROGUE_STR ClassName = "Rogue"

func (r *Rogue) Name() ClassName {
	return ROGUE_STR
}

func (r *Rogue) GetDescription() string {
	return "Stealth assassin."
}

func (r *Rogue) GetBaseDmg() int {
	return r.baseDmg
}
func (r *Rogue) GetEnergy() int {
	return r.stealth
}

func (r *Rogue) IncreaseStats(level int) {
	r.baseDmg += level * 2
	r.stealth += level * 10
}
