package cli

import (
	"os"

	"github.com/Durelius/INTEproj/internal/assets/ascii"
	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/random"
	tea "github.com/charmbracelet/bubbletea"
)

type battleState struct {
	stage battleStage
	enemy enemy.Enemy
}

type battleStage int

const (
	encounter battleStage = iota
	fight
	defeat
	victory
)

func (bs *battleState) getState() State {
	return Battle
}

func (bs *battleState) view(cli *CLI) (out string) {
	switch bs.stage {
	case encounter:
		out = bs.viewEncounter(cli)
	case fight:
		out = bs.viewFight(cli)
	case victory:
		out = bs.viewVictory(cli)
	case defeat:
		out = bs.viewDefeat(cli)
	}
	return
}

func (bs *battleState) update(cli *CLI, msg tea.KeyMsg) {
	switch bs.stage {
	case encounter:
		bs.updateEncounter(cli, msg)
	case fight:
		bs.updateFight(cli, msg)
	case victory:
		bs.updateVictory(cli, msg)
	case defeat:
		bs.updateDefeat(cli, msg)
	}

}

func (bs *battleState) viewEncounter(cli *CLI) (out string) {
	return ascii.Encounter(bs.enemy)
}

func (bs *battleState) updateEncounter(cli *CLI, msg tea.KeyMsg) {
	switch msg.String() {
	case "r":
		gamba := random.Int(0, 100)
		if gamba > 20 {
			cli.msg = "You ran away"
			cli.view = &mainState{}
		} else {
			cli.msg = "You're not going anywhere, prepare to fight!"
			cli.game.InitiateBattle(bs.enemy)
			bs.stage = fight
		}
	case "f":
		cli.game.InitiateBattle(bs.enemy)
		bs.stage = fight
	}
}

func (bs *battleState) viewFight(cli *CLI) string {
	return ascii.Fight(cli.game.Battle)
}

func (bs *battleState) updateFight(cli *CLI, msg tea.KeyMsg) {
	b := cli.game.Battle

	switch msg.String() {

	case "a":
		b.ProgressFight()

		if b.GetStatus() == battle.Victory {
			bs.stage = victory
		}
		if b.GetStatus() == battle.Defeat {
			bs.stage = defeat
		}
	}
}

func (bs *battleState) viewVictory(cli *CLI) string {
	return ascii.Victory()
}

// Should have logic here to get loot
func (bs *battleState) updateVictory(cli *CLI, msg tea.KeyMsg) {

	switch msg.String() {
	case "c":
		cli.view = &mainState{}
	}
}

func (bs *battleState) viewDefeat(cli *CLI) string {
	return ascii.Defeat()
}

func (bs *battleState) updateDefeat(cli *CLI, msg tea.KeyMsg) {

	switch msg.String() { // Exit program on any key press
	default:
		os.Exit(0)
	}
}
