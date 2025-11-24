package gamestate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
	"github.com/joho/godotenv"
)

func TestConvertGameStateToSave(t *testing.T) {
	gs := gamestate.New(player.New("testPlayer", class.MAGE_STR), room.NewRandomRoom())
	save := gs.ConvertToSaveType()
	if save == nil {
		t.Error("Save struct is nil")
		return
	}
	//player tests
	if save.Player.Class == nil {
		t.Error("Class is nil")
	}
	if save.Player.Name == "" {
		t.Error("Name is empty in save struct")
	}
	if save.Player.ID == "" {
		t.Error("ID is empty in save struct")
	}
	if save.Player.InventoryItemNames == nil {
		t.Error("item names is not supposed to be nil, ever")
	}

	if save.Player.Level == 0 {
		t.Error("Level can never be 0")
	}

	if save.Player.GearNames == nil {
		t.Error("Gear is not supposed to be nil")
	}

	if save.Player.MaxHealth == 0 {
		t.Error("Max health can never be 0")
	}

	//room tests
	if len(save.Room.PoiList) == 0 {
		t.Error("No pois, there should always be at least 1")
	}

	if save.Room.PlayerLocation == nil {
		t.Error("No player location")
	}

	if save.Room.Entry == nil {
		t.Error("No entry")
	}

	if save.Room.Height == 0 {
		t.Error("No height")
	}

	if save.Room.Width == 0 {
		t.Error("No width")
	}
	if save.Room.PoiList == nil {
		t.Error("No poi list")
	}

}

func TestConvertGameSaveToJSON(t *testing.T) {
	gs := gamestate.New(player.New("testPlayer", class.MAGE_STR), room.NewRandomRoom())
	rl := room.NewRoomList(gs.Room)
	prev := room.NewRandomRoom()
	next := room.NewRandomRoom()
	rl.Add(prev)
	rl.Add(next)
	gs.Room.SetNext(prev)
	gs.Room.SetPrev(prev)
	save := gs.ConvertToSaveType()
	data, err := save.ConvertToBytes()
	if err != nil {
		t.Errorf("Error converting save to JSON: %v", err)
	}
	if len(data) == 0 || data == nil {
		t.Error("JSON is empty")
	}

}

func TestCreateLoadSaveFile(t *testing.T) {
	//load .env file from main dir, as .env is in gitignore we ignore errors
	godotenv.Load(filepath.Join("..", "..", ".env"))
	//only run create file test in local environment when we want, not in build server
	if os.Getenv("CREATE_JSON_TEST") != "true" {
		t.Skip()
	}
	gs := gamestate.New(player.New("testPlayer", class.MAGE_STR), room.NewRandomRoom())
	if err := gs.SaveToFile(); err != nil {
		t.Errorf("Error saving file: %v", err)
		return
	}
	gs2, err := gamestate.LoadSaveFile(gs.GetFileName())
	if err != nil {
		t.Errorf("Error loading save  file: %v", err)
	}
	if gs.Player.GetID() != gs2.Player.GetID() {
		t.Error("Player IDs doesn't match after save and load")
	}

	if err := os.Remove(filepath.Join("savefiles", gs.GetFileName())); err != nil {
		t.Errorf("Save file couldn't be deleted: %v", err)
	}

}

