package item

import "fmt"

type Weapon struct {
	damage   int
	weight   int
	name     string
	itemType ItemType
	rarity   Rarity
}

func (w *Weapon) GetDamage() int {
	return w.damage
}
func (w *Weapon) GetWeight() int {
	return w.weight
}
func (w *Weapon) GetType() ItemType {
	return w.itemType
}
func (w *Weapon) GetRarity() Rarity {
	return w.rarity
}
func (w *Weapon) GetName() string {
	return w.name
}
func (w *Weapon) ToString() string {
	return fmt.Sprintf("Name: %s, Weight: %d", w.name, w.weight)
}
