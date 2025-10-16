package gamestate

import (
	"github.com/Durelius/INTEproj/internal/player"
	class "github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
)

// GameState holds all information about the current state of the game
type GameState struct {
	Player       *player.Player			// The player character
	Room         *room.Room				// The current room the player is in
}

// New creates a new GameState instance
func New(p *player.Player, room *room.Room) *GameState{
	return &GameState{
		Player: p,
		Room: room,
	}
}

// Currently this method serves to bypass the limitations of Golang not allowing optional arguments.
// This helps us get a base Game going, where player is hard coded to Josh.
func NewDefault () *GameState {
	p := player.New("Josh", &class.Rogue{})

	return &GameState{
		Player: p,
		Room: room.STARTING_AREA,
	}
}










