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
type Item interface {
	GetID() string
	GetWeight() int
}
type ItemType string

const (
	WEAPON     ItemType = "WEAPON"
	CONSUMABLE ItemType = "CONSUMABLE"
	WEARABLE   ItemType = "WEARABLE"
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

func (c *BaseItem) GetWeight() int {
	return c.weight
}
