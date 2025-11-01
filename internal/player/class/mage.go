package class

const MAGE_STR ClassName = "Mage"

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
func LoadMage(baseDmg, mana int) *Mage {
	return &Mage{
		baseDmg: baseDmg,
		mana:    mana,
	}
}

func (m *Mage) Name() ClassName {
	return MAGE_STR
}

func (m *Mage) GetDescription() string {
	return "Magic caster."
}

func (m *Mage) GetBaseDmg() int {
	return m.baseDmg
}
func (m *Mage) GetEnergy() int {
	return m.mana
}

func (m *Mage) IncreaseStats(level int) {
	m.baseDmg += level * 2
	m.mana += level * 10
}
