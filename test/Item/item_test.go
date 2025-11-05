package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/onsi/gomega"
)

// Testar random delningen mellan alla olika rareities.
// Common 50 %
// Rare 40 %
// Epic 8 %
// Legendary 2%
func TestGetRandomItemsWith(t *testing.T) {
	g := gomega.NewWithT(t)
	const length = 1000
	items := make([]item.Item, 0, length)

	for i := 0; i < length; i++ {
		items = append(items, item.GetRandomItem())
	}

	counts := make(map[item.Rarity]int)
	for _, it := range items {
		counts[it.GetRarity()]++
	}

	// Expected proportions
	g.Expect(counts[item.Common]).To(gomega.BeNumerically(">=", 460))
	g.Expect(counts[item.Common]).To(gomega.BeNumerically("<=", 540))

	g.Expect(counts[item.Rare]).To(gomega.BeNumerically(">=", 360))
	g.Expect(counts[item.Rare]).To(gomega.BeNumerically("<=", 440))

	g.Expect(counts[item.Epic]).To(gomega.BeNumerically(">=", 60))
	g.Expect(counts[item.Epic]).To(gomega.BeNumerically("<=", 90))

	g.Expect(counts[item.Legendary]).To(gomega.BeNumerically(">=", 5))
	g.Expect(counts[item.Legendary]).To(gomega.BeNumerically("<=", 35))
}
