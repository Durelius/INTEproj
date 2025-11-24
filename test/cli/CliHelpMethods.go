package cli

import (
	"os"
	"testing"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

func RunGameToInitialState() *gamestate.GameState {
	p := player.New("test", class.MAGE_STR)
	pois := make(map[room.Location]room.PointOfInterest)
	pois[room.NewLocation(5, 0)] = room.NewLoot()
	pois[room.NewLocation(0, 5)] = enemy.NewGoblin()

	r := room.NewCustomRoom(pois, 10, 10, 1, 1)
	r.SetPlayerLocation(5, 5)

	// Returnera pekaren direkt
	return gamestate.New(p, r)
}
func CreateCharacterInMainState(t *testing.T, gs *gamestate.GameState) *cli.CLI {
	if err := os.MkdirAll("savefiles", 0755); err != nil {
		t.Fatal(err)
	}
	model := cli.New(gs)
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
