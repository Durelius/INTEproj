package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/item"
	tea "github.com/charmbracelet/bubbletea"
)

type inventoryState struct{}

func (is *inventoryState) view(cli *CLI) (out string) {
	p := cli.game.Player

	out = fmt.Sprintf("\nInventory (Weight: %d):\nPress SPACE to toggle items, X to drop and E to equip.\n", p.GetInventoryWeight())

	inventoryItems := p.GetItems()
	if len(inventoryItems) == 0 {
		out += "\tInventory is empty.\n"
	}

	for i, item := range inventoryItems {
		cursor := " " // no cursor
		if i == cli.cursor {
			cursor = ">" // cursor
		}
		checked := "[ ]"
		if cli.checkedIndex == i {
			checked = "[x]"
		}

		out += fmt.Sprintf("\t%s %s %s\n", cursor, checked, item.ToString())
	}

	out += fmt.Sprintf("\nEquipped (Weight: %d):\nPress SPACE to toggle items and E to unequip.\n", p.GetEquippedWeight())

	gear := p.GetGear()
	equippedItems := []item.Item{gear.Head, gear.Upperbody, gear.Legs, gear.Feet, gear.Weapon}
	stringSlots := []string{"Head", "Torso", "Legs", "Feet", "Weapon"}

	for i, item := range equippedItems {
		adjustedIndex := i + len(inventoryItems)

		cursor := " " // no cursor
		if  cli.cursor == adjustedIndex{
			cursor = ">" // cursor
		}
		checked := "[ ]"
		if cli.checkedIndex == adjustedIndex {
			checked = "[x]"
		}

		if item == nil {
			// out += fmt.Sprintf("%s\n",stringSlots[i])
			out += fmt.Sprintf("%s\t%s %s ______\n", stringSlots[i], cursor, checked)
		} else {
			out += fmt.Sprintf("%s\t%s %s\n", cursor, checked, item.ToString())
		}
	}

	out += fmt.Sprintf("\nTotal Weight: %d/%d", p.GetTotalWeight(), p.GetMaxWeight())

	return
}

func (is *inventoryState) update(cli *CLI, msg tea.KeyMsg) {
	items := cli.game.Player.GetItems()
	equippedItemSlots := 5	// Although they may be empty, there will always be 5 rows, one per equippable item slot

	switch msg.String() {
	case "b":
		cli.view = &mainState{}
		cli.cursor = 0					// When exiting inventory, set cursor position to zero
		cli.checkedIndex = INTEGER_MAX  // When exiting inventory, uncheck
		return
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(items)-1 + equippedItemSlots {
			cli.cursor++
		}
	case " ":
		if cli.checkedIndex == cli.cursor {
			cli.checkedIndex = INTEGER_MAX
		} else {
			cli.checkedIndex = cli.cursor
		}
	case "e":
		
	}

}
