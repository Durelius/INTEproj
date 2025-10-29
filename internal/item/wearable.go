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
	color := "\033[0m" // default (reset)

	if w.rarity == Common {
		color = "\033[32m" // green
	} else if w.rarity == Rare {
		color = "\033[34m" // blue
	} else if w.rarity == Epic {
		color = "\033[35m" // purple (magenta)
	} else if w.rarity == Legendary {
		color = "\033[31m" // red
	}

	return fmt.Sprintf("Name: %s%s\033[0m, Weight: %d", color, w.name, w.weight)
}
