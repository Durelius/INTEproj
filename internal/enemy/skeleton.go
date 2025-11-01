package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Skeleton struct {
	maxHealth int
	health    int
	damage    int
	xp        int
}

func NewSkeleton() *Skeleton {
	return &Skeleton{
		maxHealth: 50,
		health:    50,
		damage:    15,
		xp:        60,
	}
}

// This is needed to make the enemy interface implement POI
func (s *Skeleton) GetType() string {
	return "ENEMY"
}

func (s *Skeleton) GetEnemyType() string {
	return "Skeleton"
}

func (s *Skeleton) GetCurrentHealth() int {
	return s.health
}

func (s *Skeleton) GetMaxHealth() int {
	return s.maxHealth
}

func (s *Skeleton) GetDamage() int {
	return s.damage
}

func (s *Skeleton) GetXPDrop() int {
	return s.xp
}

func (s *Skeleton) IsDead() bool {
	return s.health <= 0
}

func (s *Skeleton) TakeDamage(damage int) {
	s.health -= damage

	if s.health <= 0 {
		s.health = 0
	}
}

func (s *Skeleton) DropLoot() item.Item {
	return item.GetRandomItem()
}
