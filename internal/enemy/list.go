package enemy

import (
	"math/rand"
)

var ENEMY_LIST = []func() Enemy{
	func() Enemy { return &Goblin{health: 50, maxHealth: 50, damage: 4, xp: 50} },
	func() Enemy { return &Skeleton{health: 25, maxHealth: 25, damage: 10, xp: 60} },
	func() Enemy{ return &Succubus{health: 60, maxHealth: 60, damage: 14, xp: 85,}},
	func() Enemy { return &Wraith{health: 80, maxHealth: 80, damage: 15, xp:130} },
	func() Enemy { return &Hellhound{health: 120, maxHealth:120, damage: 22, xp:200} },
	func() Enemy { return &JobApplication{health: 500, maxHealth: 500, damage: 25, xp: 1000} },
}


func NewRandomEnemy() Enemy {
	return ENEMY_LIST[rand.Intn(len(ENEMY_LIST))]()
}