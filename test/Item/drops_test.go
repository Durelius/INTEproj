package itemtest

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/onsi/gomega"
)

// Create named items of all rarities. Testcase 3 from tillståndsmaskin
func TestGetItemsByName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	str := []string{"Rusty Sword", "Training Sword", "Soulfire Edge", "Crown of Thorns", "Crown of Eternity"}
	count := len(str)

	for i := 0; i < count; i++ {
		item := item.GetItemByName(str[i])
		g.Expect(str[i]).To(gomega.Equal(item.GetName()))
		g.Expect(item).NotTo(gomega.BeNil())
	}
}

// Create empty item som inte finns
func TestCreateEmptyItem(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	str := []string{"", "NotItemName", "     ", "nil", "Ru57y 5w0RD"}
	count := len(str)

	for i := 0; i < count; i++ {
		item := item.GetItemByName(str[i])
		g.Expect(item).To(gomega.BeNil())
	}
}

// Cannot Create Item From NIL
func TestCreateNilItem(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	var nilString string
	item := item.GetItemByName(nilString)
	g.Expect(item).To(gomega.BeNil())
}
