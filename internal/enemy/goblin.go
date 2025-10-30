package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Goblin struct {
	health    int
	maxHealth int
	damage    int
	xp        int
}

func NewGoblin() *Goblin {
	return &Goblin{
		health:    100,
		maxHealth: 100,
		damage:    7,
		xp:        50,
	}
}

// This is needed to make the enemy interface implement POI
func (g *Goblin) GetType() string {
	return "ENEMY"
}

func (g *Goblin) GetEnemyType() string {
	return "Goblin"
}

func (g *Goblin) GetCurrentHealth() int {
	return g.health
}

func (g *Goblin) GetMaxHealth() int {
	return g.maxHealth
}

func (g *Goblin) GetDamage() int {
	return g.damage
}

func (g *Goblin) GetXPDrop() int {
	return g.xp
}

func (g *Goblin) IsDead() bool {
	return g.health <= 0
}

func (g *Goblin) TakeDamage(damage int) {
	g.health -= damage

	if g.health <= 0 {
		g.health = 0
	}
}

func (s *Goblin) DropLoot() item.Item {
	return item.GetRandomItem()
}
