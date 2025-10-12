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
	getBase() *BaseItem
	GetWearPosition() WearPosition
	GetType() string
	ToString() string
	GetName() string
}
type WearPosition string

type ItemType string

const (
	WEAR_POSITION_HEAD       WearPosition = "HEAD"
	WEAR_POSITION_UPPER_BODY WearPosition = "UPPER"
	WEAR_POSITION_LOWER_BODY WearPosition = "LOWER"
	WEAR_POSITION_FOOT       WearPosition = "FOOT"
	WEAR_POSITION_WEAPON     WearPosition = "WEAPON"
)

// func New(name string, weight int, itemType ItemType) (*BaseItem, error) {
// 	if len(name) == 0 {
// 		return nil, fmt.Errorf("No name supplied")
// 	}
// 	if weight < 0 {
// 		return nil, fmt.Errorf("Weight cannot be negative")
// 	}

//		return &BaseItem{id: random.String(), itemType: itemType, name: name}, nil
//	}
func New(item Item) Item {
	item.getBase().id = random.String()

	return item
}

func (c *BaseItem) GetID() string {
	return c.id
}

func (c *BaseItem) GetWeight() int {
	return c.weight
}

func (c *BaseItem) getBase() *BaseItem {
	return c
}
func (c *BaseItem) GetWearPosition() WearPosition {
	return WearPosition("")
}
func (c *BaseItem) GetType() string {
	return string(c.itemType)
}
func (c *BaseItem) GetName() string {
	return c.name
}
func (c *BaseItem) ToString() string {
	return fmt.Sprintf("Name: %s, Weight: %d", c.name, c.weight)
}
