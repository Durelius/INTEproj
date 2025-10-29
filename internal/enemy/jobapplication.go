package enemy

import (
	"github.com/Durelius/INTEproj/internal/item"
)

type JobApplication struct {
	health    int
	maxHealth int
	damage    int
	xp        int
}

func NewJobApplication() *JobApplication {
	return &JobApplication{
		health:    500,
		maxHealth: 500,
		damage:    25,
		xp:        1000,
	}
}

// This is needed to make the enemy interface implement POI
func (j *JobApplication) GetType() string {
	return "ENEMY"
}

func (j *JobApplication) GetEnemyType() string {
	return "Job Application"
}

func (j *JobApplication) GetCurrentHealth() int {
	return j.health
}

func (j *JobApplication) GetMaxHealth() int {
	return j.maxHealth
}

func (j *JobApplication) GetDamage() int {
	return j.damage
}

func (j *JobApplication) GetXPDrop() int {
	return j.xp
}

func (j *JobApplication) IsDead() bool {
	return j.health <= 0
}

func (j *JobApplication) TakeDamage(damage int) {
	j.health -= damage
	if j.health <= 0 {
		j.health = 0
	}
}

func (j *JobApplication) DropLoot() item.Item {
	return item.GetRandomItem()
}
