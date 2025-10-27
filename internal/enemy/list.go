package enemy

var ENEMY_LIST = []Enemy{
	func() Enemy { return &Skeleton{health: 25, maxHealth: 25, damage: 10, xp: 75} }(),
	func() Enemy { return &Goblin{health: 50, maxHealth: 50, damage: 4, xp: 50} }(),
	func() Enemy { return &JobApplication{health: 500, maxHealth: 500, damage: 25, xp: 1000} }(),
}
