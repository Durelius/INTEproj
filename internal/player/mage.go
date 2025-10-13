package player

type Mage struct {
	Player
	mana int
}

const (
	base_mana = 100
)

func newMage(player Player) Player {
	return &Mage{mana: base_mana, Player: player}
}
