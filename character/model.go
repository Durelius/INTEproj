package character

import (
	"INTE/projekt/random"
	"fmt"
)

type Character struct {
	id     string
	health int
	name   string
}

const id_length = 16

func New(name string) (*Character, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}

	return &Character{id: random.String(id_length), health: 100, name: name}, nil
}

func (c *Character) GetID() string {
	return c.id
}

func (c *Character) GetHealth() int {
	return c.health
}
func (c *Character) GetName() string {
	return c.name
}
