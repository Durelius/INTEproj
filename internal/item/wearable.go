package item

import "fmt"

type Wearable struct {
	defense int
	slot    WearPosition
	weight  int
	name    string
	rarity  Rarity
}

func (w *Wearable) GetDefense() int {
	return w.defense
}

func (w *Wearable) GetWeight() int {
	return w.weight
}

func (w *Wearable) GetRarity() Rarity {
	return w.rarity
}
func (w *Wearable) GetName() string {
	return w.name
}

func (w *Wearable) GetSlot() WearPosition {
	return w.slot
}

func (w *Wearable) ToString() string {
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

	return fmt.Sprintf("Name: %s%s\033[0m, Defense: %d, Weight: %d", color, w.name, w.GetDefense(), w.weight)
}
