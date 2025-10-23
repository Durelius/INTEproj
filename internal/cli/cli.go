package cli

import (
	"fmt"
	"log"

	"github.com/Durelius/INTEproj/internal/assets/ascii"
	gs "github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

type view int

const (
	MAIN view = iota
	BATTLE
	INVENTORY
	LOOT
	LOOT_LIST
)

// The CLI struct is the main model for the command line interface.
// It listens for user input and updates the gamestate, as well as the current view.
type CLI struct {
	game         *gs.GameState
	view         view
	msg          string
	cursor       int // Cursor position in the current view
	checkedIndex int // Index of the currently checked item in lists
	currentPOI   room.PointOfInterest
}

func New(game *gs.GameState) CLI {
	return CLI{game: game}
}

func (cli CLI) Init() tea.Cmd {
	return nil
}

func (cli *CLI) generateMapView() string {
	out := "\nMap:\n"
	loc := cli.game.Room.GetPlayerLocation()
	for y := 0; y < cli.game.Room.GetHeight(); y++ {
		for x := 0; x < cli.game.Room.GetWidth(); x++ {
			locX, locY := loc.Get()
			if x == locX && y == locY {
				out += "@"
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return out
}

func (cli CLI) View() string {
	out := cli.getHeaderInfo()

	switch cli.view {
	case MAIN:
		out += "Press B to open bag\n"
		out += cli.generateMapView()
	case LOOT:
		out += ascii.Chest()
	case LOOT_LIST:
		loot := cli.currentPOI.(*room.Loot)
		out += "Select item from (space to toggle, enter to confirm):\n\n"
		for i, item := range loot.GetItems() {
			cursor := " " // no cursor
			if i == cli.cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if cli.checkedIndex == i {
				checked = "[x]"
			}

			out += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case INVENTORY:
		out += "Inventory:\n\n"

		for i, item := range cli.game.Player.GetItems() {
			cursor := " " // no cursor
			if i == cli.cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if cli.checkedIndex == i {
				checked = "[x]"
			}

			out += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case BATTLE:
		out += "\n BATTLE TIME:\n"
	}
	return out
}

// Update reads a message (user input) and updates the view accordingly.
func (cli CLI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return cli, tea.Quit
		}

		switch cli.view {
		case INVENTORY:
			cli.msg = "Press space to toggle items, X to drop and ENTER to consume"
			return cli.updateBag(msg)
		case BATTLE:
			return cli.updateBattle(msg)
		case LOOT:
			return cli.updateLoot(msg)
		case LOOT_LIST:
			return cli.updateLootList(msg)
		default:
			return cli.updateMain(msg)
		}

	}
	return cli, nil
}

// Fetches all relevant information which is always displayed at the top of the CLI
func (cli *CLI) getHeaderInfo() string {
	player := cli.game.Player
	room := cli.game.Room
	loc := room.GetPlayerLocation()
	x, y := loc.Get()

	s := fmt.Sprintf("Room: %s (%dx%d)\n", room.GetName(), room.GetWidth(), room.GetHeight())
	s += fmt.Sprintf("Player: %s (%v) HP:%d/%d\n", player.GetName(), player.GetClass(), player.GetCurrentHealth(), player.GetMaxHealth())
	s += fmt.Sprintf("Location: (%d,%d)\n", x, y)
	s += cli.msg + "\n"
	return s
}

// Updates the main view, where the player can move around the room
func (cli *CLI) updateMain(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
		cli.view = INVENTORY
		return cli, nil
	}

	// Check if there's a point of interest at the new location
	poi := cli.game.Room.UsePOI(x, y)
	if poi != nil { // New location has a point of interest
		cli.currentPOI = poi
		switch poi.GetType() {
		case "ENEMY":
			cli.msg = "You encountered an enemy! Prepare to fight!"
			cli.view = BATTLE
		case "LOOT":
			cli.msg = "Press E to open the chest, or S to skip!"
			cli.view = LOOT
		default:
		}
	}
	cli.game.Room.SetPlayerLocation(x, y) // Update player location in the room
	return cli, nil
}

func (cli *CLI) updateBag(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	items := cli.game.Player.GetItems()
	switch msg.String() {
	case "b":
		cli.view = MAIN
		return cli, nil
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(items)-1 {
			cli.cursor++
		}
	case " ":
		if cli.checkedIndex == cli.cursor {
			cli.checkedIndex = 1337
		} else {
			cli.checkedIndex = cli.cursor
		}
	case "enter":
		if cli.checkedIndex == 1337 {
			cli.msg = "You got nothing..."
			cli.view = MAIN
			return cli, nil
		}
		if err := cli.game.Player.PickupItem(items[cli.checkedIndex]); err != nil {
			cli.msg = err.Error()
			return cli, nil
		}
		cli.msg = "You picked up " + items[cli.checkedIndex].GetName()
		cli.view = MAIN
		cli.checkedIndex = 1337
	}

	return cli, nil
}

func (cli *CLI) updateLoot(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "e":
		cli.msg = ""
		cli.view = LOOT_LIST
	case "s":
		cli.view = MAIN
	}

	return cli, nil
}

func (cli *CLI) updateLootList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	loot := cli.currentPOI.(*room.Loot)

	switch msg.String() {
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(loot.GetItems())-1 {
			cli.cursor++
		}
	case " ":
		if cli.checkedIndex == cli.cursor {
			cli.checkedIndex = 1337
		} else {
			cli.checkedIndex = cli.cursor
		}
	case "enter":

		if cli.checkedIndex == 1337 {
			cli.msg = "You got nothing..."
			cli.view = MAIN
			return cli, nil
		}
		if err := cli.game.Player.PickupItem(loot.GetItems()[cli.checkedIndex]); err != nil {
			cli.msg = err.Error()
			return cli, nil
		}
		cli.msg = "You picked up " + loot.GetItems()[cli.checkedIndex].GetName()
		cli.view = MAIN
		cli.checkedIndex = 1337
	}

	return cli, nil
}

func (cli *CLI) updateBattle(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	log.Println(msg)
	return cli, nil
}
