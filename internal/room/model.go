package room

import (
	"fmt"
	"math/rand"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
)

type Room struct {
	// name           string
	entry          Location
	height         int
	width          int
	playerLocation Location
	poi            map[Location]PointOfInterest
	prev           *Room
	next           *Room
}

func NewRandomRoom(entry Location, height, width int) *Room {
	room := &Room{entry: entry, height: height, width: width, playerLocation: entry, poi: make(map[Location]PointOfInterest)}
	itemAmount := 5
	enemyAmount := 1
	pois := []PointOfInterest{}
	for i := 0; i < itemAmount; i++ {
		pois = append(pois, &Loot{items: []item.Item{item.GetRandomItem(), item.GetRandomItem()}})
	}
	for i := 0; i < enemyAmount; i++ {
		// This actually has no effect, the CLI will spawn a random enemy when player steps on a poi of type enemy.
		// Since POI is an interface and not a map storing locations to items / monsters,
		// CLI can not access the type of the monster through POI.
		pois = append(pois, enemy.NewRandomEnemy())
	}

	pois = append(pois, &Exit{isLocked: true})
	room.createRandomLocations(pois)
	Rooms.Add(room)

	return room
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

type Exit struct {
	isLocked bool
}

func (*Exit) GetType() string {
	return "EXIT"
}

func (e *Exit) IsLocked(room *Room) bool {
	if room.HasEnemies() {
		return e.isLocked
	}
	e.SetIsLocked()
	return e.isLocked
}

func (e *Exit) SetIsLocked() {
	e.isLocked = false
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

// func (r *Room) GetName() string {
// 	return r.name
// }

func (r *Room) GetPlayerLocation() Location {
	return r.playerLocation
}

func (r *Room) GetPOI() map[Location]PointOfInterest {
	return r.poi
}

func (r *Room) HasEnemies() bool {
	for _, poi := range r.poi {
		if poi.GetType() == "ENEMY" {
			return true
		}
	}
	return false
}

// fetches a point of interest and removes it from the room map
func (r *Room) UsePOI(x, y int) PointOfInterest {
	loc := NewLocation(x, y)
	poi := r.poi[loc]
	if poi != nil && poi.GetType() == "EXIT" {
		return poi
	}
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
