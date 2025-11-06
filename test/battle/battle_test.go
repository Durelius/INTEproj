package battle_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	. "github.com/onsi/gomega"
)

func TestPlayerWin(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.MAGE_STR)

	skeleton := enemy.NewSkeleton()
	b := battle.New(p, skeleton, true) // Player starts first
	g.Expect(b.PlayerTurn()).To(BeTrue())
	expectedSkeletonHealth := skeleton.GetCurrentHealth() - p.GetDamage()
	b.ProgressFight()
	actualSkeletonHealth := skeleton.GetCurrentHealth()

	g.Expect(actualSkeletonHealth).To(Equal(expectedSkeletonHealth))

	expectedPlayerHealth := p.GetCurrentHealth() - skeleton.GetDamage()
	b.ProgressFight() // Enemy's turn
	actualPlayerHealth := p.GetCurrentHealth()

	g.Expect(actualPlayerHealth).To(Equal(expectedPlayerHealth))

	for !b.IsOver() {
		b.ProgressFight()
	}

	b.ProgressFight()

	g.Expect(skeleton.IsDead()).To(BeTrue())
	g.Expect(p.IsDead()).To(BeFalse())
	g.Expect(b.GetStatus()).To(Equal(battle.Victory))
}

func TestPlayerLoss(t *testing.T) {
	g := NewWithT(t)

	p := player.New("TestPlayer", class.MAGE_STR)
	jobApplication := enemy.NewJobApplication()
	b := battle.New(p, jobApplication, false)

	for !b.IsOver() {
		b.ProgressFight()
	}

	g.Expect(jobApplication.IsDead()).To(BeFalse())
	g.Expect(p.IsDead()).To(BeTrue())
}
