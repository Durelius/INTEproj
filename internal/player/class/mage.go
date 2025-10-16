package class

type Mage struct {
	baseDmg int
	mana    int
}

func NewMage() *Mage {
	return &Mage{
		baseDmg: 10,
		mana:    100,
	}
}

func (m *Mage) Name() string {
	return "Mage"
}

func (m *Mage) GetDescription() string {
	return "Magic caster."
}

func (m *Mage) GetBaseDmg() int {
	return m.baseDmg
}

func (m *Mage) IncreaseStats(level int) {
	m.baseDmg += level * 2
	m.mana += level * 10
}