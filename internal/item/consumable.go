package item

type Consumable struct {
	Item
}

func (w *Consumable) GetEffect() string {
	return "Nothing ever happens"
}
