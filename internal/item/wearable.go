package item

import "fmt"

type Wearable struct {
	defense  int
	slot     WearPosition
	weight   int
	name     string
	itemType ItemType
	rarity   Rarity
}

func (w *Wearable) GetDefense() int {
	return w.defense
}

func (w *Wearable) GetWeight() int {
	return w.weight
}
func (w *Wearable) GetType() ItemType {
	return w.itemType
}
func (w *Wearable) GetRarity() Rarity {
	return w.rarity
}
func (w *Wearable) GetName() string {
	return w.name
}
func (w *Wearable) ToString() string {
	return fmt.Sprintf("Name: %s, Weight: %d", w.name, w.weight)
}
