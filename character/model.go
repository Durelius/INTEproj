package character

import (
	"INTE/projekt/item"
	"INTE/projekt/random"
	"fmt"
)

type BaseCharacter struct {
	id       string
	health   int
	name     string
	equipped map[WearPosition]item.Item
}
type Character interface {
	GetID() string
	GetHealth() int
	GetName() string
}
type Fightable interface {
	Attack(rec Fightable)
}
type WearPosition string

const (
	WEAR_POSITION_HEAD       WearPosition = "HEAD"
	WEAR_POSITION_UPPER_BODY WearPosition = "UPPER"
	WEAR_POSITION_LOWER_BODY WearPosition = "LOWER"
	WEAR_POSITION_FOOT       WearPosition = "FOOT"
	WEAR_POSITION_LEFT_ARM   WearPosition = "LEFT"
	WEAR_POSITION_RIGHT_ARM  WearPosition = "RIGHT"
)

func New(name string) (Character, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}
	char := &BaseCharacter{id: random.String(), health: 100, name: name}

	char.equipped = make(map[WearPosition]item.Item)

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
