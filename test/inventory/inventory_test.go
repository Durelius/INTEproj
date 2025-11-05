package inventory_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player/inventory"
)

func TestNewEmptyInventory(t *testing.T) {
	inv := inventory.New()
	if len(inv.GetItems()) != 0 {
		t.Errorf("Inventory initialized without items should be empty")
	}

	sword := item.FindItemByName("Crimson Edge")

	inv.AddItem(sword)

	items  := inv.GetItems()
	if items[0] != sword || len(items) != 1 {
		t.Errorf("Expected inventory to have 1 item, Crimson Edge, found %s", items)
	}
}


func TestInventoryWithItems(t *testing.T) {
	sword := item.FindItemByName("Crimson Edge")
	inv := inventory.New(sword)

	items := inv.GetItems()
	if items[0] != sword || len(items) != 1 {
		t.Errorf("Expected inventory to have 1 item, Crimson Edge, found %s", items)
	}
}