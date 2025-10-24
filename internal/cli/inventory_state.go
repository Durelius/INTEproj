package cli

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type inventoryState struct{}

func (is *inventoryState) view(cli *CLI) (out string) {
	cli.msg = "Press space to toggle items, X to drop and ENTER to consume"

	out = "Inventory:\n\n"

	for i, item := range cli.game.Player.GetItems() {
		cursor := " " // no cursor
		if i == cli.cursor {
			cursor = ">" // cursor
		}
		checked := "[ ]"
		if cli.checkedIndex == i {
			checked = "[x]"
		}

		out += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
	}
	return
}

func (is *inventoryState) update(cli *CLI, msg tea.KeyMsg) {
	items := cli.game.Player.GetItems()
	switch msg.String() {
	case "b":
		cli.view = &mainState{}
		return
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(items)-1 {
			cli.cursor++
		}
	case " ":
		if cli.checkedIndex == cli.cursor {
			cli.checkedIndex = INTEGER_MAX
		} else {
			cli.checkedIndex = cli.cursor
		}
	case "enter":
		if cli.checkedIndex == INTEGER_MAX {
			cli.msg = "You got nothing..."
			cli.view = &mainState{}
			return
		}
		if err := cli.game.Player.PickupItem(items[cli.checkedIndex]); err != nil {
			cli.msg = err.Error()
			return
		}
		cli.msg = "You picked up " + items[cli.checkedIndex].GetName()
		cli.view = &mainState{}
		cli.checkedIndex = INTEGER_MAX
	}

}
