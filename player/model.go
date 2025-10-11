package player

import (
	"INTE/projekt/character"
	"INTE/projekt/item"
	"fmt"
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
	PickupItem(item.Item) error
	GetTotalWeight() int
}
type Class string

const (
	CLASS_ROGUE        Class = "ROGUE"
	CLASS_PALADIN      Class = "PALADIN"
	CLASS_MAGE         Class = "MAGE"
	default_max_weight int   = 500
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
	curItem := p.GetItem(item.WEAR_POSITION_WEAPON)
	if len(curItem.GetID()) > 0 {
		weapon := curItem.(*item.Weapon)
		damage += weapon.GetDamage()
	}

	return damage
}
func (p *BasePlayer) ReceiveDamage(damage int) int {
	p.SetHealth(p.GetHealth() - damage)

	return p.GetHealth()
}
func (p *BasePlayer) Attack(rec character.Character) (int, error) {
	pFightable, ok := p.IsFightable()
	if !ok {
		return 0, fmt.Errorf("Attacker can't fight")
	}
	eFightable, ok := rec.IsFightable()
	if !ok {
		return 0, fmt.Errorf("Receiver can't fight")
	}

	return eFightable.ReceiveDamage(pFightable.GetDamage()), nil
}

func (p *BasePlayer) PickupItem(item item.Item) error {
	if p.GetTotalWeight()+item.GetWeight() > p.maxWeight {
		return fmt.Errorf("Overburdened")
	}
	p.Character.AddItemToBag(item)

	return nil
}
func (p *BasePlayer) GetTotalWeight() int {
	bag := p.GetBag()
	return bag.GetTotalWeight()
}
func (p *BasePlayer) GetType() string {
	return "PLAYER"
}
