package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

type initialState struct {
	stage         initialStage
	newPlayerName string
}

const (
	NEW_STR = "New"
)

type initialStage int

const (
	initial initialStage = iota
	load
	new
	chooseClass
)

func (is *initialState) menuItemsDisplay(cli *CLI) []string {
	menuItems := []string{NEW_STR}
	savefiles, err := cli.game.GetSaveFiles()
	if err != nil {
		return menuItems
	}
	for i, filename := range savefiles {
		formatted := fmt.Sprintf("%s - %d", strings.Split(filename, "_")[0], +i+1)
		menuItems = append(menuItems, formatted)
	}
	return menuItems
}
func (is *initialState) menuItems(cli *CLI) []string {
	menuItems := []string{NEW_STR}
	savefiles, err := cli.game.GetSaveFiles()
	if err != nil {
		return menuItems
	}

	return append(menuItems, savefiles...)
}

func (is *initialState) view(cli *CLI) (out string) {
	switch is.stage {
	case initial:
		out = is.viewInitial(cli)
	case new:
		out = is.viewNew(cli)
	case chooseClass:
		out = is.viewChooseClass(cli)
	}
	return

}

func (is *initialState) update(cli *CLI, msg tea.KeyMsg) {
	switch is.stage {
	case initial:
		is.updateInitial(cli, msg)
	case new:
		is.updateNew(cli, msg)
	case chooseClass:
		is.updateChooseClass(cli, msg)
	}

}
func (is *initialState) viewInitial(cli *CLI) (out string) {
	cli.msg = "Load a save file or start from scratch"

	for i, menuItem := range is.menuItemsDisplay(cli) {
		cursor := " " // no cursor
		if i == cli.cursor {
			cursor = ">" // cursor
		}
		checked := "[ ]"
		if cli.checkedIndex == i {
			checked = "[x]"
		}

		out += fmt.Sprintf("%s %s %s\n", cursor, checked, menuItem)
	}
	return
}

// TODO
func (is *initialState) viewNew(cli *CLI) (out string) {
	cli.msg = "Choose character name, enter to save"
	out += fmt.Sprintf("Player name: %s\n", is.newPlayerName)

	return
}

func (is *initialState) viewChooseClass(cli *CLI) (out string) {
	cli.msg = "Choose character class with arrows and space, enter to choose"
	for i, class := range class.CLASS_LIST {
		cursor := " " // no cursor
		if i == cli.cursor {
			cursor = ">" // cursor
		}
		checked := "[ ]"
		if cli.checkedIndex == i {
			checked = "[x]"
		}
		out += fmt.Sprintf("%s %s %s\n", cursor, checked, class)

	}
	out += fmt.Sprintf("Player name: %s\n", is.newPlayerName)

	return
}

func (is *initialState) updateInitial(cli *CLI, msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(is.menuItems(cli))-1 {
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
			cli.msg = "Nothing ever happens"
			return
		}
		if is.menuItems(cli)[cli.checkedIndex] == NEW_STR {
			cli.msg = "Choose character name, enter to save"
			is.stage = new
			return
		}

		gs, err := gamestate.LoadSaveFile(is.menuItems(cli)[cli.checkedIndex])
		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}
		cli.game = gs
		cli.view = &mainState{}

		cli.checkedIndex = INTEGER_MAX
	}

}
func (is *initialState) updateNew(cli *CLI, msg tea.KeyMsg) {
	nameLen := len(is.newPlayerName)
	if nameLen > 0 {
		nameLen -= 1
	}
	switch msg.String() {
	case "backspace":
		is.newPlayerName = is.newPlayerName[:nameLen]
		return
	case "enter":
		cli.checkedIndex = INTEGER_MAX
		cli.cursor = 0
		cli.msg = "Choose character class with arrows and space, enter to choose"
		is.stage = chooseClass
		return
	}
	is.newPlayerName += msg.String()

}
func (is *initialState) updateChooseClass(cli *CLI, msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		if cli.cursor > 0 {
			cli.cursor--
		}
	case "down":
		if cli.cursor < len(class.CLASS_LIST)-1 {
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
			cli.msg = "Nothing ever happens"
			return
		}
		cli.game = gamestate.New(player.New(is.newPlayerName, class.CLASS_LIST[cli.checkedIndex]), room.NewRandomRoom(room.NewLocation(0, 0), 25, 50))
		if err := cli.game.SaveToFile(); err != nil {
			log.Fatal(err)
		}
		cli.view = &mainState{}

		cli.checkedIndex = INTEGER_MAX
	}

}
