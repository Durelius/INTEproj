package inventory_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player/inventory"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestNewEmptyInventory(t *testing.T) {
	g := gomega.NewWithT(t)

	inv := inventory.New()

	g.Expect(len(inv.GetItems())).To(Equal(0))

	sword := item.FindItemByName("Crimson Edge")
	inv.AddItem(sword)

	g.Expect(len(inv.GetItems())).To(Equal(1))
	g.Expect(inv.GetItems()[0]).To(Equal(sword))
}


func TestInventoryWithItems(t *testing.T) {
	g := gomega.NewWithT(t)

	sword := item.FindItemByName("Crimson Edge")
	inv := inventory.New(sword)
	items := inv.GetItems()

	g.Expect(items).To(ContainElement(sword))
	g.Expect(items).To(HaveLen(1))
}











