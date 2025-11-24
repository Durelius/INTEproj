package enemy_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/enemy"
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
		g.Expect(e.GetType()).NotTo(BeNil())
		g.Expect(e.GetEnemyType()).NotTo(BeNil())
		g.Expect(e.GetCurrentHealth()).NotTo(BeNil())
		g.Expect(e.GetMaxHealth()).NotTo(BeNil())
		g.Expect(e.GetDamage()).NotTo(BeNil())
		g.Expect(e.GetXPDrop()).NotTo(BeNil())
		g.Expect(e.DropLoot()).NotTo(BeNil())

		e.IsDead()
		g.Expect(e.IsDead()).To(BeFalse())

		e.TakeDamage(10)
		g.Expect(e.GetCurrentHealth()).To(Equal(e.GetMaxHealth() - 10))

		e.TakeDamage(e.GetCurrentHealth() + 5)
		g.Expect(e.GetCurrentHealth()).To(Equal(0))
		g.Expect(e.IsDead()).To(BeTrue())
	}
}
