package enemy

import (
	"math/rand"
)

var ENEMY_LIST = []func() Enemy{
	func() Enemy { return &Goblin{health: 100, maxHealth: 100, damage: 7, xp: 50} },
	func() Enemy { return &Skeleton{health: 50, maxHealth: 50, damage: 15, xp: 70} },
	func() Enemy{ return &Succubus{health: 280, maxHealth: 280, damage: 15, xp: 150,}},
	func() Enemy { return &Wraith{health: 150, maxHealth: 150, damage: 25, xp:150} },
	func() Enemy { return &Hellhound{health: 300, maxHealth:300, damage: 15, xp:250} },
	func() Enemy { return &JobApplication{health: 750, maxHealth: 750, damage: 30, xp: 1000} },
	func() Enemy {return &HenkeB{health: 2000, maxHealth: 2000, damage:30, xp:3000}},
}


func NewRandomEnemy() Enemy {
	return ENEMY_LIST[rand.Intn(len(ENEMY_LIST))]()
}