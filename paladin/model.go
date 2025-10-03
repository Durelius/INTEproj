package rouge

import (
	"INTE/projekt/player"
)

type Paladin struct {
	player.Player
	holiness int
}

const (
	base_holiness = 100
)

func New() (*Paladin, error) {
	return &Paladin{holiness: base_holiness}, nil
}
