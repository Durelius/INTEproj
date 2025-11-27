package itemtest

import (
	"fmt"
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

func TestWeaponToString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	tests := []struct {
		name   string
		damage int
		weight int
	}{
		{"Rusty Sword", 10, 30},
		{"Steel Sword", 25, 35},
		{"Dragonbone Axe", 45, 50},
		{"Soulfire Edge", 63, 45},
	}

	for _, tt := range tests {
		item := item.GetItemByName(tt.name)

		g.Expect(item).NotTo(gomega.BeNil())
		out := item.ToString()

		g.Expect(out).To(gomega.ContainSubstring("Name:"))
		g.Expect(out).To(gomega.ContainSubstring(tt.name))
		g.Expect(out).To(gomega.ContainSubstring(fmt.Sprintf("Damage: %d", tt.damage)))
		g.Expect(out).To(gomega.ContainSubstring(fmt.Sprintf("Weight: %d", tt.weight)))
	}
}
func TestWearableToString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	tests := []struct {
		name    string
		defense int
		weight  int
	}{
		{"Three Stripes", 13, 10},
		{"Steel Helm", 25, 35},
		{"Helm of Embersteel", 42, 40},
		{"Crown of Eternity", 65, 45},
	}

	for _, tt := range tests {
		it := item.GetItemByName(tt.name)
		g.Expect(it).NotTo(gomega.BeNil())

		out := it.ToString()

		// These always exist regardless of ANSI
		g.Expect(out).To(gomega.ContainSubstring("Name:"))
		g.Expect(out).To(gomega.ContainSubstring(tt.name))
		g.Expect(out).To(gomega.ContainSubstring(fmt.Sprintf("Defense: %d", tt.defense)))
		g.Expect(out).To(gomega.ContainSubstring(fmt.Sprintf("Weight: %d", tt.weight)))
	}
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
