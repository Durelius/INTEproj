package player

type Rogue struct {
	Player
	stealth int
}

const (
	baseStealth = 100
)

func newRogue() *Rogue {
	return &Rogue{stealth: baseStealth}
}
