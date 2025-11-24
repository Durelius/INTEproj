package room_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/room"
	. "github.com/onsi/gomega"
)

// TEST: Tests for room_list
func TestNewRoomList(t *testing.T) {
	g := NewWithT(t)

	room1 := room.NewRandomRoom()
	rl := room.NewRoomList(room1)

	g.Expect(rl.GetLevelCounter()).To(Equal(1))
	g.Expect(room1.GetLevel()).To(Equal(1))
	g.Expect(rl.GetHead()).To(Equal(room1))
	g.Expect(rl.GetTail()).To(Equal(room1))

	room2 := room.NewRandomRoom()
	rl.Add(room2)

	g.Expect(rl.GetLevelCounter()).To(Equal(2))
	g.Expect(room2.GetLevel()).To(Equal(2))
	g.Expect(rl.GetHead()).To(Equal(room1))
	g.Expect(rl.GetTail()).To(Equal(room2))

	room3 := room.NewRandomRoom()
	rl.Add(room3)

	g.Expect(rl.GetLevelCounter()).To(Equal(3))
	g.Expect(room3.GetLevel()).To(Equal(3))
	g.Expect(rl.GetHead()).To(Equal(room1))
	g.Expect(rl.GetTail()).To(Equal(room3))
}

// TEST: Tests for room
func TestNewRandomRoom(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()

	var exit int
	var loot int
	var enemy int

	for _, poi := range room.GetPOI() {
		switch poi.GetType() {
		case "EXIT":
			exit++
		case "LOOT":
			loot++
		case "ENEMY":
			enemy++
		}
	}

	g.Expect(exit).To(Equal(1))
	g.Expect(enemy).To(SatisfyAll(
		BeNumerically(">", 0),
		BeNumerically("<=", 3),
	))
	g.Expect(loot).To(SatisfyAll(
		BeNumerically(">", 0),
		BeNumerically("<=", 5),
	))
}
func TestUsePOINull(t *testing.T) {
	room := room.NewRandomRoom()
	poi := room.UsePOI(-5, -5)
	if poi != nil {
		t.Errorf("Use POI should return nil if values for X or Y is negative, value: %v", poi)
	}
}
func TestGetLoot(t *testing.T) {
	poi := room.NewLoot()
	loot := poi.(*room.Loot)
	items := loot.GetItems()
	if len(items) != 2 {
		t.Error("Loot should return 2 items")
	}

}
func TestGetLocation(t *testing.T) {
	loc := room.NewLocation(5, 5)
	x, y := loc.Get()
	if x != 5 || y != 5 {
		t.Error("New location doesn't return expected values")
	}
}

// TEST: See that room can handle a lot of POIs
func TestStressRoom(t *testing.T) {
	g := NewWithT(t)
	room := room.NewCustomRoom([]room.PointOfInterest{}, 10, 10, 500, 500)

	var exit int
	var loot int
	var enemy int

	for _, poi := range room.GetPOI() {
		switch poi.GetType() {
		case "EXIT":
			exit++
		case "LOOT":
			loot++
		case "ENEMY":
			enemy++
		}
	}

	g.Expect(exit).To(Equal(1))
	g.Expect(enemy).To(SatisfyAll(
		BeNumerically(">", 0),
		BeNumerically("<=", 500),
	))
	g.Expect(loot).To(SatisfyAll(
		BeNumerically(">", 0),
		BeNumerically("<=", 500),
	))
}

func TestRoomDimensions(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()
	g.Expect(room.GetHeight()).To(Equal(20))
	g.Expect(room.GetWidth()).To(Equal(50))
}

func TestPlayerStartsAtEntry(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()
	g.Expect(room.GetPlayerLocation()).To(Equal(room.GetEntry()))
}

func TestSetPlayerLocation(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()
	room.SetPlayerLocation(10, 15)
	playerLocation := room.GetPlayerLocation()
	g.Expect(playerLocation.GetX()).To(Equal(10))
	g.Expect(playerLocation.GetY()).To(Equal(15))
}

func TestHasEnemiesIsTrue(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()

	g.Expect(room.HasEnemies()).To(BeTrue())
}

func TestHasEnemiesIsFalse(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()

	for location := range room.GetPOI() {
		room.UsePOI(location.GetX(), location.GetY())
	}

	g.Expect(room.HasEnemies()).To(BeFalse())
}

// level is added in roomlist add so if you don't add it to roomlist it will be 0 and return nil
func TestGetNextRoomNil(t *testing.T) {
	room := room.NewRandomRoom()
	next := room.GetNextRoom()
	if next != nil {
		t.Error("next is supposed to be nil here")
	}

}

// level is added in roomlist add so if you don't ddd it to roomlist it will be 0 and return nil
func TestGetPrevRoomNil(t *testing.T) {
	room := room.NewRandomRoom()
	prev := room.GetPrevRoom()
	if prev != nil {
		t.Error("prev is supposed to be nil here")
	}
}

func TestUsePOI(t *testing.T) {
	g := NewWithT(t)
	room := room.NewRandomRoom()

	for location := range room.GetPOI() {
		poi := room.UsePOI(location.GetX(), location.GetY())
		if poi.GetType() != "EXIT" {
			_, exists := room.GetPOI()[location]
			g.Expect(exists).To(BeFalse())
		}
	}
}

func TestExitRemainsAfterUse(t *testing.T) {
	g := NewWithT(t)
	r := room.NewRandomRoom()
	var exitLoc room.Location
	for location, poi := range r.GetPOI() {
		if poi.GetType() == "EXIT" {
			exitLoc = location
			break
		}
	}
	r.UsePOI(exitLoc.GetX(), exitLoc.GetY())
	_, exists := r.GetPOI()[room.NewLocation(
		exitLoc.GetX(),
		exitLoc.GetY(),
	)]
	g.Expect(exists).To(BeTrue())
}

func TestExitLockedWithEnemies(t *testing.T) {
	g := NewWithT(t)
	r := room.NewRandomRoom()
	var exit *room.Exit
	for _, poi := range r.GetPOI() {
		if poi.GetType() == "EXIT" {
			exit = poi.(*room.Exit)
			break
		}
	}
	g.Expect(exit).NotTo(BeNil())
	g.Expect(exit.GetLockedStatus(r)).To(BeTrue())
}

func TestExitUnlocksWithoutEnemies(t *testing.T) {
	g := NewWithT(t)
	r := room.NewRandomRoom()
	var exit *room.Exit
	for location, poi := range r.GetPOI() {
		if poi.GetType() == "EXIT" {
			exit = poi.(*room.Exit)
		} else {
			r.UsePOI(location.GetX(), location.GetY())
		}
	}
	g.Expect(exit).NotTo(BeNil())
	g.Expect(exit.GetLockedStatus(r)).To(BeFalse())
}
