package item

type Wearable struct {
	Item
	defense int
}

func (w *Wearable) GetDefense() int {
	return w.defense
}
