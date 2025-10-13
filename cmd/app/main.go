package main

import (
	"fmt"
	"os"

	gs "github.com/Durelius/INTEproj/internal/gamestate"

	tea "github.com/charmbracelet/bubbletea"
)



func main() {
	gamestate := gs.NewDefault();
	p := tea.NewProgram(gamestate)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
