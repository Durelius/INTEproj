package class

type Class interface {
	Name() string
	GetDescription() string
	GetBaseDmg() int
	IncreaseStats(int)
}
