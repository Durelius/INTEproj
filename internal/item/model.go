package item

import (
	"fmt"
	"math/rand"

	"github.com/Durelius/INTEproj/internal/random"
)

// ------------------------------------------------ Seed -----------------------------------------------------
var globalRand *rand.Rand

func SetSeed(seed int64) {
	globalRand = rand.New(rand.NewSource(seed))
}

//------------------------------------------- Setters for shit -----------------------------------------------

func (c *BaseItem) GetID() string {
	return c.id
}
func (c *BaseItem) GetWeight() int {
	return c.weight
}
func (c *BaseItem) getBase() *BaseItem {
	return c
}
func (c *BaseItem) GetType() ItemType {
	return c.itemType
}
func (c *BaseItem) GetRarity() Rarity {
	return c.rarity
}
func (c *BaseItem) GetName() string {
	return c.name
}
func (c *BaseItem) ToString() string {
	return fmt.Sprintf("Name: %s, Weight: %d", c.name, c.weight)
}

// -------------------------------------------- Variables (ish) ------------------------------------------------------

type WearPosition string
type ItemType string
type BaseItem struct {
	id       string
	weight   int
	name     string
	itemType ItemType
	rarity   Rarity
}
type Item interface {
	GetID() string
	GetWeight() int
	getBase() *BaseItem
	GetType() ItemType
	ToString() string
	GetName() string
	GetRarity() Rarity
}

const (
	WEAR_POSITION_HEAD       WearPosition = "HEAD"
	WEAR_POSITION_UPPER_BODY WearPosition = "UPPER"
	WEAR_POSITION_LOWER_BODY WearPosition = "LOWER"
	WEAR_POSITION_FOOT       WearPosition = "FOOT"
	WEAR_POSITION_WEAPON     WearPosition = "WEAPON"
)

// ----------------------- RARITIES --------------------------

type Rarity int

const (
	Common Rarity = iota
	Rare
	Epic
	Legendary
)

var rarityPools = map[Rarity][]Item{
	Common:    COMMON_ITEMS,
	Rare:      RARE_ITEMS,
	Epic:      EPIC_ITEMS,
	Legendary: LEGENDARY_ITEMS,
}

// -------------------------------------------- Functions -------------------------------------------------

// make new item, idk if this is rly needed. But trqst the prqccess
func New(item Item) Item {
	item.getBase().id = random.String()

	return item
}

// Get a random item of the chosen rarity.
// Later todo, make it more dynamic for VAR
func GetRandomItemByRarity(r Rarity) Item {
	itemPool := rarityPools[r]
	return itemPool[globalRand.Intn(len(itemPool))]
}
func GetRandomItem() Item {
	selected := rand.Int63n(100) + 1
	var r Rarity
	switch {
	case selected <= 50:
		r = Common
	case selected <= 89:
		r = Rare
	case selected <= 98:
		r = Epic
	case selected <= 100:
		r = Legendary
	}
	return GetRandomItemByRarity(r)
}
