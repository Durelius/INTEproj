package itemtest

import (
	"strconv"
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
)

// Hjälpfunktion
func createItemWithSeed(seed int64) item.Item {
	random.SetSeed(seed)
	return item.GetRandomItem()

}

// Testar items med samma seed om det ger samma item
func TestGetRandomItem(t *testing.T) {
	for i := 1; i < 20; i++ {
		item1 := createItemWithSeed(int64(i))
		item2 := createItemWithSeed(int64(i))

		if item1 != item2 {
			t.Error("Expected that seed: ", strconv.Itoa(i), " gave same item twice. It gave 1: ", item1.GetName(), ". And 2:", item2.GetName(), ". Which is not same item")
		} else {
			//t.Logf(item1.GetName(), " == ", item2.GetName()) Debugging
		}
	}
}
