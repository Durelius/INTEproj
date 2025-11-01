package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Succubus struct {
	maxHealth int
	health    int
	damage    int
	xp        int
}

func NewSuccubus() *Succubus {
	return &Succubus{
		maxHealth: 280,  // tougher than Skeleton (25)
		health:    280,
		damage:    15,  // stronger attacks
		xp:        150, // more rewarding to defeat
	}
}

// This is needed to make the enemy interface implement POI
func (s *Succubus) GetType() string {
	return "ENEMY"
}

func (s *Succubus) GetEnemyType() string {
	return "Succubus"
}

func (s *Succubus) GetCurrentHealth() int {
	return s.health
}

func (s *Succubus) GetMaxHealth() int {
	return s.maxHealth
}

func (s *Succubus) GetDamage() int {
	return s.damage
}

func (s *Succubus) GetXPDrop() int {
	return s.xp
}

func (s *Succubus) IsDead() bool {
	return s.health <= 0
}

func (s *Succubus) TakeDamage(damage int) {
	s.health -= damage

	if s.health <= 0 {
		s.health = 0
	}
}

func (s *Succubus) DropLoot() item.Item {
	return item.GetRandomItem()
}