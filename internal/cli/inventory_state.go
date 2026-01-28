package cli

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/item"
	tea "github.com/charmbracelet/bubbletea"
)

type inventoryState struct{}

func (is *inventoryState) getState() State {
	return Inventory
}

func (is *inventoryState) view(cli *CLI) (out string) {
	p := cli.game.Player

	out = "Inventory & Gear Setup\nPress B to return.\n\n"
	out += fmt.Sprintf("Inventory (Weight: %d) - (Use arrows to navigate, X to drop, E to equip)\n",
		p.GetInventoryWeight())

	inventoryItems := p.GetItems()
	if len(inventoryItems) == 0 {
		out += "    Inventory is empty.\n" // 4 spaces instead of tab
	}

	// Inventory items
	for i, item := range inventoryItems {
		cursor := "  [ ]" // no cursor
		if i == cli.cursor {
			cursor = "> [x]" // cursor
		}
		// %-4s reserves 4 characters for cursor, left-aligned
		out += fmt.Sprintf("    %-4s %s\n", cursor, item.ToString())
	}

	out += fmt.Sprintf("\nEquipped (Weight: %d) - (Use arrows to navigate and E to unequip.)\n",
		p.GetEquippedWeight())

	gear := p.GetGear()
	equippedItems := []item.Item{gear.Head, gear.Upperbody, gear.Legs, gear.Feet, gear.Weapon}
	stringSlots := []string{"Head", "Torso", "Legs", "Feet", "Weapon"}

	// Equipped items
	for i, itm := range equippedItems {
		adjustedIndex := i + len(inventoryItems)

		cursor := "  [ ]"
		if adjustedIndex == cli.cursor {
			cursor = "> [x]"
		}

		if itm == nil {
			// %-8s reserves 8 chars for slot name; %-5s reserves 5 for cursor
			out += fmt.Sprintf("%-8s %-5s ______\n", stringSlots[i], cursor)
		} else {
			out += fmt.Sprintf("%-8s %-5s %s\n", stringSlots[i], cursor, itm.ToString())
		}
	}

	out += fmt.Sprintf("\nTotal Weight: %d/%d", p.GetTotalWeight(), p.GetMaxWeight())
	return
}

func (is *inventoryState) update(cli *CLI, msg tea.KeyMsg) {
	items := cli.game.Player.GetItems()
	equippedItemSlots := 5 // Although they may be empty, there will always be 5 rows, one per equippable item slot

	switch msg.String() {
	case "b":
		cli.msg = ""
		cli.view = &mainState{}
		cli.cursor = 0 // When exiting inventory, reset cursor position
		return
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(items)-1+equippedItemSlots {
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
				cli.cursor++ // This is a bit of a work around to the problem where the cursor moves when unequipping an itemm, because of the way that index is counted.
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
