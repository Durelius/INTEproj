package main

import (
	"fmt"
	"os"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/gamestate"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	gamestate := gamestate.GameState{}

	cli := cli.New(&gamestate)

	p := tea.NewProgram(cli)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
