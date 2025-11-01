package class

type ClassName string

type Class interface {
	Name() ClassName
	GetDescription() string
	GetBaseDmg() int
	IncreaseStats(int)
	GetEnergy() int
}
