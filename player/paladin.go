package player

import (
	"INTE/projekt/character"
	"INTE/projekt/item"
)

type Paladin struct {
	Player
	holiness int
}

const (
	base_holiness = 100
)

func newPaladin(player Player) Player {
	paladin := &Paladin{holiness: base_holiness, Player: player}
	paladin.SetItem(character.WEAR_POSITION_LEFT_ARM, item.New(item.IRON_SWORD))
	return paladin
}
