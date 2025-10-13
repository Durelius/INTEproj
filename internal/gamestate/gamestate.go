package gamestate

import (
	"fmt"
	"log"

	"github.com/Durelius/INTEproj/assets/ascii"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/room"
	"github.com/Durelius/INTEproj/internal/view"
	tea "github.com/charmbracelet/bubbletea"
)

// GameState holds all information about the current state of the game
type GameState struct {
	Player       player.Player			// The player character
	Room         *room.Room				// The current room the player is in
	Msg          string					// Message to display to the player
	GameView         view.View					// Current view (main, battle, bag, loot, etc.)
	Cursor       int					// Cursor position in the current view
	CheckedIndex int					// Index of the currently checked item in lists
	CurrentPOI   room.PointOfInterest	// Currently selected point of interest in the room
}

// New creates a new GameState instance
func New(p player.Player, room *room.Room, msg string, view view.View, cursor int, checkedIndex int, currentPoi room.PointOfInterest) *GameState{
	return &GameState{
		Player: p,
		Room: room,
		Msg: msg,
		GameView: view,
		Cursor: cursor,
		CheckedIndex: checkedIndex,
		CurrentPOI: currentPoi,
	}
}

// Currently this method serves to bypass the limitations of Golang not allowing optional arguments.
// This helps us get a base Game going, where player is hard coded to Josh.
func NewDefault () *GameState {
	p, err := player.New(player.CLASS_ROGUE, "Josh")
	if err != nil {
		log.Fatal(err)
	}

	return &GameState{
		Player: p,
		Room: room.STARTING_AREA,
		Msg: "Use arrow keys to move. q to quit.",
		GameView: view.MAIN,
		Cursor: 0, 
		CheckedIndex: 1337, // Temporary unreachable index
	}
}


func (*GameState) Init() tea.Cmd {
	return nil
}

func (gs *GameState) View() string {
	loc := gs.Room.GetPlayerLocation()
	x, y := loc.Get()
	s := fmt.Sprintf("Room: %s (%dx%d)\n", gs.Room.GetName(), gs.Room.GetWidth(), gs.Room.GetHeight())
	s += fmt.Sprintf("Player: %s (%s) HP:%d\n", gs.Player.GetName(), gs.Player.GetClass(), gs.Player.GetHealth())
	s += fmt.Sprintf("Location: (%d,%d)\n", x, y)
	s += gs.Msg + "\n"

	switch gs.GameView {
	case view.MAIN:
		s += "Press B to open bag\n"
		s += "\nMap:\n"
		for y := 0; y < gs.Room.GetHeight(); y++ {
			for x := 0; x < gs.Room.GetWidth(); x++ {
				locX, locY := loc.Get()
				if x == locX && y == locY {
					s += "@"
				} else {
					s += "."
				}
			}
			s += "\n"

		}
	case view.LOOT:
		s += ascii.Chest()
	case view.LOOT_LIST:
		loot := gs.CurrentPOI.(*room.Loot)
		s += "Select item from (space to toggle, enter to confirm):\n\n"
		for i, item := range loot.GetItems() {
			cursor := " " // no cursor
			if i == gs.Cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if gs.CheckedIndex == i {
				checked = "[x]"
			}

			s += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case view.BAG:
		s += "Inventory:\n\n"
		bag := gs.Player.GetBag()

		for i, item := range bag.GetItems() {
			cursor := " " // no cursor
			if i == gs.Cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if gs.CheckedIndex == i {
				checked = "[x]"
			}

			s += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case view.BATTLE:
		s += "\n BATTLE TIME:\n"
	}
	return s
}

func (gs *GameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return gs, tea.Quit
		}
		loc := gs.Room.GetPlayerLocation()
		x, y := loc.Get()

		switch gs.GameView {
		case view.BAG:
			gs.Msg = "Press space to toggle items, X to drop and ENTER to consume"
			return gs.UpdateBag(msg)
		case view.BATTLE:
			return gs.UpdateBattle(msg, x, y)
		case view.LOOT:
			return gs.UpdateLoot(msg)
		case view.LOOT_LIST:
			return gs.UpdateLootList(msg, x, y)
		default:
			return gs.UpdateMain(msg, x, y)
		}

	}
	return gs, nil
}

