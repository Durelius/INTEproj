package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
)

// Full Random Item
func TestRandomItemsCreation(t *testing.T) {
	for i := 0; i < 10; i++ {
		v := item.GetRandomItem()

		if v == nil {
			t.Errorf("expected non-nil Item, got nil")
		}
	}
}

// Check Rarity IOTA works correctly
func TestCreateOutOfBoundsIota(t *testing.T) {
	for i := 4; i < 10; i++ {
		v := item.GetRandomItemByRarity(item.Rarity(i))
		if v != nil {
			t.Errorf("expected nil, got %T", v)
		}
	}
}

// Checks that all normal rarities can be created
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
