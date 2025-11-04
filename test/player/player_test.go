package playertest

import (
	"math"
	"testing"

	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
)

// Tests that a player levels up correctly and that their damage increases as expected
func TestIncreaseXpAndLevelUp(t *testing.T) {
	p := player.New("TestPlayer", class.MAGE_STR)

	requiredXpToLevel2 := p.CalculateNextLevelExp()

	p.IncreaseExperience(50)
	if p.GetLevel() != 1 {
		t.Errorf("Expected level 1, got %d", p.GetLevel())
	}
	if p.GetExperience() != 50 {
		t.Errorf("Expected 50 xp, got %d", p.GetExperience())
	}
	if p.GetDamage() != 10 {
		t.Errorf("Expected 10 dmg for lvl 1 mage, got %d", p.GetDamage())
	}

	p.IncreaseExperience(100)

	if p.GetLevel() != 2 {
		t.Errorf("Expected level 2, got %d", p.GetLevel())
	}

	expectedExperience := 150 - requiredXpToLevel2
	if p.GetExperience() != expectedExperience {
		t.Errorf("Expected %d xp, got %d", expectedExperience, p.GetExperience())
	}

	if p.GetDamage() != 14 {
		t.Errorf("Expected 14 dmg for lvl 2 mage, got %d", p.GetDamage())
	}

	if p.GetMaxHealth() != 120 {
		t.Errorf("Expected 120 max health for lvl 2 player, got %d", p.GetMaxHealth())
	}
}

func TestLevelUpMultipleTimesOnOneXpDrop(t *testing.T) {
	p := player.New("TestPlayer", class.ROGUE_STR)

	xpToLevel7 := CalculateXpToLevel(7)
	additionalXp := 57

	p.IncreaseExperience(xpToLevel7 + additionalXp)
	if p.GetLevel() != 7 {
		t.Errorf("Expected level 7, got %d", p.GetLevel())
	}
	if p.GetExperience() != additionalXp {
		t.Errorf("Expected %d xp, got %d", additionalXp, p.GetExperience())
	}
	if p.GetDamage() != 61 {
		t.Errorf("Expected 61 dmg for lvl 7 rogue, got %d", p.GetDamage())
	}
}

// func TestEquipItems(t *testing.T) {
// 	item.
// }

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
