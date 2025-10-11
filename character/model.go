package character

import (
	"INTE/projekt/bag"
	"INTE/projekt/item"
	"INTE/projekt/random"
	"fmt"
)

type BaseCharacter struct {
	id       string
	health   int
	name     string
	equipped map[item.WearPosition]item.Item
	bag      *bag.Bag
}
type Character interface {
	GetID() string
	GetHealth() int
	GetName() string
	SetEquippedItem(item.WearPosition, item.Item)
	GetItem(item.WearPosition) item.Item
	IsFightable() (Fightable, bool)
	SetHealth(int)
	IsAlive() bool
	GetBag() bag.Bag
	AddItemToBag(item.Item)
}
type Fightable interface {
	Attack(rec Character) (int, error)
	GetDamage() int
	ReceiveDamage(int) int
	Character
}

func New(name string) (Character, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}
	char := &BaseCharacter{id: random.String(), health: 100, name: name}

	char.initializeItems()
	return char, nil
}

func (c *BaseCharacter) GetID() string {
	return c.id
}

func (c *BaseCharacter) GetHealth() int {
	return c.health
}
func (c *BaseCharacter) GetName() string {
	return c.name
}
func (c *BaseCharacter) SetHealth(health int) {
	c.health = health
}
func (c *BaseCharacter) IsFightable() (Fightable, bool) {
	return nil, false
}
func (c *BaseCharacter) IsAlive() bool {
	return c.health > 0
}
func (c *BaseCharacter) initializeItems() {
	c.bag = bag.New()
	c.equipped = make(map[item.WearPosition]item.Item)
	c.equipped[item.WEAR_POSITION_HEAD] = item.NOTHING
	c.equipped[item.WEAR_POSITION_UPPER_BODY] = item.NOTHING
	c.equipped[item.WEAR_POSITION_LOWER_BODY] = item.NOTHING
	c.equipped[item.WEAR_POSITION_FOOT] = item.NOTHING
	c.equipped[item.WEAR_POSITION_WEAPON] = item.NOTHING
}

func (c *BaseCharacter) SetEquippedItem(wp item.WearPosition, item item.Item) {
	c.equipped[wp] = item
}
func (c *BaseCharacter) GetItem(wp item.WearPosition) item.Item {
	return c.equipped[wp]
}
func (c *BaseCharacter) GetBag() bag.Bag {
	return *c.bag
}
func (c *BaseCharacter) AddItemToBag(item item.Item) {
	c.bag.AddItem(item)
}
