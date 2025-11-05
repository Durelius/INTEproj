package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
)

// TODO TestCreateAllItemTypes kontrollerar att varje item-typ blir korrekt skapad
func TestCreateAllItemTypes(t *testing.T) {
	// tests := []struct {
	// 	name     string
	// 	itemType item.ItemType
	// }{
	// 	{"consumable", item.itemType.consumable},
	// 	{"weapon", item.ItemType.weapon},
	// 	{"wearable", item.ItemType.wearable},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		/*if tt.item.item == nil {
	// 			t.Errorf("%s has nil itemType", tt.name)
	// 		} not sure why no work. */
	// 		if tt.itemT.GetType() != tt.itemType {
	// 			t.Errorf("%s has wrong itemtype: got %v want %v", tt.name, tt.item.GetType(), tt.itemType)
	// 		}
	// 	})
	// }
}

// TestCreateAllRarities kontrollerar att varje rarity-item är korrekt skapad
func TestCreateAllRarities(t *testing.T) {
	tests := []struct {
		name   string
		rarity item.Rarity
	}{
		{"Rusty Sword", item.Common},
		{"Elven Longbow", item.Rare},
		{"Runeblade", item.Epic},
		{"Eclipse Fang", item.Legendary},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := item.FindItemByName(tt.name)
			if got == nil {
				t.Errorf("%s has nil BaseItem", tt.name)
			}
			if got.GetRarity() != tt.rarity {
				t.Errorf("%s has wrong rarity: got %v, want %v", tt.name, got.GetRarity(), tt.rarity)
			}
		})
	}
}

// Testar att GetRandomItemByRarity() fungerar korrekt.
func TestGetRandomItemByRarity(t *testing.T) {

	for _, rarity := range []item.Rarity{item.Common, item.Rare, item.Epic, item.Legendary} {
		got := item.GetRandomItemByRarity(rarity)
		if got == nil {
			t.Errorf("Expected item for rarity %v, got nil", rarity)
			continue
		}
		if got.GetRarity() != rarity {
			t.Errorf("Expected rarity %v, got %v for item %v", rarity, got.GetRarity(), got.GetName())
		}
	}
}

// Testar random delningen mellan alla olika rareities. Kör testet individuellt för logging
// Ger aldrig rätt eller fel, men ger en lista som visar division av drop rates
// Common 50 %
// Rare 40 %
// Epic 8 %
// Legendary 2%
func TestGetRandomItemsWith(t *testing.T) {
	length := 1000
	items := []item.Item{}
	for i := 0; i < length; i++ {
		items = append(items, item.GetRandomItem())
	}

	counts := make(map[item.Rarity]int)
	for _, it := range items {
		counts[it.GetRarity()]++
	}

	/*for rarity, count := range counts {
		t.Logf("%v: %d", rarity, count)
	}*/
}
