package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type Enemy interface {
	GetType() string	 // This is needed to make the enemy interface implement POI 
	GetEnemyType() string
	GetCurrentHealth() int
	GetMaxHealth() int
	GetDamage() int
	GetXPDrop() int
	IsDead() bool
	TakeDamage(damage int)
	DropLoot() *item.Weapon	// This should be item.Item, but weapons do not implement item.Item interface
}

