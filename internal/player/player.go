package player

import (
	"fmt"
	"math"

	"github.com/Durelius/INTEproj/internal/item"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/player/gear"
	"github.com/Durelius/INTEproj/internal/player/inventory"
)

type Player struct {
	name string
	class     class.Class
	inventory *inventory.Inventory
	gear    *gear.Gear
	maxWeight int
	maxHealth int
	currentHealth int
	level      int
	experience int
	dead	  bool
}

func New(name string, class class.Class) *Player {
	return &Player{
		name: name,
		class: class, 
		inventory: inventory.New(),
		gear: &gear.Gear{},
		maxWeight: 500, 
		maxHealth: 100,
		currentHealth: 100,
		level: 1,
		experience: 0,
	}
}

// --------------------------------------------
// Getters
// --------------------------------------------

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetClass() class.Class {
	return p.class
}

func (p *Player) GetLevel() int {
	return p.level
}

func (p *Player) GetExperience() int {
	return p.experience
}

func (p *Player) GetItems() []item.Item {
	return p.inventory.GetItems()
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

func (p *Player) GetGear() *gear.Gear {
	return p.gear
}

func (p *Player) IsDead() bool {
	return p.dead
}

// Returns the total weight of all items in inventory and gear slots
func (p *Player) GetTotalWeight() int {
	invWeight := p.inventory.GetTotalWeight()

	return invWeight + p.gear.GetTotalWeight()
}

func (p *Player) GetInventoryWeight() int {
	return p.inventory.GetTotalWeight()
}

func (p *Player) GetEquippedWeight() int {
	return p.gear.GetTotalWeight()
}

// GetDamage returns the total damage of the player including weapon damage
func (p *Player) GetDamage() int {
	damage := p.class.GetBaseDmg()

	w := p.gear.Weapon
	if w != nil {
		weapon := w.(*item.Weapon)
		damage += weapon.GetDamage()
	}

	return damage
}


// --------------------------------------------
// Combat methods
// --------------------------------------------

// Reduce current health
func (p *Player) ReceiveDamage(damage int) int {
	p.currentHealth = p.currentHealth - damage

	if p.currentHealth <= 0 {
		p.currentHealth = 0
		p.dead = true
	}

	return p.currentHealth
}

// --------------------------------------------
// Experience and Leveling methods
// --------------------------------------------

func (p *Player) IncreaseExperience(exp int) {
	p.experience += exp

	for p.experience >= p.CalculateNextLevelExp() {
		p.experience -= p.CalculateNextLevelExp()
		p.levelUp()
	}
}

// Handles all logic for leveling up a player
func (p *Player) levelUp () {
	p.level++	// Increase level by 1
	p.class.IncreaseStats(p.level)	// Increase class stats based on new level

	p.maxHealth += 20 // Increase max health
	p.currentHealth = p.maxHealth 	// Heal to full health on level up
}

func (p *Player) CalculateNextLevelExp() int {
	exp := 50 * math.Pow(1.5,float64(p.level))

	return int(math.Round(exp))
}


// --------------------------------------------
// Inventory methods
// --------------------------------------------

// PickupItem adds an item to the players inventory unless that would increase the total weight above maxWeight
func (p *Player) PickupItem(item item.Item) error {
	if p.GetTotalWeight()+item.GetWeight() > p.maxWeight {
		return fmt.Errorf("Overburdened")
	}
	p.inventory.AddItem(item)

	return nil
}
