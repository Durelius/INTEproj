package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
)

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