func (gs *GameState) UpdateMain(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "up":
		if y > 0 {
			y--
		}
	case "down":
		if y < gs.Room.GetHeight()-1 {
			y++
		}
	case "left":
		if x > 0 {
			x--
		}
	case "right":
		if x < gs.Room.GetWidth()-1 {
			x++
		}
	case "b":
		gs.GameView = view.BAG
		return gs, nil
	}

	if x < 0 || y < 0 {
		return gs, nil
	}

	poi := gs.Room.UsePOI(x, y)
	if poi != nil {
		gs.CurrentPOI = poi
		switch poi.GetType() {
		case "ENEMY":
			gs.Msg = "You encountered an enemy! Prepare to fight!"
			gs.GameView = view.BATTLE
		case "LOOT":
			gs.Msg = "Press E to open the chest, or S to skip!"
			gs.GameView = view.LOOT
		default:
			gs.Msg = poi.GetType()
			gs.GameView = view.MAIN
		}
	}
	gs.Room.SetPlayerLocation(x, y)
	return gs, nil
}

func (gs *GameState) UpdateBag(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	bag := gs.Player.GetBag()
	items := bag.GetItems()
	switch msg.String() {
	case "b":
		gs.GameView = view.MAIN
		return gs, nil
	case "up":
		if gs.Cursor > 0 {
			gs.Cursor--
		}
	case "down":
		if gs.Cursor < len(items)-1 {
			gs.Cursor++
		}
	case " ":
		if gs.CheckedIndex == gs.Cursor {
			gs.CheckedIndex = 1337
		} else {
			gs.CheckedIndex = gs.Cursor
		}
	case "enter":
		if gs.CheckedIndex == 1337 {
			gs.Msg = "You got nothing..."
			gs.GameView = view.MAIN
			return gs, nil
		}
		if err := gs.Player.PickupItem(items[gs.CheckedIndex]); err != nil {
			gs.Msg = err.Error()
			return gs, nil
		}
		gs.Msg = "You picked up " + items[gs.CheckedIndex].GetName()
		gs.GameView = view.MAIN
		gs.CheckedIndex = 1337
	}

	return gs, nil
}


func (gs *GameState) UpdateLoot(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "e":
		gs.Msg = ""
		gs.GameView = view.LOOT_LIST
	case "s":
		gs.GameView = view.MAIN
	}

	return gs, nil
}


func (gs *GameState) UpdateLootList(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {
	loot := gs.CurrentPOI.(*room.Loot)
	switch msg.String() {
	case "up":
		if gs.Cursor > 0 {
			gs.Cursor--
		}
	case "down":
		if gs.Cursor < len(loot.GetItems())-1 {
			gs.Cursor++
		}
	case " ":
		if gs.CheckedIndex == gs.Cursor {
			gs.CheckedIndex = 1337
		} else {
			gs.CheckedIndex = gs.Cursor
		}
	case "enter":

		if gs.CheckedIndex == 1337 {
			gs.Msg = "You got nothing..."
			gs.GameView = view.MAIN
			return gs, nil
		}
		if err := gs.Player.PickupItem(loot.GetItems()[gs.CheckedIndex]); err != nil {
			gs.Msg = err.Error()
			return gs, nil
		}
		gs.Msg = "You picked up " + loot.GetItems()[gs.CheckedIndex].GetName()
		gs.GameView = view.MAIN
		gs.CheckedIndex = 1337
	}

	return gs, nil
}

func (gs *GameState) UpdateBattle(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {
	return gs, nil
}