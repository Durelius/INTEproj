package rouge

import (
	"INTE/projekt/player"
)

type Rouge struct {
	player.Player
	stealth int
}

const (
	baseStealth = 100
)

func New() (*Rouge, error) {
	return &Rouge{stealth: baseStealth}, nil
}
