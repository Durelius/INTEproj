package player

type Mage struct {
	Player
	mana int
}

const (
	base_mana = 100
)

func newMage() *Mage {
	return &Mage{mana: base_mana}
}
