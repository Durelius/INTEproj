package cli

import (
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

type mainState struct{}

func (ms *mainState) view(cli *CLI) (out string) {
	out = "Press B to open bag\n"
	out += cli.generateMapView()
	return
}

func (ms *mainState) update(cli *CLI, msg tea.KeyMsg) {
	loc := cli.game.Room.GetPlayerLocation()
	x, y := loc.Get()

	// Update player locatioon x or y based on input
	switch msg.String() {
	case "up":
		if y > 0 {
			y--
		}
	case "down":
		if y < cli.game.Room.GetHeight()-1 {
			y++
		}
	case "left":
		if x > 0 {
			x--
		}
	case "right":
		if x < cli.game.Room.GetWidth()-1 {
			x++
		}
	case "b":
		cli.view = &inventoryState{}

		return
	}

	// Check if there's a point of interest at the new location
	poi := cli.game.Room.UsePOI(x, y)
	cli.msg = ""

	if poi != nil { // New location has a point of interest
		cli.currentPOI = poi
		switch poi.GetType() {
		case "ENEMY":
			// Spawn random enemy
			// enemy := enemy.ENEMY_LIST[rand.Intn(len(enemy.ENEMY_LIST)) ]
			cli.view = &enemyState{stage: encounter, enemy: enemy.NewSkeleton()}
			return
		case "LOOT":
			cli.msg = "Press E to open the chest, or S to skip!"
			cli.view = &lootState{stage: chest}
			return
		case "EXIT":
			exit := poi.(*room.Exit)
			if exit.IsLocked(cli.game.Room) {
				cli.msg = "The door is locked until all enemies are killed"
			} else {
				cli.game.UpdateRoom(room.NewRandomRoom("Starting area", room.NewLocation(0, 0), 25, 50, nil, nil))
				cli.view = &mainState{}
			}

		default:

		}
	}
	cli.game.Room.SetPlayerLocation(x, y) // Update player location in the room
}
