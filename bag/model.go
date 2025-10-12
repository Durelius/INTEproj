package bag

import (
	"INTE/projekt/item"
	"fmt"
)

type Bag struct {
	items []item.Item
}

func New() *Bag {
	return &Bag{items: []item.Item{}}
}

func (b *Bag) AddItem(item item.Item) {
	b.items = append(b.items, item)
}

func (b *Bag) RemoveItem(item item.Item) error {
	for i, x := range b.items {
		if x.GetID() == item.GetID() {
			b.items[i] = b.items[len(b.items)-1]
			b.items = b.items[:len(b.items)-1]
			return nil
		}
	}
	return fmt.Errorf("This item didn't exist in bag")
}
func (b *Bag) GetTotalWeight() (weight int) {
	for _, item := range b.items {
		weight += item.GetWeight()
	}
	return weight
}
func (b *Bag) GetItems() []item.Item {
	return b.items
}
