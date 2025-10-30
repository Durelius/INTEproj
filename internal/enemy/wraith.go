package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Wraith struct {
	maxHealth int
	health    int
	damage    int
	xp        int
}

func NewWraith() *Wraith {
	return &Wraith{
		maxHealth: 150, 
		health:    150,
		damage:    25,
		xp:        150,
	}
}

// This is needed to make the enemy interface implement POI
func (w *Wraith) GetType() string {
	return "ENEMY"
}

func (w *Wraith) GetEnemyType() string {
	return "Wraith"
}

func (w *Wraith) GetCurrentHealth() int {
	return w.health
}

func (w *Wraith) GetMaxHealth() int {
	return w.maxHealth
}

func (w *Wraith) GetDamage() int {
	return w.damage
}

func (w *Wraith) GetXPDrop() int {
	return w.xp
}

func (w *Wraith) IsDead() bool {
	return w.health <= 0
}

func (w *Wraith) TakeDamage(damage int) {
	w.health -= damage
	if w.health <= 0 {
		w.health = 0
	}
}

func (w *Wraith) DropLoot() item.Item {
	return item.GetRandomItem()
}