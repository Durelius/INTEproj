package room

import (
	"math/rand"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/random"
)

const (
	PoiTypeEnemy = "ENEMY"
	PoiTypeLoot  = "LOOT"
	PoiTypeExit  = "EXIT"

	roomHeight          = 20
	roomWidth           = 50
	standardItemAmount  = 5
	standardEnemyAmount = 3
)

var roomEntryLocation = NewLocation(0, 0)

type Room struct {
	id             string
	level          int
	entry          Location
	height         int
	width          int
	playerLocation Location
	poi            map[Location]PointOfInterest
	prev           *Room
	next           *Room
}

func NewRandomRoom() *Room {
	room := &Room{
		id:             random.String(),
		entry:          roomEntryLocation,
		height:         roomHeight,
		width:          roomWidth,
		playerLocation: roomEntryLocation,
		poi:            make(map[Location]PointOfInterest),
		next:           nil,
		prev:           nil,
	}

	randomizedItemAmount := rand.Intn(standardItemAmount) + 1
	randomizedEnemyMAount := rand.Intn(standardEnemyAmount) + 1

	pointsOfInterest := []PointOfInterest{}

	for range randomizedItemAmount {
		pointsOfInterest = append(pointsOfInterest, NewLoot())
	}

	for range randomizedEnemyMAount {
		pointsOfInterest = append(pointsOfInterest, enemy.NewRandomEnemy())
	}

	pointsOfInterest = append(pointsOfInterest, &Exit{isLocked: true})

	room.assignLocationsToRoom(
		pointsOfInterest,
		roomHeight,
		roomWidth,
	)

	return room
}
func NewCustomRoom(pois map[Location]PointOfInterest, amount ...int) *Room {
	itemAmount := standardItemAmount
	enemyAmount := standardEnemyAmount
	thisRoomHeight := roomHeight
	thisRoomWidth := roomWidth
	if len(amount) == 4 {
		thisRoomHeight = amount[0]
		thisRoomWidth = amount[1]
		itemAmount = amount[2]
		enemyAmount = amount[3]
	}
	room := &Room{
		id:             random.String(),
		entry:          roomEntryLocation,
		height:         thisRoomHeight,
		width:          thisRoomWidth,
		playerLocation: roomEntryLocation,
		poi:            make(map[Location]PointOfInterest),
		next:           nil,
		prev:           nil,
	}

	pointsOfInterest := []PointOfInterest{}
	if len(pois) == 0 {
		pointsOfInterest = append(pointsOfInterest, &Exit{isLocked: true})
		for range itemAmount {
			pointsOfInterest = append(pointsOfInterest, NewLoot())
		}

		for range enemyAmount {
			pointsOfInterest = append(pointsOfInterest, enemy.NewRandomEnemy())
		}

		room.assignLocationsToRoom(
			pointsOfInterest,
			roomHeight,
			roomWidth,
		)
		return room
	}
	room.poi = pois

	return room
}

func (r *Room) assignLocationsToRoom(
	pointsOfInterest []PointOfInterest,
	height,
	width int,
) {
	for _, point := range pointsOfInterest {
		counter := 0
		for {
			location := NewLocation(
				rand.Intn(width-2)+1,
				rand.Intn(height-2)+1,
			)
			counter++
			if _, exist := r.poi[location]; !exist {
				r.poi[location] = point
				break
			} else if counter > 10 {
				break
			}
		}
	}
}

func Load(
	level,
	height,
	width int,
	entry,
	playerLocation Location,
	poi map[Location]PointOfInterest,
) *Room {
	return &Room{
		level:          level,
		height:         height,
		width:          width,
		entry:          entry,
		playerLocation: playerLocation,
		poi:            poi,
	}
}

func (r *Room) SetNext(next *Room) {
	r.next = next
}

func (r *Room) SetPrev(prev *Room) {
	r.prev = prev
}

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
func (r *Room) GetID() string {
	return r.id
}

func (r *Room) HasEnemies() bool {
	for _, p := range r.poi {
		if p.GetType() == PoiTypeEnemy {
			return true
		}
	}
	return false
}

// UsePOI fetches a point of interest and removes it from the room map
func (r *Room) UsePOI(x, y int) PointOfInterest {
	if x < 0 || x >= r.width || y < 0 || y >= r.height {
		return nil
	}

	loc := NewLocation(x, y)
	poi := r.poi[loc]

	if poi != nil && poi.GetType() == PoiTypeExit {
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
	return &Loot{
		items: []item.Item{
			item.GetRandomItem(),
			item.GetRandomItem(),
		},
	}
}

func (l *Loot) GetType() string {
	return PoiTypeLoot
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
	return PoiTypeExit
}

func (e *Exit) GetLockedStatus(room *Room) bool {
	if room.HasEnemies() {
		return e.isLocked
	}
	e.unlockDoor()
	return e.isLocked
}

func (e *Exit) unlockDoor() {
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
