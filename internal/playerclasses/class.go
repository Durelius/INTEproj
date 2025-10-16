package playerclasses

type Class interface {
	Name() string
	GetDescription() string
	GetBaseDmg() int
}
