package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/assets/ascii"
	gs "github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/room"
	"github.com/Durelius/INTEproj/internal/view"
	tea "github.com/charmbracelet/bubbletea"
)

type CLI struct {
	game *gs.GameState
	view  view.View
	msg string
	cursor	int			// Cursor position in the current view
	checkedIndex int	// Index of the currently checked item in lists
	currentPOI room.PointOfInterest
}

func New(game *gs.GameState) CLI {
    return CLI{game: game}
}

func (cli CLI) Init() tea.Cmd { 
	return nil 
}

func (cli CLI) View() string {
	out := cli.getHeaderInfo()

	switch cli.view {
	case view.MAIN:
		out += "Press B to open bag\n"
		out += "\nMap:\n"

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
	case view.LOOT:
		out += ascii.Chest()
	case view.LOOT_LIST:
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
	case view.BAG:
		out += "Inventory:\n\n"
		inv := cli.game.Player.GetInventory()

		for i, item := range inv.GetItems() {
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
	case view.BATTLE:
		out += "\n BATTLE TIME:\n"
	}
	return out
}

func (cli CLI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return cli, tea.Quit
		}
		
		switch cli.view {
		case view.BAG:
			cli.msg = "Press space to toggle items, X to drop and ENTER to consume"
			return cli.updateBag(msg)
		case view.BATTLE:
			return cli.updateBattle(msg)
		case view.LOOT:
			return cli.updateLoot(msg)
		case view.LOOT_LIST:
			return cli.updateLootList(msg)
		default:
			return cli.updateMain(msg)
		}

	}
	return cli, nil
}

func (cli *CLI) getHeaderInfo() string {
	player := cli.game.Player
	room := cli.game.Room
	loc := room.GetPlayerLocation()
	x, y := loc.Get()

	s := fmt.Sprintf("Room: %s (%dx%d)\n", room.GetName(), room.GetWidth(), room.GetHeight())
	s += fmt.Sprintf("Player: %s (%s) HP:%d/%d\n", player.GetName(), player.GetClass(), player.GetCurrentHealth(), player.GetMaxHealth())
	s += fmt.Sprintf("Location: (%d,%d)\n", x,y)
	s += cli.msg + "\n"
	return s
}

func (cli *CLI) updateMain(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	loc := cli.game.Room.GetPlayerLocation()
	x, y := loc.Get()

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
		cli.view = view.BAG
		return cli, nil
	}

	if x < 0 || y < 0 {
		return cli, nil
	}

	poi := cli.game.Room.UsePOI(x, y)
	if poi != nil {
		cli.currentPOI = poi
		switch poi.GetType() {
		case "ENEMY":
			cli.msg = "You encountered an enemy! Prepare to fight!"
			cli.view = view.BATTLE
		case "LOOT":
			cli.msg = "Press E to open the chest, or S to skip!"
			cli.view = view.LOOT
		default:
			cli.msg = poi.GetType()
			cli.view = view.MAIN
		}
	}
	cli.game.Room.SetPlayerLocation(x, y)
	return cli, nil
}

func (cli *CLI) updateBag(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	bag := cli.game.Player.GetInventory()
	items := bag.GetItems()
	switch msg.String() {
	case "b":
		cli.view = view.MAIN
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
			cli.view = view.MAIN
			return cli, nil
		}
		if err := cli.game.Player.PickupItem(items[cli.checkedIndex]); err != nil {
			cli.msg = err.Error()
			return cli, nil
		}
		cli.msg = "You picked up " + items[cli.checkedIndex].GetName()
		cli.view = view.MAIN
		cli.checkedIndex = 1337
	}

	return cli, nil
}

func (cli *CLI) updateLoot(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "e":
		cli.msg = ""
		cli.view = view.LOOT_LIST
	case "s":
		cli.view = view.MAIN
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
			cli.view = view.MAIN
			return cli, nil
		}
		if err := cli.game.Player.PickupItem(loot.GetItems()[cli.checkedIndex]); err != nil {
			cli.msg = err.Error()
			return cli, nil
		}
		cli.msg = "You picked up " + loot.GetItems()[cli.checkedIndex].GetName()
		cli.view = view.MAIN
		cli.checkedIndex = 1337
	}

	return cli, nil
}

func (cli *CLI) updateBattle(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return cli, nil
}
