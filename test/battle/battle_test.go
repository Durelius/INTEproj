package battle_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
)


func TestPlayerWin(t *testing.T) {
	p := player.New("TestPlayer", class.NewMage())

	skeleton := enemy.NewSkeleton()
	

	b := battle.New(p, skeleton, true)	// Player starts first

	expectedSkeletonHealth := skeleton.GetCurrentHealth() - p.GetDamage()
 
	b.ProgressFight()

	actualSkeletonHealth := skeleton.GetCurrentHealth()

	if actualSkeletonHealth != expectedSkeletonHealth {
		t.Errorf("Expected enemy health to be %d, got %d", expectedSkeletonHealth, actualSkeletonHealth)
	}

	expectedPlayerHealth := p.GetCurrentHealth() - skeleton.GetDamage()

	b.ProgressFight() // Enemy's turn

	actualPlayerHealth := p.GetCurrentHealth()

	if actualPlayerHealth != expectedPlayerHealth {
		t.Errorf("Expected player health to be %d, got %d", expectedPlayerHealth, actualPlayerHealth)
	}

	for !b.IsOver(){
		b.ProgressFight()
	}

	if !skeleton.IsDead(){
		t.Errorf("Expected enemy to be defeated, but has %d health left", skeleton.GetCurrentHealth())
	}
	if p.IsDead() {
		t.Errorf("Expected player to be alive, but is dead")
	}

	if b.GetStatus() != battle.Victory {
		t.Errorf("Expected status to be victory, got: %d", b.GetStatus())
	}

}

func TestPlayerLoss(t *testing.T) {
	p := player.New("TestPlayer", class.NewMage())

	jobApplication := enemy.NewJobApplication()
	
	b := battle.New(p, jobApplication, false)
 
	for !b.IsOver() {
		b.ProgressFight()
	}


	if jobApplication.IsDead() {
		t.Errorf("Expected job application to win")
	}

	if !p.IsDead() {
		t.Errorf("Expected player to be dead, but has %d health left", p.GetCurrentHealth())
	}
}