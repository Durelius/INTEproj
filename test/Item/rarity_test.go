package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/onsi/gomega"
)

// Check Rarity IOTA works correctly
func TestCreateOutOfBoundsIota(t *testing.T) {
	outOfIotaRange := 4
	g := gomega.NewWithT(t)
	g.Expect(item.GetRandomItemByRarity(item.Rarity(outOfIotaRange))).To(gomega.BeNil())
}

// Checks that all normal rarities can be created
func TestGetRandomItemByRarity(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	for _, rarity := range []item.Rarity{item.Common, item.Rare, item.Epic, item.Legendary} {
		got := item.GetRandomItemByRarity(rarity)
		g.Expect(got).NotTo(gomega.BeNil())
		g.Expect(got.GetRarity()).To(gomega.Equal(rarity))
	}
}
