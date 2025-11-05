package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
	"github.com/onsi/gomega"
)

// Hjälpfunktion
func createItemWithSeed(seed int64) item.Item {
	random.SetSeed(seed)
	return item.GetRandomItem()

}

// Testar items med samma seed om det ger samma item
func TestGetRandomItem(t *testing.T) {
	g := gomega.NewWithT(t)
	for i := 1; i < 20; i++ {
		g.Expect(createItemWithSeed(int64(i))).To(gomega.Equal(createItemWithSeed(int64(i))))
	}
}
