package item

import (
	"INTE/projekt/random"
	"fmt"
)

type BaseItem struct {
	id       string
	weight   int
	name     string
	itemType ItemType
}
type ItemType string

const (
	Stick  ItemType = "Stick"
	Food   ItemType = "Food"
	Potion ItemType = "Potion"
)

func New(name string, weight int, itemType ItemType) (*BaseItem, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}
	if weight < 0 {
		return nil, fmt.Errorf("Weight cannot be negative")
	}

	return &BaseItem{id: random.String(), itemType: itemType, name: name}, nil
}

func (c *BaseItem) GetID() string {
	return c.id
}

func (c *BaseItem) getWeight() int {
	return c.weight
}
