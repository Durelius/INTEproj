package gamestate

import (
	"encoding/json"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/player/gear"
	"github.com/Durelius/INTEproj/internal/room"
)

type GameSave struct {
	Player       *playerSave `json:"player"`       // The player character
	Room         *roomSave   `json:"room"`         // The current room the player is in
	LevelCounter int         `json:"levelCounter"` // amount of rooms existing
}

func (gs *GameState) ConvertToSaveType() *GameSave {
	var save GameSave
	save.Player = convertPlayerToSave(gs.Player)
	save.Room = convertRoomToSave(gs.Room)
	save.LevelCounter = gs.RoomList.GetLevelCounter()
	return &save
}
func (save *GameSave) ConvertSaveToGameState() *GameState {
	var gs GameState
	gs.Player = convertSaveToPlayer(save.Player)
	gs.Room = convertSaveToRoom(save.Room)
	gs.RoomList = room.LoadRoomList(gs.Room, save.LevelCounter)
	return &gs
}
func (gs *GameSave) ConvertToBytes() ([]byte, error) {
	return json.Marshal(gs)
}

func (gs *GameSave) getFileName() string {
	return gs.Player.Name + "_" + gs.Player.ID + ".json"
}

func convertRoomToSave(room *room.Room) *roomSave {
	var save roomSave
	save.Entry = convertLocationToSave(room.GetEntry())
	save.Height = room.GetHeight()
	save.Width = room.GetWidth()
	save.Level = room.GetLevel()
	nextRoom := room.GetNextRoom()
	if nextRoom != nil && nextRoom.GetLevel() > 0 {
		save.Next = convertRoomToSave(nextRoom)
	}
	prevRoom := room.GetPrevRoom()
	if prevRoom != nil && prevRoom.GetLevel() > 0 {
		save.Prev = convertRoomToSave(prevRoom)
	}
	save.PlayerLocation = convertLocationToSave(room.GetPlayerLocation())
	save.PoiList = []poiSave{}
	for k, v := range room.GetPOI() {
		save.PoiList = append(save.PoiList, poiSave{Location: convertLocationToSave(k), Poi: v.GetType()})
	}

	return &save
}
func convertSaveToRoom(save *roomSave) *room.Room {
	poiMap := make(map[room.Location]room.PointOfInterest)
	for _, v := range save.PoiList {
		var poi room.PointOfInterest
		switch v.Poi {
		case "ENEMY":
			poi = enemy.NewRandomEnemy()
		case "EXIT":
			poi = room.NewExit()
		case "LOOT":
			poi = room.NewLoot()
		default:
			return nil
		}
		poiMap[convertSaveToLocation(*v.Location)] = poi

	}
	room := room.Load(save.Level, save.Height, save.Width, convertSaveToLocation(*save.Entry), convertSaveToLocation(*save.PlayerLocation), poiMap)
	if save.Next != nil {
		room.SetNext(convertSaveToRoom(save.Next))
	}
	if save.Prev != nil {
		room.SetNext(convertSaveToRoom(save.Prev))
	}
	return room
}
func convertLocationToSave(l room.Location) *locationSave {
	var save locationSave
	save.X = l.GetX()
	save.Y = l.GetY()
	return &save

}
func convertSaveToLocation(save locationSave) room.Location {
	return room.NewLocation(save.X, save.Y)

}

func convertPlayerToSave(p *player.Player) *playerSave {
	var save playerSave
	save.CurrentHealth = p.GetCurrentHealth()
	save.Dead = p.IsDead()
	save.Experience = p.GetExperience()
	save.ID = p.GetID()
	save.Level = p.GetLevel()
	save.MaxHealth = p.GetMaxHealth()
	save.MaxWeight = p.GetMaxWeight()
	save.Name = p.GetName()

	class := p.GetClass()
	save.Class = &classSave{BaseDmg: class.GetBaseDmg(), Energy: class.GetEnergy(), ClassName: p.GetClassName()}

	save.InventoryItemNames = []string{}
	for _, item := range p.GetItems() {
		save.InventoryItemNames = append(save.InventoryItemNames, item.GetName())
	}

	gear := p.GetGear()
	save.GearNames = &gearSave{
		HeadItemName: func() string {
			if gear.Head != nil {
				return gear.Head.GetName()
			}
			return ""
		}(),
		UpperbodyItemName: func() string {
			if gear.Upperbody != nil {
				return gear.Upperbody.GetName()
			}
			return ""
		}(),
		LegItemName: func() string {
			if gear.Legs != nil {
				return gear.Legs.GetName()
			}
			return ""
		}(),
		WeaponItemName: func() string {
			if gear.Weapon != nil {
				return gear.Weapon.GetName()
			}
			return ""
		}(),
		FeetItemName: func() string {
			if gear.Feet != nil {
				return gear.Feet.GetName()
			}
			return ""
		}(),
	}

	return &save

}
func convertSaveToPlayer(save *playerSave) *player.Player {
	var itemlist []item.Item
	for _, name := range save.InventoryItemNames {
		itemlist = append(itemlist, item.FindItemByName(name))
	}
	gear := gear.Gear{Head: item.FindItemByName(save.GearNames.HeadItemName),
		Upperbody: item.FindItemByName(save.GearNames.UpperbodyItemName),
		Legs:      item.FindItemByName(save.GearNames.LegItemName),
		Weapon:    item.FindItemByName(save.GearNames.WeaponItemName),
		Feet:      item.FindItemByName(save.GearNames.FeetItemName),
	}
	p := player.Load(save.Name, save.Class.ClassName, itemlist, gear, save.ID, save.MaxWeight, save.MaxHealth, save.CurrentHealth, save.Level, save.Experience, save.Class.BaseDmg, save.Class.Energy)
	return p

}

type roomSave struct {
	Level          int           `json:"level"`
	Entry          *locationSave `json:"entry"`
	Height         int           `json:"height"`
	Width          int           `json:"width"`
	PlayerLocation *locationSave `json:"playerLocation"`
	PoiList        []poiSave     `json:"poiList"`
	Next           *roomSave     `json:"nextRoom"`
	Prev           *roomSave     `json:"prevRoom"`
}
type poiSave struct {
	Location *locationSave `json:"location"`
	Poi      string        `json:"poi"`
}
type locationSave struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type playerSave struct {
	Name               string     `json:"name"`
	Class              *classSave `json:"class"`
	InventoryItemNames []string   `json:"inventory"`
	GearNames          *gearSave  `json:"gear"`
	MaxWeight          int        `json:"maxWeight"`
	MaxHealth          int        `json:"maxHealth"`
	CurrentHealth      int        `json:"currentHealth"`
	Level              int        `json:"level"`
	Experience         int        `json:"experience"`
	Dead               bool       `json:"dead"`
	ID                 string     `json:"id"`
}
type gearSave struct {
	HeadItemName      string `json:"head"`
	UpperbodyItemName string `json:"upperBody"`
	LegItemName       string `json:"legs"`
	WeaponItemName    string `json:"weapon"`
	FeetItemName      string `json:"feet"`
}
type classSave struct {
	BaseDmg   int             `json:"baseDmg"`
	Energy    int             `json:"energy"`
	ClassName class.ClassName `json:"className"`
}
