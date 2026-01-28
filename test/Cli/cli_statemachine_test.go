package cli

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/cli"
	"github.com/Durelius/INTEproj/internal/gamestate"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onsi/gomega"
)

func inputName(cli *cli.CLI, s string) {
	for _, c := range s {
		cli.Update(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune{c},
		})
	}
}

func TestCreateCharacterAndRedoName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	gameState := gamestate.GameState{}
	c := cli.New(&gameState)

	// Game should open in the start menu
	g.Expect(c.GetState()).To(gomega.Equal(cli.StartMenu))

	// Select the first option in the start menu, which is to create a new character.
	// Press space to select it, then enter to confirm
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.NameMenu))

	// Enter a character name
	inputName(c, "FoulOyster")
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.ClassMenu))

	// CTRL + C returns the player to the previous menu, IE the name menu
	c.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	g.Expect(c.GetState()).To(gomega.Equal(cli.NameMenu))

	inputName(c, "SnusMumriken123")
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.ClassMenu))

	// Select and confirm the first option, Mage class
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.Map))
}

func TestCreateCharacterAndRedoNameAndGoBackToStartMenuAndLoadSavedCharacter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	gameState := gamestate.GameState{}
	c := cli.New(&gameState)

	// Select the first option in the start menu, which is to create a new character.
	// Press space to select it, then enter to confirm
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.NameMenu))

	// Enter a character name
	inputName(c, "BilboBaggins")
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.ClassMenu))

	// CTRL + C returns the player to the previous menu, IE the name menu
	c.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	g.Expect(c.GetState()).To(gomega.Equal(cli.NameMenu))

	// CTRL + C again returns the player to the start menu
	c.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	g.Expect(c.GetState()).To(gomega.Equal(cli.StartMenu))

	// Selects to create a new character again
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.NameMenu))

	// CTRL + C again returns the player to the start menu
	c.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	g.Expect(c.GetState()).To(gomega.Equal(cli.StartMenu))

	// Arrow down scrolls to a previously saved character, space selects and enter confirms selection
	c.Update(tea.KeyMsg{Type: tea.KeyDown})
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.Map))
}

func TestLoadSavedCharacterImmediately(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	gameState := gamestate.GameState{}
	c := cli.New(&gameState)

	// Game should open in the start menu
	g.Expect(c.GetState()).To(gomega.Equal(cli.StartMenu))

	// Arrow down scrolls to a previously saved character, space selects and enter confirms selection
	c.Update(tea.KeyMsg{Type: tea.KeyDown})
	c.Update(tea.KeyMsg{Type: tea.KeySpace})
	c.Update(tea.KeyMsg{Type: tea.KeyEnter})
	g.Expect(c.GetState()).To(gomega.Equal(cli.Map))
}
