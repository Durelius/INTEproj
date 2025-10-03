package player

import (
	"INTE/projekt/character"
)

type BasePlayer struct {
	character.Character
	class     Class
	maxWeight int
}
type Player interface {
	GetClass() Class
	GetMaxWeight() int
	character.Character
}
type Class string

const (
	CLASS_ROUGE        Class = "ROUGE"
	CLASS_PALADIN      Class = "PALADIN"
	CLASS_MAGE         Class = "MAGE"
	default_max_weight int   = 50
)

func New(class Class, name string) (Player, error) {
	char, err := character.New(name)
	if err != nil {
		return nil, err
	}
	return &BasePlayer{class: class, maxWeight: default_max_weight, Character: char}, nil
}

func (p *BasePlayer) GetClass() Class {
	return p.class
}

func (p *BasePlayer) GetMaxWeight() int {
	return p.maxWeight
}
