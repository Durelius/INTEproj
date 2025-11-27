package item

import "fmt"

type Weapon struct {
	damage int
	weight int
	name   string
	rarity Rarity
}

func (w *Weapon) GetDamage() int {
	return w.damage
}
func (w *Weapon) GetWeight() int {
	return w.weight
}

func (w *Weapon) GetRarity() Rarity {
	return w.rarity
}
func (w *Weapon) GetName() string {
	return w.name
}
func (w *Weapon) ToString() string {
	color := "\033[0m" // default (reset)

	switch w.rarity {
	case Common:
		color = "\033[32m" // green
	case Rare:
		color = "\033[34m" // blue
	case Epic:
		color = "\033[35m" // purple (magenta)
	case Legendary:
		color = "\033[31m" // red
	}

	return fmt.Sprintf("Name: %s%s\033[0m, Damage: %d, Weight: %d", color, w.name, w.GetDamage(), w.weight)
}
