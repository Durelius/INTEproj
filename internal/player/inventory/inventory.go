package inventory

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/item"
)

type Inventory struct {
	items []item.Item
}

func New() *Inventory {
	return &Inventory{items: []item.Item{}}
}

func (inv *Inventory) AddItem(item item.Item) {
	inv.items = append(inv.items, item)
}

func (inv *Inventory) RemoveItem(item item.Item) error {
	for i, x := range inv.items {
		if x.GetName() == item.GetName() {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("This item didn't exist in inventory")
}

func (inv *Inventory) GetTotalWeight() (weight int) {
	for _, item := range inv.items {
		weight += item.GetWeight()
	}
	return weight
}

func (inv *Inventory) GetItems() []item.Item {
	return inv.items
}
