package playerclasses

type Mage struct {
	Level      int
	Experience int
	BaseDmg    int
	Mana       int
}

func NewMage() *Mage {
	return &Mage{
		Level:      1,
		Experience: 0,
		BaseDmg:    5,
		Mana:       100,
	}
}

func (m *Mage) Name() string {
	return "Mage"
}

func (m *Mage) GetDescription() string {
	return "Magic caster."
}

func (m *Mage) GetBaseDmg() int {
	return m.BaseDmg
}