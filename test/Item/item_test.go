package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
	"github.com/onsi/gomega"
)

func TestRandomItemsRate(t *testing.T) {
	getRandomItemsWith(int64(5), t)
	getRandomItemsWith(int64(6), t)
	getRandomItemsWith(int64(7), t)
	getRandomItemsWith(int64(8), t)
	getRandomItemsWith(int64(9), t)
}

// Testar random delningen mellan alla olika rareities.
// Common 50 %
// Rare 40 %
// Epic 8 %
// Legendary 2%

func getRandomItemsWith(seed int64, t *testing.T) {
	g := gomega.NewWithT(t)
	random.SetSeed(seed)
	length := 1000
	items := []item.Item{}
	for i := 0; i < length; i++ {
		items = append(items, item.GetRandomItem())
	}

	counts := make(map[item.Rarity]int)
	for _, it := range items {
		counts[it.GetRarity()]++
	}

	g.Expect(counts[item.Common]).To(gomega.BeNumerically(">=", 470))
	g.Expect(counts[item.Common]).To(gomega.BeNumerically("<=", 530))

	g.Expect(counts[item.Rare]).To(gomega.BeNumerically(">=", 370))
	g.Expect(counts[item.Rare]).To(gomega.BeNumerically("<=", 430))

	g.Expect(counts[item.Epic]).To(gomega.BeNumerically(">=", 70))
	g.Expect(counts[item.Epic]).To(gomega.BeNumerically("<=", 90))

	g.Expect(counts[item.Legendary]).To(gomega.BeNumerically(">=", 15))
	g.Expect(counts[item.Legendary]).To(gomega.BeNumerically("<=", 25))

}
