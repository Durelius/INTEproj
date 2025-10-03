package character

import (
	"INTE/projekt/random"
	"fmt"
)

type BaseCharacter struct {
	id     string
	health int
	name   string
}
type Character interface {
	GetID() string
	GetHealth() int
	GetName() string
}

func New(name string) (Character, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}

	return &BaseCharacter{id: random.String(), health: 100, name: name}, nil
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
