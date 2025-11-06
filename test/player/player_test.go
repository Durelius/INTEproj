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

func TestEquipItems(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.PALADIN_STR)

	baseDmg := p.GetDamage()

	g.Expect(baseDmg).To(Equal(5))
	g.Expect(p.GetTotalDefense()).To(Equal(0))

	helm := item.GetItemByName("Crown of Eternity")
	p.EquipItem(helm)

	g.Expect(p.GetTotalDefense()).To(Equal(65))

	helm2 := item.GetItemByName("Helm of Embersteel")
	p.EquipItem(helm2)

	g.Expect(p.GetTotalDefense()).To(Equal(42))

	weapon := item.GetItemByName("Bloodforged Sword")
	p.EquipItem(weapon)

	g.Expect(p.GetDamage()).To(Equal(baseDmg + 43))

	// Damage reduction is calculated as defence / defence + 100
	def := float32(p.GetTotalDefense())
	expectedDamageReduction := def / (def + 100)

	g.Expect(p.GetDamageReduction()).To(Equal(expectedDamageReduction))

	p.ReceiveDamage(50)
	expectedDamageTaken := int((float32(50) * (1 - p.GetDamageReduction())))

	g.Expect(p.GetCurrentHealth()).To(Equal(p.GetMaxHealth() - expectedDamageTaken))
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
