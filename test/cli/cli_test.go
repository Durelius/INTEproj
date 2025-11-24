package cli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/gamestate"
	tea "github.com/charmbracelet/bubbletea"
)

func CLITest(t *testing.T) {
	gamestate := gamestate.GameState{}

	cli := cli.New(&gamestate)

	p := tea.NewProgram(cli)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
