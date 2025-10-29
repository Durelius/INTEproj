package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/item"
	tea "github.com/charmbracelet/bubbletea"
)

type inventoryState struct{}

func (is *inventoryState) view(cli *CLI) (out string) {
	p := cli.game.Player

	out = fmt.Sprintf("\nInventory (Weight: %d):\nUse arrows to navigate, X to drop, E to equip, and B to exit inventory.\n", p.GetInventoryWeight())

	inventoryItems := p.GetItems()
	if len(inventoryItems) == 0 {
		out += "\tInventory is empty.\n"
	}

	for i, item := range inventoryItems {
		cursor := "  [ ]" // no cursor
		if i == cli.cursor {
			cursor = "> [x]" // cursor
		}

		out += fmt.Sprintf("\t%s %s\n", cursor, item.ToString())
	}

	out += fmt.Sprintf("\nEquipped (Weight: %d):\nUse arrows to navigate and E to unequip an item.\n", p.GetEquippedWeight())

	gear := p.GetGear()
	equippedItems := []item.Item{gear.Head, gear.Upperbody, gear.Legs, gear.Feet, gear.Weapon}
	stringSlots := []string{"Head", "Torso", "Legs", "Feet", "Weapon"}

	for i, item := range equippedItems {
		adjustedIndex := i + len(inventoryItems)

		cursor := "  [ ]" // no cursor
		if adjustedIndex == cli.cursor {
			cursor = "> [x]" // cursor
		}

		if item == nil {
			// out += fmt.Sprintf("%s\n",stringSlots[i])
			out += fmt.Sprintf("%s\t%s ______\n", stringSlots[i], cursor)
		} else {
			out += fmt.Sprintf("%s\t%s %s\n", stringSlots[i], cursor, item.ToString())
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
		cli.cursor = 0		// When exiting inventory, reset cursor position
		return
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(items)-1 + equippedItemSlots {
			cli.cursor++
		}
	case "x":
		if CursorIsInIventory(cli) {
			cli.game.Player.DropItem(items[cli.cursor])
		}
	case "e":
		if CursorIsInIventory(cli) {
			cli.game.Player.EquipItem(items[cli.cursor])
		}
		if CursorIsInEquipped(cli) {
			p := cli.game.Player
			
			adjustedCursorIndex := cli.cursor - len(items)
			didUnequipItem := false
			switch adjustedCursorIndex {
			case 0:
				didUnequipItem = p.UnequipHead()
			case 1: 
				didUnequipItem = p.UnequipUpperBody()
			case 2:
				didUnequipItem = p.UnequipLowerBody()
			case 3:
				didUnequipItem = p.UnequipFeet()
			case 4:
				didUnequipItem = p.UnequipWeapon()
			}
			if didUnequipItem {
				cli.cursor ++ 	// This is a bit of a work around to the problem where the cursor moves when unequipping an itemm, because of the way that index is counted. 
			}
		}
	}
}



func CursorIsInIventory(cli *CLI) bool {
	return cli.cursor < len(cli.game.Player.GetItems())
}

func CursorIsInEquipped(cli *CLI) bool {
	return cli.cursor >= len(cli.game.Player.GetItems())
}