func TestRoundTripGameStateSave(t *testing.T) {
	gs := gamestate.New(player.New("testPlayer", class.MAGE_STR), room.NewRandomRoom())
	rl := room.NewRoomList(gs.Room)
	prev := room.NewRandomRoom()
	next := room.NewRandomRoom()
	rl.Add(prev)
	rl.Add(next)
	gs.Room.SetNext(prev)
	gs.Room.SetPrev(prev)
	p1 := gs.Player
	itemForInv := item.GetRandomItem()
	p1.PickupItem(itemForInv)

	head := item.GetItemByName("Leather Cap")
	p1.PickupItem(head)
	p1.EquipItem(head)

	weapon := item.GetItemByName("The Worldbreaker")
	p1.PickupItem(weapon)
	p1.EquipItem(weapon)

	upperBody := item.GetItemByName("Padded Vest")
	p1.PickupItem(upperBody)
	p1.EquipItem(upperBody)

	lowerBody := item.GetItemByName("Leather Breeches")
	p1.PickupItem(lowerBody)
	p1.EquipItem(lowerBody)

	foot := item.GetItemByName("Ironstride Boots")
	p1.PickupItem(foot)
	p1.EquipItem(foot)

	save := gs.ConvertToSaveType()
	gs2 := save.ConvertSaveToGameState()

	p2 := gs2.Player
	// player tests
	if gs2.RoomList.GetLevelCounter() == 0 {
		t.Errorf("Level counter is always supposed to be initialized at 1, got %d", gs2.RoomList.GetLevelCounter())
	}
	if p1.GetDamage() != p2.GetDamage() {
		t.Errorf("Player damage mismatch: expected %v, got %v", p1.GetDamage(), p2.GetDamage())
	}
	if p1.GetCurrentHealth() != p2.GetCurrentHealth() {
		t.Errorf("Player current health mismatch: expected %v, got %v", p1.GetCurrentHealth(), p2.GetCurrentHealth())
	}
	if p1.GetClassName() != p2.GetClassName() {
		t.Errorf("Player class name mismatch: expected %v, got %v", p1.GetClassName(), p2.GetClassName())
	}
	if p1.GetDamageReduction() != p2.GetDamageReduction() {
		t.Errorf("Player damage reduction mismatch: expected %v, got %v", p1.GetDamageReduction(), p2.GetDamageReduction())
	}
	if p1.GetEquippedWeight() != p2.GetEquippedWeight() {
		t.Errorf("Player equipped weight mismatch: expected %v, got %v", p1.GetEquippedWeight(), p2.GetEquippedWeight())
	}
	if p1.GetExperience() != p2.GetExperience() {
		t.Errorf("Player experience mismatch: expected %v, got %v", p1.GetExperience(), p2.GetExperience())
	}
	if (p1.GetGear().Head != nil) != (p2.GetGear().Head != nil) {
		t.Errorf("Player gear head nil mismatch: expected %v, got %v", p1.GetGear().Head != nil, p2.GetGear().Head != nil)
	}
	if (p1.GetGear().Weapon != nil) != (p2.GetGear().Weapon != nil) {
		t.Errorf("Player gear weapon nil mismatch: expected %v, got %v", p1.GetGear().Weapon != nil, p2.GetGear().Weapon != nil)
	}
	if p1.GetGear().Head != nil && p2.GetGear().Head != nil && p1.GetGear().Head.GetName() != p2.GetGear().Head.GetName() {
		t.Errorf("Player gear head mismatch: expected %v, got %v", p1.GetGear().Head.GetName(), p2.GetGear().Head.GetName())
	}
	if p1.GetGear().Weapon != nil && p2.GetGear().Weapon != nil && p1.GetGear().Weapon.GetName() != p2.GetGear().Weapon.GetName() {
		t.Errorf("Player gear weapon mismatch: expected %v, got %v", p1.GetGear().Weapon.GetName(), p2.GetGear().Weapon.GetName())
	}
	if p1.GetID() != p2.GetID() {
		t.Errorf("Player ID mismatch: expected %v, got %v", p1.GetID(), p2.GetID())
	}
	if p1.GetInventoryWeight() != p2.GetInventoryWeight() {
		t.Errorf("Player inventory weight mismatch: expected %v, got %v", p1.GetInventoryWeight(), p2.GetInventoryWeight())
	}
	if len(p1.GetItems()) != len(p2.GetItems()) {
		t.Errorf("Player inventory length mismatch: expected %v, got %v", len(p1.GetItems()), len(p2.GetItems()))
	}
	if p1.GetLevel() != p2.GetLevel() {
		t.Errorf("Player level mismatch: expected %v, got %v", p1.GetLevel(), p2.GetLevel())
	}
	if p1.GetMaxHealth() != p2.GetMaxHealth() {
		t.Errorf("Player max health mismatch: expected %v, got %v", p1.GetMaxHealth(), p2.GetMaxHealth())
	}
	if p1.GetMaxWeight() != p2.GetMaxWeight() {
		t.Errorf("Player max weight mismatch: expected %v, got %v", p1.GetMaxWeight(), p2.GetMaxWeight())
	}
	if p1.GetName() != p2.GetName() {
		t.Errorf("Player name mismatch: expected %v, got %v", p1.GetName(), p2.GetName())
	}
	if p1.GetTotalDefense() != p2.GetTotalDefense() {
		t.Errorf("Player total defense mismatch: expected %v, got %v", p1.GetTotalDefense(), p2.GetTotalDefense())
	}
	if p1.GetTotalWeight() != p2.GetTotalWeight() {
		t.Errorf("Player total weight mismatch: expected %v, got %v", p1.GetTotalWeight(), p2.GetTotalWeight())
	}
	if p1.IsDead() != p2.IsDead() {
		t.Errorf("Player death state mismatch: expected %v, got %v", p1.IsDead(), p2.IsDead())
	}

	//room tests
	r1 := gs.Room
	r2 := gs2.Room

	if r1.GetEntry() != r2.GetEntry() {
		t.Errorf("Room entry  mismatch: expected %v, got %v", r1.GetEntry(), r2.GetEntry())
	}
	if r1.GetHeight() != r2.GetHeight() {
		t.Errorf("Room height mismatch: expected %v, got %v", r1.GetHeight(), r2.GetHeight())
	}
	if r1.GetLevel() != r2.GetLevel() {
		t.Errorf("Room level mismatch: expected %v, got %v", r1.GetLevel(), r2.GetLevel())
	}
	if r1.GetWidth() != r2.GetWidth() {
		t.Errorf("Room width mismatch: expected %v, got %v", r1.GetWidth(), r2.GetWidth())
	}
	if (r1.GetNextRoom() != nil) != (r2.GetNextRoom() != nil) {
		t.Errorf("Room next nil mismatch: expected %v, got %v", r1.GetNextRoom() != nil, r2.GetNextRoom() != nil)
	}
	if (r1.GetPrevRoom() != nil) != (r2.GetPrevRoom() != nil) {
		t.Errorf("Room prev nil mismatch: expected %v, got %v", r1.GetPrevRoom() != nil, r2.GetPrevRoom() != nil)
	}
	if (r1.GetPrevRoom() != nil) && (r2.GetPrevRoom() != nil) && r1.GetPrevRoom().GetLevel() != r2.GetPrevRoom().GetLevel() {
		t.Errorf("Room prev level mismatch: expected %v, got %v", r1.GetPrevRoom().GetLevel(), r2.GetPrevRoom().GetLevel())
	}
	if (r1.GetNextRoom() != nil) && (r2.GetNextRoom() != nil) && r1.GetNextRoom().GetLevel() != r2.GetNextRoom().GetLevel() {
		t.Errorf("Room next level mismatch: expected %v, got %v", r1.GetNextRoom().GetLevel(), r2.GetNextRoom().GetLevel())
	}
	if len(r1.GetPOI()) != len(r2.GetPOI()) {
		t.Errorf("Room POI length mismatch: expected %v, got %v", len(r1.GetPOI()), len(r2.GetPOI()))
	}
	if r1.GetPlayerLocation() != r2.GetPlayerLocation() {
		t.Errorf("Room player location mismatch: expected %v, got %v", r1.GetPlayerLocation(), r2.GetPlayerLocation())
	}
	if r1.HasEnemies() != r2.HasEnemies() {
		t.Errorf("Room has enemies mismatch: expected %v, got %v", r1.HasEnemies(), r2.HasEnemies())
	}

	//roomlist
	if gs.RoomList.GetLevelCounter() != gs2.RoomList.GetLevelCounter() {
		t.Errorf("Roomlist has level counter mismatch: expected %v, got %v", gs.RoomList.GetLevelCounter(), gs2.RoomList.GetLevelCounter())
	}

}
