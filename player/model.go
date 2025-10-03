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
	CLASS_ROGUE        Class = "ROGUE"
	CLASS_PALADIN      Class = "PALADIN"
	CLASS_MAGE         Class = "MAGE"
	default_max_weight int   = 50
)

func New(class Class, name string) (Player, error) {
	char, err := character.New(name)
	if err != nil {
		return nil, err
	}
	player := &BasePlayer{class: class, maxWeight: default_max_weight, Character: char}
	switch class {
	case CLASS_PALADIN:
		p := newPaladin()
		p.Player = player
		return p, nil
	case CLASS_ROGUE:
		p := newRogue()
		p.Player = player
		return p, nil
	case CLASS_MAGE:
		p := newMage()
		p.Player = player
		return p, nil
	}

	return &BasePlayer{class: class, maxWeight: default_max_weight, Character: char}, nil
}

func (p *BasePlayer) GetClass() Class {
	return p.class
}

func (p *BasePlayer) GetMaxWeight() int {
	return p.maxWeight
}
func (p *BasePlayer) Fight(rec character.Fightable) {

}
