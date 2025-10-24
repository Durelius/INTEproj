package cli

import (
	"github.com/Durelius/INTEproj/internal/assets/ascii"
	"github.com/Durelius/INTEproj/internal/random"
	tea "github.com/charmbracelet/bubbletea"
)

type enemyState struct {
	stage enemyStage
}

type enemyStage int

const (
	encounter enemyStage = iota
	fight
)

func (es *enemyState) view(cli *CLI) (out string) {
	switch es.stage {
	case encounter:

		out = es.viewEncounter(cli)
	case fight:

		out = es.viewFight(cli)
	}
	return

}

func (es *enemyState) update(cli *CLI, msg tea.KeyMsg) {
	switch es.stage {
	case encounter:

		es.updateEncounter(cli, msg)
	case fight:
		es.updateFight(cli, msg) //TODO: implement
	}

}
func (es *enemyState) updateEncounter(cli *CLI, msg tea.KeyMsg) {
	switch msg.String() {
	case "r":
		gamba := random.Int(0, 100)
		if gamba > 75 {
			cli.msg = "You ran away"
			cli.view = &mainState{}
		} else {
			es.stage = fight
		}
	case "f":
		es.stage = fight
	}
}
func (es *enemyState) updateFight(cli *CLI, msg tea.KeyMsg) {
	switch msg.String() {
	case "r":
		cli.msg = "You ran away"
		cli.view = &mainState{}
	case "a":
		//implement attack
	}
}
func (es *enemyState) viewEncounter(cli *CLI) (out string) {
	return ascii.Encounter(cli.game.Player.GetCurrentHealth(), cli.game.Player.GetMaxHealth(), 100, 100)

}
func (es *enemyState) viewFight(cli *CLI) string {

	return ascii.Encounter(cli.game.Player.GetCurrentHealth(), cli.game.Player.GetMaxHealth(), 100, 100)
}
