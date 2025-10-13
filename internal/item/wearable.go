package item

type Head struct {
	Item
	defense int
}

func (w *Head) GetDefense() int {
	return w.defense
}
func (w *Head) GetWearPosition() WearPosition {
	return WEAR_POSITION_HEAD
}

type UpperBody struct {
	Item
	defense int
}

func (w *UpperBody) GetDefense() int {
	return w.defense
}
func (w *UpperBody) GetWearPosition() WearPosition {
	return WEAR_POSITION_UPPER_BODY
}

type LowerBody struct {
	Item
	defense int
}

func (w *LowerBody) GetDefense() int {
	return w.defense
}
func (w *LowerBody) GetWearPosition() WearPosition {
	return WEAR_POSITION_LOWER_BODY
}

type Foot struct {
	Item
	defense int
}

func (w *Foot) GetDefense() int {
	return w.defense
}
func (w *Foot) GetWearPosition() WearPosition {
	return WEAR_POSITION_FOOT
}
