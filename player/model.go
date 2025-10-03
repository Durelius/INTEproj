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
}
type Class string

const (
	class_rouge        Class = "ROUGE"
	class_paladin      Class = "PALADIN"
	class_mage         Class = "MAGE"
	default_max_weight int   = 50
)

func New(class Class) (Player, error) {
	return &BasePlayer{class: class, maxWeight: default_max_weight}, nil
}

func (p *BasePlayer) GetClass() Class {
	return p.class
}

func (p *BasePlayer) GetMaxWeight() int {
	return p.maxWeight
}
