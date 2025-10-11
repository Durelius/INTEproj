package room

import (
	"INTE/projekt/enemy"
	"INTE/projekt/item"
)

func New() *Room {
	return &Room{}
}

type Room struct {
	entry          Location
	height         int
	width          int
	enemies        []enemy.Enemy
	loot           []Loot
	playerLocation Location
}
type Loot struct {
	location Location
	items    []item.Item
}
type Location struct {
	x int
	y int
}

func (l *Location) Get() (x int, y int) {
	return l.x, l.y
}
