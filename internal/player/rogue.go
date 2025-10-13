package player

type Rogue struct {
	Player
	stealth int
}

const (
	baseStealth = 100
)

func newRogue(player Player) Player {
	return &Rogue{stealth: baseStealth, Player: player}
}
