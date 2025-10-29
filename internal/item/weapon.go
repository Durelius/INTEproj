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
	color := "\033[0m" // default (reset)

	if w.rarity == Common {
		color = "\033[32m" // green
	} else if w.rarity == Rare {
		color = "\033[34m" // blue
	} else if w.rarity == Epic {
		color = "\033[33m" // yellow/orange
	} else if w.rarity == Legendary {
		color = "\033[93m" // gold-yellow
	}

	return fmt.Sprintf("%sName: %s, Weight: %d\033[0m", color, w.name, w.weight)
}
