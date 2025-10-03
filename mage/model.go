package rouge

import (
	"INTE/projekt/player"
)

type Mage struct {
	player.Player
	mana int
}

const (
	base_mana = 100
)

func New() (*Mage, error) {
	return &Mage{mana: base_mana}, nil
}
