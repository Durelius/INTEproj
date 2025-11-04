package player

import (
	"fmt"
	"math"

	"github.com/Durelius/INTEproj/internal/item"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/player/gear"
	"github.com/Durelius/INTEproj/internal/player/inventory"
	"github.com/Durelius/INTEproj/internal/random"
)

type Player struct {
	name          string
	class         class.Class
	inventory     *inventory.Inventory
	gear          *gear.Gear
	maxWeight     int
	maxHealth     int
	currentHealth int
	level         int
	experience    int
	dead          bool
	id            string
}

func New(name string, classParam class.ClassName) *Player {
	var playerClass class.Class
	switch classParam {
	case class.MAGE_STR:
		playerClass = class.NewMage()
	case class.PALADIN_STR:
		playerClass = class.NewPaladin()
	case class.ROGUE_STR:
		playerClass = class.NewRogue()
	}
	return &Player{
		name:          name,
		class:         playerClass,
		inventory:     inventory.New(),
		gear:          &gear.Gear{},
		maxWeight:     500,
		maxHealth:     100,
		currentHealth: 100,
		level:         1,
		experience:    0,
		id:            random.String(),
	}
}
func Load(name string, classParam class.ClassName, items []item.Item, gear gear.Gear, id string, maxWeight, maxHealth, currentHealth, level, experience, baseDmg, energy int) *Player {
	var playerClass class.Class
	switch classParam {
	case class.MAGE_STR:
		playerClass = class.LoadMage(baseDmg, energy)
	case class.PALADIN_STR:
		playerClass = class.LoadPaladin(baseDmg, energy)
	case class.ROGUE_STR:
		playerClass = class.LoadRogue(baseDmg, energy)
	}
	return &Player{
		name:          name,
		class:         playerClass,
		inventory:     inventory.New(items...),
		gear:          &gear,
		maxWeight:     maxWeight,
		maxHealth:     maxHealth,
		currentHealth: currentHealth,
		level:         level,
		experience:    experience,
		id:            id,
	}
}

// --------------------------------------------
// Getters
// --------------------------------------------

func (p *Player) GetName() string {
	return p.name
}
func (p *Player) GetID() string {
	return p.id
}

func (p *Player) GetClassName() class.ClassName {
	return p.class.Name()
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
	damageReduction := p.GetDamageReduction()

	p.currentHealth = p.currentHealth - int((float32(damage) * (1 - damageReduction)))

	if p.currentHealth <= 0 {
		p.currentHealth = 0
		p.dead = true
	}

	return p.currentHealth
}

// Damage reduction is given as a float, 0.25 = 25% damage reduction .
func (p *Player) GetDamageReduction() float32 {

	totalDefense := float32(p.GetTotalDefense())

	// This formula gives armor diminishing returns,
	// ie, 200 armor is not twice as good as 100 armor.
	// 0 armor = 100% damage taken, 50 armor = 67% damage taken, 400 armor = 20% damage taken
	return totalDefense / (totalDefense + 100)
}

func (p *Player) GetTotalDefense() int {
	g := p.gear
	itemSlots := []item.Item{g.Head, g.Upperbody, g.Legs, g.Feet}

	totalDefense := 0

	for _, slot := range itemSlots {
		if equippedItem, ok := slot.(*item.Wearable); ok {
			totalDefense += equippedItem.GetDefense()
		}
	}

	return totalDefense
}

// func (p *Player) CalculateDamageAfterDamageReduction(damage int) {
// 	damageReduction := p.GetDamageReduction()

// }

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
func (p *Player) levelUp() {
	p.level++                      // Increase level by 1
	p.class.IncreaseStats(p.level) // Increase class stats based on new level

	p.maxHealth += 20             // Increase max health
	p.currentHealth = p.maxHealth // Heal to full health on level up
}

func (p *Player) CalculateNextLevelExp() int {
	exp := 50 * math.Pow(1.5, float64(p.level))

	return int(math.Round(exp))
}

// --------------------------------------------
// Inventory and Gear methods
// --------------------------------------------

// PickupItem adds an item to the players inventory unless that would increase the total weight above maxWeight
func (p *Player) PickupItem(item item.Item) error {
	if p.GetTotalWeight()+item.GetWeight() > p.maxWeight {
		return fmt.Errorf("Overburdened")
	}
	p.inventory.AddItem(item)

	return nil
}

func (p *Player) DropItem(item item.Item) {
	p.inventory.RemoveItem(item)
}

// This method has an insane amount of duplication it is what it is
func (p *Player) EquipItem(i item.Item) {
	pGear := p.GetGear()

	// item is a weapon
	if w, ok := i.(*item.Weapon); ok {
		if pGear.Weapon == nil {
			pGear.Weapon = w
			p.inventory.RemoveItem(w)
		} else {
			p.inventory.RemoveItem(w)
			p.inventory.AddItem(pGear.Weapon)
			pGear.Weapon = w
		}
	}
	// item is a piece of armor
	if w, ok := i.(*item.Wearable); ok {
		switch w.GetSlot() {
		case item.WEAR_POSITION_HEAD:
			if pGear.Head == nil {
				pGear.Head = w
				p.inventory.RemoveItem(w)
			} else {
				p.inventory.RemoveItem(w)
				p.inventory.AddItem(pGear.Head)
				pGear.Head = w
			}
		case item.WEAR_POSITION_UPPER_BODY:
			if pGear.Upperbody == nil {
				pGear.Upperbody = w
				p.inventory.RemoveItem(w)
			} else {
				p.inventory.RemoveItem(w)
				p.inventory.AddItem(pGear.Upperbody)
				pGear.Upperbody = w
			}
		case item.WEAR_POSITION_LOWER_BODY:
			if pGear.Legs == nil {
				pGear.Legs = w
				p.inventory.RemoveItem(w)
			} else {
				p.inventory.RemoveItem(w)
				p.inventory.AddItem(pGear.Legs)
				pGear.Legs = w
			}
		case item.WEAR_POSITION_FOOT:
			if pGear.Feet == nil {
				pGear.Feet = w
				p.inventory.RemoveItem(w)
			} else {
				p.inventory.RemoveItem(w)
				p.inventory.AddItem(pGear.Feet)
				pGear.Feet = w
			}
		}
	}
}

func (p *Player) UnequipHead() bool {
	if p.gear.Head != nil {
		p.inventory.AddItem(p.gear.Head)
		p.gear.Head = nil
		return true
	}
	return false
}

func (p *Player) UnequipUpperBody() bool {
	if p.gear.Upperbody != nil {
		p.inventory.AddItem(p.gear.Upperbody)
		p.gear.Upperbody = nil
		return true
	}
	return false
}

func (p *Player) UnequipLowerBody() bool {
	if p.gear.Legs != nil {
		p.inventory.AddItem(p.gear.Legs)
		p.gear.Legs = nil
		return true
	}
	return false
}

func (p *Player) UnequipFeet() bool {
	if p.gear.Feet != nil {
		p.inventory.AddItem(p.gear.Feet)
		p.gear.Feet = nil
		return true
	}
	return false
}

func (p *Player) UnequipWeapon() bool {
	if p.gear.Weapon != nil {
		p.inventory.AddItem(p.gear.Weapon)
		p.gear.Weapon = nil
		return true
	}
	return false
}
