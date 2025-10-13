package player

import (
	"github.com/Durelius/INTEproj/internal/item"
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
	paladin.SetEquippedItem(item.WEAR_POSITION_WEAPON, item.New(&item.IRON_SWORD))
	return paladin
}
