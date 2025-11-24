package enemy_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/item"
	. "github.com/onsi/gomega"
)

// TEST: Tests for spawning enemies

func TestSpawningEnemies(t *testing.T) {
	g := NewWithT(t)

	enemies := []enemy.Enemy{
		enemy.NewGoblin(),
		enemy.NewHellhound(),
		enemy.NewHenkeB(),
		enemy.NewJobApplication(),
		enemy.NewSkeleton(),
		enemy.NewSuccubus(),
		enemy.NewWraith(),
	}

	for _, e := range enemies {
		e.GetType()
		e.GetEnemyType()
		e.GetCurrentHealth()
		e.GetMaxHealth()
		e.GetDamage()
		e.GetXPDrop()

		i := e.DropLoot()
		_, ok := i.(item.Item)
		g.Expect(ok).To(BeTrue())

		e.IsDead()
		g.Expect(e.IsDead()).To(BeFalse())

		e.TakeDamage(10)
		g.Expect(e.GetCurrentHealth()).To(Equal(e.GetMaxHealth() - 10))

		e.TakeDamage(e.GetCurrentHealth() + 5)
		g.Expect(e.GetCurrentHealth()).To(Equal(0))
		g.Expect(e.IsDead()).To(BeTrue())
	}
}
