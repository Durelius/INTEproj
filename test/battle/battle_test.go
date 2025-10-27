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

	for !b.IsOver() {
		b.ProgressFight()
	}

	if !skeleton.IsDead(){
		t.Errorf("Expected enemy to be defeated, but has %d health left", skeleton.GetCurrentHealth())
	}
	if p.IsDead() {
		t.Errorf("Expected player to be alive, but is dead")
	}
}

func TestPlayerMultipleCombatPlayerLoss(t *testing.T) {
	p := player.New("TestPlayer", class.NewMage())

	goblin := enemy.NewGoblin()
	skeleton := enemy.NewSkeleton()
	skeleton2 := enemy.NewSkeleton()
	skeleton3 := enemy.NewSkeleton()
	
	b := battle.New(p, goblin, false)
 
	for !b.IsOver() {
		b.ProgressFight()
	}

	b2 := battle.New(p, skeleton, false)	
	for !b2.IsOver() {
		b2.ProgressFight()
	}


	b3 := battle.New(p, skeleton2, false)	
	for !b3.IsOver() {
		b3.ProgressFight()
	}

	b4 := battle.New(p, skeleton3, false)	
	for !b4.IsOver() {
		b4.ProgressFight()
	}

	if !goblin.IsDead() || !skeleton.IsDead() || !skeleton2.IsDead() {
		t.Errorf("Expected first three enemies to be dead")
	}

	if skeleton3.IsDead() {
		t.Errorf("Expected third skeleton to live")
	}

	if !p.IsDead() {
		t.Errorf("Expected player to be dead, but has %d health left", p.GetCurrentHealth())
	}
}