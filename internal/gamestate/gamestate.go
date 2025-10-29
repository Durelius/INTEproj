package gamestate

import (
	"math/rand"

	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/player"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
)

// GameState holds all information about the current state of the game
type GameState struct {
	Player *player.Player // The player character
	Room   *room.Room     // The current room the player is in
	Battle *battle.Battle
}

// New creates a new GameState instance
func New(p *player.Player, room *room.Room) *GameState {
	return &GameState{
		Player: p,
		Room:   room,
		Battle: nil,
	}
}

func (gs *GameState) InitiateBattle(e enemy.Enemy) {
	playerTurn := rand.Intn(2) == 1
	gs.Battle = battle.New(gs.Player, e, playerTurn)
}

// Currently this method serves to bypass the limitations of Golang not allowing optional arguments.
// This helps us get a base Game going, where player is hard coded to Josh.
func NewDefault() *GameState {
	p := player.New("Josh", class.NewRogue())

	return &GameState{
		Player: p,
		Room:   room.NewRandomRoom(room.NewLocation(0, 0), 25, 50),
		Battle: nil,
	}
}

func (gs *GameState) UpdateRoom() {
	gs.Room = room.NewRandomRoom(room.NewLocation(0, 0), 25, 50)
}
