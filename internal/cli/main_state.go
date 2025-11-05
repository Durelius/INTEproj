package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

type mainState struct{}

func (ms *mainState) view(cli *CLI) (out string) {

	out = cli.getHeaderInfo()

	out += "Press B to open bag\n"
	out += cli.generateMapView()

	room := cli.game.Room
	playerLocation := room.GetPlayerLocation()
	x, y := playerLocation.Get()

	out += fmt.Sprintf("Room: Size=(%dx%d)\n", room.GetWidth(), room.GetHeight())
	out += fmt.Sprintf("Location: Level=(%d) Pos=(%d,%d)\n\n", room.GetLevel(), x, y)
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
			enemy := poi.(enemy.Enemy)
			cli.view = &battleState{stage: encounter, enemy: enemy}
			return
		case "LOOT":
			cli.view = &lootState{stage: chest}
			return
		case "EXIT":
			exit := poi.(*room.Exit)
			if exit.IsLocked(cli.game.Room) {
				cli.msg = "The door is locked until all enemies are killed"
			} else {
				cli.game.UpdateRoom()
				cli.view = &mainState{}
				return
			}

		default:

		}
	}
	cli.game.Room.SetPlayerLocation(x, y) // Update player location in the room
}
