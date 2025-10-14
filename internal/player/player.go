package player

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/character"
	"github.com/Durelius/INTEproj/internal/inventory"
	"github.com/Durelius/INTEproj/internal/item"
	class "github.com/Durelius/INTEproj/internal/playerclasses"
)

type Gear struct {
	head item.Item
	upperbody item.Item
	legs item.Item
	weapon item.Item
	feet item.Item
}

type Player struct {
	name string
	class     class.Class
	inventory *inventory.Inventory
	maxWeight int
	maxHealth int
	currentHealth int
	gear    *Gear
}

func New(name string, class class.Class) *Player {
	return &Player{
		name: name,
		class: class, 
		maxWeight: 500, 
		inventory: inventory.New(),
		maxHealth: 100,
		currentHealth: 100,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetClass() *class.Class {
	return &p.class
}

func (p *Player) GetInventory() *inventory.Inventory {
	return p.inventory
}

func (p *Player) GetMaxWeight() int {
	return p.maxWeight
}

func (p *Player) GetCurrentHealth() int {
	return p.currentHealth
}

func (p *Player) GetMaxHealth() int {
	return p.maxHealth
}

func (p *Player) GetTotalWeight() int {
	return p.inventory.GetTotalWeight()
}

func (p *Player) GetDamage() int {
	damage := p.class.GetBaseDmg()
	curItem := p.gear.weapon
	if len(curItem.GetID()) > 0 {	// Do we need to check this?
		weapon := curItem.(*item.Weapon)
		damage += weapon.GetDamage()
	}

	return damage
}

func (p *Player) ReceiveDamage(damage int) int {
	p.currentHealth = p.currentHealth - damage
	return p.currentHealth
}

func (p *Player) Attack(rec character.Character) (int, error) {
	eFightable, ok := rec.IsFightable()
	if !ok {
		return 0, fmt.Errorf("Receiver can't fight")
	}

	return eFightable.ReceiveDamage(p.GetDamage()), nil
}

func (p *Player) PickupItem(item item.Item) error {
	if p.GetTotalWeight()+item.GetWeight() > p.maxWeight {
		return fmt.Errorf("Overburdened")
	}
	p.inventory.AddItem(item)

	return nil
}
