package player

import (
	"INTE/projekt/character"
	"INTE/projekt/item"
	"fmt"
	"log"
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
	character.Fightable
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
		return newPaladin(player), nil
	case CLASS_ROGUE:
		return newRogue(player), nil
	case CLASS_MAGE:
		return newMage(player), nil
	}

	return nil, fmt.Errorf("No class provided")
}

func (p *BasePlayer) GetClass() Class {
	return p.class
}
func (p *BasePlayer) IsFightable() (fightable character.Fightable, ok bool) {
	return p, true
}

func (p *BasePlayer) GetMaxWeight() int {
	return p.maxWeight
}
func (p *BasePlayer) GetDamage() int {
	damage := 0
	left := p.GetItem(character.WEAR_POSITION_LEFT_ARM)
	right := p.GetItem(character.WEAR_POSITION_RIGHT_ARM)
	if !left.IsNothing() {
		weapon := left.(item.Weapon)
		damage += weapon.GetDamage()
	}
	if !right.IsNothing() {
		weapon := right.(item.Weapon)
		damage += weapon.GetDamage()
	}
	return damage
}
func (p *BasePlayer) Attack(rec character.Fightable) error {
	pFightable, ok := p.IsFightable()
	if !ok {
		return fmt.Errorf("Attacker can't fight")
	}
	eFightable, ok := rec.IsFightable()
	if !ok {
		return fmt.Errorf("Receiver can't fight")
	}
	log.Println(pFightable.GetDamage())
	log.Println(eFightable.GetHealth())

	return nil
}
