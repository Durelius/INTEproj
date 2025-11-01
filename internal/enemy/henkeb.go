package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type HenkeB struct {
	health    int
	maxHealth int
	damage    int
	xp        int
}

func NewHenkeB() *HenkeB {
	return &HenkeB{
		health:    2000,
		maxHealth: 2000,
		damage:    30,
		xp:        3000,
	}
}

// This is needed to make the enemy interface implement POI
func (h *HenkeB) GetType() string {
	return "ENEMY"
}

func (h *HenkeB) GetEnemyType() string {
	return "HENKE B."
}

func (h *HenkeB) GetCurrentHealth() int {
	return h.health
}

func (h *HenkeB) GetMaxHealth() int {
	return h.maxHealth
}

func (h *HenkeB) GetDamage() int {
	return h.damage
}

func (h *HenkeB) GetXPDrop() int {
	return h.xp
}

func (h *HenkeB) IsDead() bool {
	return h.health <= 0
}

func (h *HenkeB) TakeDamage(damage int) {
	h.health -= damage
	if h.health <= 0 {
		h.health = 0
	}
}

func (h *HenkeB) DropLoot() item.Item {
	return item.GetRandomItem()
}
