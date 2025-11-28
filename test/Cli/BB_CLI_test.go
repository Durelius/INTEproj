package cli

import (
	"os"
	"strings"
	"testing"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

func goFromInitialToMainstate(t *testing.T) *cli.CLI {
	if err := os.MkdirAll("savefiles", 0755); err != nil {
		t.Fatal(err)
	}
	model := cli.New(gamestate.New(player.New("test", class.MAGE_STR), room.NewRandomRoom()))
	model.Update(tea.KeyMsg{Type: tea.KeySpace}) // Kryssa i [x]
	model.Update(tea.KeyMsg{Type: tea.KeyEnter}) // Bekräfta

	// 2. Name Input: Skriv "Hero", sen Enter
	model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("d")})
	model.Update(tea.KeyMsg{Type: tea.KeyEnter})

	// 3. Choose Class: Första klassen är förvald.
	// Markera med Space, sen Enter.
	model.Update(tea.KeyMsg{Type: tea.KeySpace}) // Kryssa i [x]
	model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	return model
}

func TestCorrectStringMapOutPut(t *testing.T) {
	cli := goFromInitialToMainstate(t)

	arr2 := [5]string{"Mage", "Level", "HP", "STATS", "PLAYER"}
	s := ""
	view := cli.View()

	for i := 0; i < len(arr2); i++ {
		if !strings.Contains(view, arr2[i]) {
			s += "CLI does not contain " + arr2[i] + " in output\n"
		}
	}

	if s != "" {
		t.Errorf("%s\nView:\n%s", s, view)
	}
}
