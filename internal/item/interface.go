package item

import (
	"strings"
	"sync"

	"github.com/Durelius/INTEproj/internal/random"
)

// ------------------------------------------------ Seed -----------------------------------------------------

// -------------------------------------------- Variables (ish) ------------------------------------------------------

type ItemType string

type Item interface {
	GetWeight() int
	GetType() ItemType
	ToString() string
	GetName() string
	GetRarity() Rarity
}

type WearPosition int

const (
	WEAR_POSITION_HEAD = iota
	WEAR_POSITION_UPPER_BODY
	WEAR_POSITION_LOWER_BODY
	WEAR_POSITION_FOOT
)

// ----------------------- RARITIES --------------------------

type Rarity int

const (
	Common Rarity = iota
	Rare
	Epic
	Legendary
)

var rarityIndex = map[Rarity][]Item{}
var rarityMu sync.RWMutex //Används för att se till att rarityindex alltid initierar korrekt

// On init så delar den upp rarities i rarityindex. Snabbare sökning eftersom vi ofta kallar den
func init() {
	rarityMu.Lock()
	defer rarityMu.Unlock()

	for _, item := range AllItems {
		rarityIndex[item.GetRarity()] = append(rarityIndex[item.GetRarity()], item)
	}
}

// -------------------------------------------- Functions -------------------------------------------------

// Get a random item of the chosen rarity.
// Later todo, make it more dynamic for VAR
func GetRandomItemByRarity(r Rarity) Item {
	rarityMu.RLock()
	defer rarityMu.RUnlock()

	pool := rarityIndex[r]
	if len(pool) > 0 {
		return pool[random.IntList(len(pool))]
	} else {
		return nil
	}

}

func GetRandomItem() Item {

	selected := random.Int(0, 100)
	var r Rarity
	switch {
	case selected <= 50:
		r = Common
	case selected <= 90:
		r = Rare
	case selected <= 98:
		r = Epic
	case selected <= 100:
		r = Legendary
	}
	return GetRandomItemByRarity(r)
}

func FindItemByName(name string) Item {
	for _, item := range AllItems {
		if strings.EqualFold(item.GetName(), name) {
			return item
		}
	}
	return nil
}
