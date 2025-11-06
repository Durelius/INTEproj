package playertest

import (
	"math"
	"testing"

	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	. "github.com/onsi/gomega"
)

// Tests that a player levels up correctly and that their damage increases as expected
func TestIncreaseXpAndLevelUp(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.MAGE_STR)
	requiredXpToLevel2 := p.CalculateNextLevelExp()
	p.IncreaseExperience(50)

	g.Expect(p.GetLevel()).To(Equal(1))
	g.Expect(p.GetExperience()).To(Equal(50))
	g.Expect(p.GetDamage()).To(Equal(10))

	p.IncreaseExperience(100)
	expectedExperience := 150 - requiredXpToLevel2

	g.Expect(p.GetLevel()).To(Equal(2))
	g.Expect(p.GetExperience()).To(Equal(expectedExperience))
	g.Expect(p.GetDamage()).To(Equal(14))
	g.Expect(p.GetMaxHealth()).To(Equal(120))
}

func TestLevelUpMultipleTimesOnOneXpDrop(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.ROGUE_STR)
	xpToLevel7 := CalculateXpToLevel(7)
	additionalXp := 57
	p.IncreaseExperience(xpToLevel7 + additionalXp)

	g.Expect(p.GetLevel()).To(Equal(7))
	g.Expect(p.GetExperience()).To(Equal(additionalXp))
	g.Expect(p.GetDamage()).To(Equal(61))
}

func TestPickupAndDropItem(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.ROGUE_STR)
	i := item.GetItemByName("Runeblade")

	p.EquipItem(i)
	g.Expect(p.GetTotalWeight()).To(Equal(i.GetWeight()))

	for p.GetTotalWeight()+i.GetWeight() < p.GetMaxWeight() {
		p.PickupItem(i)
	}

	g.Expect(p.PickupItem(i)).To(MatchError("Overburdened"))

	p.DropItem(i)

	g.Expect(p.PickupItem(i)).To(BeNil())
}

func TestEquipItems(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.PALADIN_STR)

	baseDmg := p.GetDamage()

	g.Expect(baseDmg).To(Equal(5))
	g.Expect(p.GetTotalDefense()).To(Equal(0))

	helm := item.GetItemByName("Crown of Eternity")
	torso := item.GetItemByName("Padded Vest")
	legs := item.GetItemByName("Cloth Trousers")
	boots := item.GetItemByName("Worn Boots")
	weapon := item.GetItemByName("Soulfire Edge")
	itemSet1 := []item.Item{helm, torso, legs, boots, weapon}
	gear := p.GetGear()

	for _, item := range itemSet1 {
		p.EquipItem(item)
	}

	g.Expect(gear.Head).To(Equal(helm))
	g.Expect(gear.Upperbody).To(Equal(torso))
	g.Expect(gear.Legs).To(Equal(legs))
	g.Expect(gear.Feet).To(Equal(boots))
	g.Expect(gear.Weapon).To(Equal(weapon))
	g.Expect(p.GetDamage()).To(Equal(baseDmg + 63))
	g.Expect(p.GetTotalDefense()).To(Equal(65 + 18 + 12 + 8))

	helm2 := item.GetItemByName("Steel Helm")
	torso2 := item.GetItemByName("Reinforced Jerkin")
	legs2 := item.GetItemByName("Chain Leggings")
	boots2 := item.GetItemByName("Ironshod Boots")
	weapon2 := item.GetItemByName("Twilight Katana")
	itemSet2 := []item.Item{helm2, torso2, legs2, boots2, weapon2}

	for _, item := range itemSet2 {
		p.EquipItem(item)
	}

	gear = p.GetGear()
	g.Expect(gear.Head).To(Equal(helm2))
	g.Expect(gear.Upperbody).To(Equal(torso2))
	g.Expect(gear.Legs).To(Equal(legs2))
	g.Expect(gear.Feet).To(Equal(boots2))
	g.Expect(gear.Weapon).To(Equal(weapon2))

	g.Expect(p.UnequipHead()).To(BeTrue())
	g.Expect(p.UnequipUpperBody()).To(BeTrue())
	g.Expect(p.UnequipLowerBody()).To(BeTrue())
	g.Expect(p.UnequipFeet()).To(BeTrue())
	g.Expect(p.UnequipWeapon()).To(BeTrue())

	g.Expect(p.UnequipHead()).To(BeFalse())
	g.Expect(p.UnequipUpperBody()).To(BeFalse())
	g.Expect(p.UnequipLowerBody()).To(BeFalse())
	g.Expect(p.UnequipFeet()).To(BeFalse())
	g.Expect(p.UnequipWeapon()).To(BeFalse())

	inv := p.GetItems()
	g.Expect(inv).To(ContainElements(itemSet1))
	g.Expect(inv).To(ContainElements(itemSet2))
}

func TestClasses(t *testing.T) {
	g := NewWithT(t)

	// Rogue
	r := class.NewRogue()

	g.Expect(r.Name()).To(Equal(class.ROGUE_STR))
	g.Expect(r.GetDescription()).To(Equal("Stealth assassin."))
	g.Expect(r.GetBaseDmg()).To(Equal(7))
	g.Expect(r.GetEnergy()).To(Equal(100))
	r.IncreaseStats(2) // Simulate player getting lvl 2
	g.Expect(r.GetBaseDmg()).To(Equal(11))
	g.Expect(r.GetEnergy()).To(Equal(120))

	// Mage
	m := class.NewMage()

	g.Expect(m.Name()).To(Equal(class.MAGE_STR))
	g.Expect(m.GetDescription()).To(Equal("Magic caster."))
	g.Expect(m.GetBaseDmg()).To(Equal(10))
	g.Expect(m.GetEnergy()).To(Equal(100))
	m.IncreaseStats(2) // Simulate player getting lvl 2
	g.Expect(m.GetBaseDmg()).To(Equal(14))
	g.Expect(m.GetEnergy()).To(Equal(120))

	// Paladin
	p := class.NewPaladin()

	g.Expect(p.Name()).To(Equal(class.PALADIN_STR))
	g.Expect(p.GetDescription()).To(Equal("Magic tank."))
	g.Expect(p.GetBaseDmg()).To(Equal(5))
	g.Expect(p.GetEnergy()).To(Equal(100))
	p.IncreaseStats(2) // Simulate player getting lvl 2
	g.Expect(p.GetBaseDmg()).To(Equal(8))
	g.Expect(p.GetEnergy()).To(Equal(120))

}

// Utility function to aid testing
// Calculates the total XP required to reach a certain level, from level 1
func CalculateXpToLevel(level int) int {
	if level <= 1 {
		return 0
	}

	totalXp := 0
	for i := 1; i < level; i++ {
		totalXp += int(math.Round(50 * math.Pow(1.5, float64(i))))
	}

	return totalXp
}
