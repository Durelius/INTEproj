package item

import (
	"INTE/projekt/random"
	"fmt"
)

type BaseItem struct {
	id       string
	weight   int
	name     string
	itemType string
}

func New(name string, weight int, itemType string) (item, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name supplied")
	}

	return &BaseCharacter{id: random.String(), health: 100, name: name}, nil
}
