package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Hellhound struct {
	maxHealth int
	health    int
	damage    int
	xp        int
}

func NewHellhound() *Hellhound {
	return &Hellhound{
		maxHealth: 120, 
		health:    120,
		damage:    22,
		xp:        200, 
	}
}

// This is needed to make the enemy interface implement POI
func (h *Hellhound) GetType() string {
	return "ENEMY"
}

func (h *Hellhound) GetEnemyType() string {
	return "Hellhound"
}

func (h *Hellhound) GetCurrentHealth() int {
	return h.health
}

func (h *Hellhound) GetMaxHealth() int {
	return h.maxHealth
}

func (h *Hellhound) GetDamage() int {
	return h.damage
}

func (h *Hellhound) GetXPDrop() int {
	return h.xp
}

func (h *Hellhound) IsDead() bool {
	return h.health <= 0
}

func (h *Hellhound) TakeDamage(damage int) {
	h.health -= damage
	if h.health <= 0 {
		h.health = 0
	}
}

func (h *Hellhound) DropLoot() item.Item {
	return item.GetRandomItem()
}