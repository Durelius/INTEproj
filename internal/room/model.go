package room

import (
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
)

type Room struct {
	level          int
	entry          Location
	height         int
	width          int
	playerLocation Location
	poi            map[Location]PointOfInterest
	prev           *Room
	next           *Room
}

func NewRandomRoom(entry Location, height, width int, extraPOIs ...int) *Room {
	room := &Room{entry: entry, height: height, width: width, playerLocation: entry, poi: make(map[Location]PointOfInterest), next: nil, prev: nil}
	var extraAmount int
	if len(extraPOIs) > 0 {
		extraAmount = extraPOIs[0]
	}
	itemAmount := random.IntList(5) + 1 + extraAmount
	enemyAmount := random.IntList(3) + 1 + extraAmount
	pois := []PointOfInterest{}

	for range itemAmount {
		pois = append(pois, &Loot{items: []item.Item{item.GetRandomItem(), item.GetRandomItem()}})
	}

	for range enemyAmount {
		pois = append(pois, enemy.NewRandomEnemy())
	}

	pois = append(pois, &Exit{isLocked: true})

	room.createRandomLocations(pois, height, width)

	return room
}

func (r *Room) createRandomLocations(pois []PointOfInterest, height, width int) {
	for _, poi := range pois {
		counter := 0
		for {
			loc := NewLocation(random.IntList((width-2)+1), random.IntList((height-2)+1))
			counter++
			if _, exist := r.poi[loc]; !exist {
				r.poi[loc] = poi
				break
			} else if counter > 100 {
				break
			}
		}
	}
}
func Load(level, height, width int, entry, playerLocation Location, poi map[Location]PointOfInterest) *Room {
	return &Room{level: level, height: height, width: width, entry: entry, playerLocation: playerLocation, poi: poi}
}
func (r *Room) SetNext(next *Room) {
	r.next = next
}
func (r *Room) SetPrev(prev *Room) {
	r.prev = prev
}

// func (room *Room) createRandomLocations(pois []PointOfInterest) error {
// 	for _, poi := range pois {
// 		attempts := 0
// 		var location Location
// 		for {
// 			location = Location{x: random.Int(1, room.width), y: random.Int(1, room.height)}
// 			jump := false
// 			for i := 1; i <= 5; i++ {
// 				tempLocXUpper := location
// 				tempLocXLower := location
// 				tempLocXLower.x = location.x - i
// 				tempLocXUpper.x = location.x + i
// 				if room.poi[tempLocXLower] != nil || room.poi[tempLocXUpper] != nil {
// 					jump = true
// 					break
// 				}
// 				tempLocYUpper := location
// 				tempLocYLower := location
// 				tempLocYLower.y = location.y - i
// 				tempLocYUpper.y = location.y + i
// 				if room.poi[tempLocYLower] != nil || room.poi[tempLocYUpper] != nil {
// 					jump = true
// 					break
// 				}
// 			}
// 			attempts++
// 			if !jump {
// 				room.poi[location] = poi
// 				break
// 			}
// 			if attempts > 100 {
// 				return fmt.Errorf("Too many POIs in a room, they can't fit")
// 			}
// 		}
// 	}
// 	return nil
// }

func (r *Room) GetHeight() int {
	return r.height
}

func (r *Room) GetWidth() int {
	return r.width
}

func (r *Room) GetLevel() int {
	return r.level
}

func (r *Room) GetPlayerLocation() Location {
	return r.playerLocation
}

func (r *Room) GetPOI() map[Location]PointOfInterest {
	return r.poi
}
func (r *Room) GetEntry() Location {
	return r.entry
}
func (r *Room) GetNextRoom() *Room {
	if r.level == 0 {
		return nil
	}
	return r.next
}
func (r *Room) GetPrevRoom() *Room {
	if r.level == 0 {
		return nil
	}
	return r.prev
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

type PointOfInterest interface {
	GetType() string
}

type Loot struct {
	items []item.Item
}

func NewLoot() PointOfInterest {
	return &Loot{items: []item.Item{item.GetRandomItem(), item.GetRandomItem()}}
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

func NewExit() PointOfInterest {
	return &Exit{isLocked: true}
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
func (l *Location) GetX() (x int) {
	return l.x
}
func (l *Location) GetY() (y int) {
	return l.y
}

func NewLocation(x int, y int) Location {
	return Location{x: x, y: y}
}
