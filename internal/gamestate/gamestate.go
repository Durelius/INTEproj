package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/player"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/random"
	"github.com/Durelius/INTEproj/internal/room"
)

// GameState holds all information about the current state of the game
type GameState struct {
	Player   *player.Player // The player character
	Room     *room.Room     // The current room the player is in
	Battle   *battle.Battle
	RoomList *room.RoomList
}

// New creates a new GameState instance
func New(p *player.Player, r *room.Room) *GameState {
	gs := GameState{
		Player:   p,
		Room:     r,
		Battle:   nil,
		RoomList: room.NewRoomList(r),
	}
	TimerSave(&gs)
	return &gs
}

// save every 5 seconds
func TimerSave(gs *GameState) {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if gs == nil || gs.Room == nil {
				continue
			}
			err := gs.SaveToFile()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
}

func (gs *GameState) InitiateBattle(e enemy.Enemy) {
	playerTurn := random.IntList(2) == 1
	gs.Battle = battle.New(gs.Player, e, playerTurn)
}

func (gs *GameState) UpdateRoom() {
	gs.Room = room.NewRandomRoom()
	gs.RoomList.Add(gs.Room)
}

// Currently this method serves to bypass the limitations of Golang not allowing optional arguments.
// This helps us get a base Game going, where player is hard coded to Josh.
func NewDefault() *GameState {
	p := player.New("Josh", class.ROGUE_STR)
	r := room.NewRandomRoom()
	rl := room.NewRoomList(r)
	gs := GameState{
		Player:   p,
		Room:     r,
		Battle:   nil,
		RoomList: rl,
	}
	go TimerSave(&gs)
	return &gs
}

func (gs *GameState) GetFileName() string {
	return gs.Player.GetName() + "_" + gs.Player.GetID() + ".json"
}

func (gs *GameState) SaveToFile() error {
	save := gs.ConvertToSaveType()
	data, err := save.ConvertToBytes()
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join("savefiles", save.getFileName()), data, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func (gs *GameState) GetSaveFiles() ([]string, error) {
	savefileNames := []string{}
	data, err := os.ReadDir(filepath.Join("savefiles"))
	if err != nil {
		return nil, err
	}
	for _, file := range data {
		savefileNames = append(savefileNames, file.Name())
	}
	return savefileNames, nil
}

func LoadSaveFile(filename string) (*GameState, error) {
	data, err := os.ReadFile(filepath.Join("savefiles", filename))
	if err != nil {
		return nil, err
	}
	var save GameSave

	if err := json.Unmarshal(data, &save); err != nil {
		return nil, err
	}
	gs := save.ConvertSaveToGameState()
	return gs, nil
}
