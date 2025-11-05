package itemtest

import (
	"strings"
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
)

// Create named items of all rarities. Testcase 3 from tillståndsmaskin
func TestGetItemsByName(t *testing.T) {
	str := []string{"Rusty Sword", "Training Sword", "Soulfire Edge", "Crown of Thorns", "Crown of Eternity"}
	count := len(str)

	for i := 0; i < count; i++ {
		item := item.FindItemByName(str[i])
		if item == nil {
			t.Errorf("Item %q not found", str[i])
			continue
		}
		if !strings.EqualFold(str[i], item.GetName()) {
			t.Errorf("Expected name %q, got %q", str[i], item.GetName())
		}
	}
}

// Create empty item som inte finns
func TestCreateEmptyItem(t *testing.T) {
	str := []string{"", "NotItemName", "     ", "nil", "Ru57y 5w0RD"}
	count := len(str)

	for i := 0; i < count; i++ {
		item := item.FindItemByName(str[i])
		if item == nil {
			continue
		} else {
			t.Errorf("Expected name %q, got %q", str[i], item.GetName())
		}

	}
}
func TestCreateNilItem(t *testing.T) {
	var nilString string
	item := item.FindItemByName(nilString)
	if item != nil {
		t.Errorf("item %q should not have been created from nil", item.GetName())
	}

}
