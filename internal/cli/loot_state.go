package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/assets/ascii"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

type lootState struct {
	stage lootStage
}

type lootStage int

const (
	chest lootStage = iota
	list
)

func (ls *lootState) update(cli *CLI, msg tea.KeyMsg) {
	switch ls.stage {
	case chest:
		ls.updateChest(cli, msg)
	case list:
		ls.updateList(cli, msg)
	}

}

func (ls *lootState) view(cli *CLI) (out string) {
	switch ls.stage {
	case chest:
		out = ls.viewChest(cli)
	case list:
		out = ls.viewList(cli)
	}
	return
}

func (is *lootState) updateList(cli *CLI, msg tea.KeyMsg) {
	loot := cli.currentPOI.(*room.Loot)

	switch msg.String() {
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(loot.GetItems())-1 {
			cli.cursor++
		}
	case "enter":
		item := loot.GetItems()[cli.cursor]
		if err := cli.game.Player.PickupItem(item); err != nil {
           cli.msg = err.Error()
           return
        }
		cli.msg = "You picked up " + item.GetName()
		cli.cursor = 0
		cli.view = &mainState{}
	}
}

func (ls *lootState) updateChest(cli *CLI, msg tea.KeyMsg) {

	switch msg.String() {
	case "e":
		cli.msg = ""
		ls.stage = list
	case "s":
		cli.msg = ""
		cli.view = &mainState{}
	}
}
func (ls *lootState) viewChest(cli *CLI) (out string) {

	out = ascii.Chest()
	return
}

func (ls *lootState) viewList(cli *CLI) (out string) {
	cli.msg = ""

	loot := cli.currentPOI.(*room.Loot)
	out = "Select item with up and down arrows, press enter to confirm:\n\n"
	for i, item := range loot.GetItems() {
	cursor := "  [ ]" // no cursor
		if i == cli.cursor {
			cursor = "> [x]" // cursor
		}

		out += fmt.Sprintf("%s %s\n", cursor, item.ToString())
	}
	return
}
