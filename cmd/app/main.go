package main

import (
	"fmt"
	"os"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/gamestate"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create default initial gamestate. This should be replaced with a proper character creation,
	// as well as the option to load a previously saved game.
	gamestate := gamestate.GameState{}

	cli := cli.New(&gamestate)

	p := tea.NewProgram(cli)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
