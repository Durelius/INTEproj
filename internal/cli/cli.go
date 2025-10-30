package cli

import (
	"fmt"

	gs "github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

const INTEGER_MAX int = int(^uint(0) >> 1) // unsigned int with inverse bits with 1 bit shift to get max signed int

// The CLI struct is the main model for the command line interface.
// It listens for user input and updates the gamestate, as well as the current view.
type CLI struct {
	game         *gs.GameState
	msg          string
	cursor       int // Cursor position in the current view
	checkedIndex int
	currentPOI   room.PointOfInterest
	view         cliState
}

type cliState interface {
	view(*CLI) string
	update(*CLI, tea.KeyMsg)
}

func New(game *gs.GameState) *CLI {
	return &CLI{game: game, view: &initialState{}, checkedIndex: INTEGER_MAX}
}

func (cli *CLI) Init() tea.Cmd {
	return nil
}

func (cli *CLI) generateMapView() string {
	out := "\nMap:\n"
	loc := cli.game.Room.GetPlayerLocation()
	locX, locY := loc.Get()
	poiMap := cli.game.Room.GetPOI()

	for y := 0; y < cli.game.Room.GetHeight(); y++ {
		for x := 0; x < cli.game.Room.GetWidth(); x++ {
			curLoc := room.NewLocation(x, y)
			if x == locX && y == locY {
				out += "\033[34m@\033[0m" // blue "@"
			} else if poi, ok := poiMap[curLoc]; ok {
				switch cli.game.Room.GetPOI()[curLoc].GetType() {
				case "LOOT", "ENEMY":
					out += "\033[33m?\033[0m"
				case "EXIT":
					exit := poi.(*room.Exit)
					if exit.IsLocked(cli.game.Room) {
						out += "\033[31m#\033[0m"
					} else {
						out += "\033[32m#\033[0m"
					}
				}
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return out
}

func (cli *CLI) View() (out string) {

	out += cli.view.view(cli)

	return
}

// Update reads a message (user input) and updates the view accordingly.
func (cli *CLI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return cli, tea.Quit
		}

		cli.view.update(cli, msg)
		return cli, nil
	}
	return cli, nil
}

// Fetches all relevant information which is always displayed at the top of the CLI
func (cli *CLI) getHeaderInfo() string {
	player := cli.game.Player
	room := cli.game.Room
	loc := room.GetPlayerLocation()
	x, y := loc.Get()

	s := fmt.Sprintf("Room: %d (%dx%d)\n", room.GetLevel(), room.GetWidth(), room.GetHeight())
	s += fmt.Sprintf("Player: %s (%v) HP:%d/%d\n", player.GetName(), player.GetClassName(), player.GetCurrentHealth(), player.GetMaxHealth())
	s += fmt.Sprintf("Location: (%d,%d)\n", x, y)
	s += cli.msg + "\n"
	return s
}
