package item

type Head struct {
	Item
	defense int
}

func (w *Head) GetDefense() int {
	return w.defense
}

type UpperBody struct {
	Item
	defense int
}

func (w *UpperBody) GetDefense() int {
	return w.defense
}

type LowerBody struct {
	Item
	defense int
}

func (w *LowerBody) GetDefense() int {
	return w.defense
}

type Foot struct {
	Item
	defense int
}

func (w *Foot) GetDefense() int {
	return w.defense
}
