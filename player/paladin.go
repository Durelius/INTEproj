package player

type Paladin struct {
	Player
	holiness int
}

const (
	base_holiness = 100
)

func newPaladin() *Paladin {

	return &Paladin{holiness: base_holiness}
}
