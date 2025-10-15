package room

import (
	"INTE/projekt/enemy"
	"INTE/projekt/item"
	"INTE/projekt/random"
	"fmt"
	"math/rand"
)

func newRandomRoom(name string, entry Location, height, width int) *Room {
	room := &Room{name: name, entry: entry, height: height, width: width, playerLocation: entry, poi: make(map[Location]PointOfInterest)}
	itemAmount := 50
	enemyAmount := 5
	pois := []PointOfInterest{}
	for i := 0; i < itemAmount; i++ {
		itemsInChest := random.Int(1, 5)
		loot := &Loot{items: []item.Item{}}
		for i := 0; i < itemsInChest; i++ {
			index := random.Int(0, len(item.ITEM_LIST_DROPPABLE)-1)
			loot.items = append(loot.items, item.ITEM_LIST_DROPPABLE[index])
		}
		pois = append(pois, loot)

	}
	for i := 0; i < enemyAmount; i++ {
		index := rand.Intn(len(enemy.ENEMY_LIST))
		pois = append(pois, enemy.ENEMY_LIST[index])
	}
	room.createRandomLocations(pois)

	return room
}

type Room struct {
	name           string
	entry          Location
	height         int
	width          int
	playerLocation Location
	poi            map[Location]PointOfInterest
	seed           int
}
type PointOfInterest interface {
	GetType() string
}
type Loot struct {
	items []item.Item
}

func (l *Loot) GetType() string {
	return "LOOT"
}
func (l *Loot) GetItems() []item.Item {
	return l.items
}

type Location struct {
	x int
	y int
}

func (l *Location) Get() (x int, y int) {
	return l.x, l.y
}
func NewLocation(x int, y int) Location {
	return Location{x: x, y: y}
}
func (r *Room) GetHeight() int {
	return r.height
}
func (r *Room) GetWidth() int {
	return r.width
}
func (r *Room) GetName() string {
	return r.name
}
func (r *Room) GetPlayerLocation() Location {
	return r.playerLocation
}

// fetches a point of interest and removes it from the room map
func (r *Room) UsePOI(x, y int) PointOfInterest {
	loc := NewLocation(x, y)
	poi := r.poi[loc]
	delete(r.poi, loc)
	return poi
}

func (r *Room) SetPlayerLocation(x, y int) {
	r.playerLocation = NewLocation(x, y)
}
func (room *Room) createRandomLocations(pois []PointOfInterest) error {
	for _, poi := range pois {
		attempts := 0
		var location Location
		for {
			location = Location{x: random.Int(1, room.width), y: random.Int(1, room.height)}
			jump := false
			for i := 1; i <= 5; i++ {
				tempLocXUpper := location
				tempLocXLower := location
				tempLocXLower.x = location.x - i
				tempLocXUpper.x = location.x + i
				if room.poi[tempLocXLower] != nil || room.poi[tempLocXUpper] != nil {
					jump = true
					break
				}
				tempLocYUpper := location
				tempLocYLower := location
				tempLocYLower.y = location.y - i
				tempLocYUpper.y = location.y + i
				if room.poi[tempLocYLower] != nil || room.poi[tempLocYUpper] != nil {
					jump = true
					break
				}
			}
			attempts++
			if !jump {
				room.poi[location] = poi
				break
			}
			if attempts > 100 {
				return fmt.Errorf("Too many POIs in a room, they can't fit")
			}
		}
	}
	return nil
}
