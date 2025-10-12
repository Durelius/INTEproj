// package main

// import (
// 	"INTE/projekt/enemy"
// 	"INTE/projekt/player"
// 	"INTE/projekt/room"
// 	"log"
// )

// func main() {
// 	player, err := player.New(player.CLASS_PALADIN, "josh")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(player.GetHealth())

// 	enemy, err := enemy.New(enemy.CLASS_GOBLIN, "gobgosh")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	newHP, err := player.Attack(enemy)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Print(newHP)
// 	log.Println(room.STARTING_AREA)
// }

package main

import (
	"fmt"
	"log"
	"os"

	"INTE/projekt/ascii"
	"INTE/projekt/player"
	"INTE/projekt/room"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	player       player.Player
	room         *room.Room
	msg          string
	view         view
	cursor       int
	checkedIndex int
	currentPOI   room.PointOfInterest
}

type view string

const (
	VIEW_MAIN      view = "MAIN"
	VIEW_BATTLE    view = "BATTLE"
	VIEW_BAG       view = "BAG"
	VIEW_LOOT      view = "LOOT"
	VIEW_LOOT_LIST view = "LOOT_LIST"
)

const (
	NULL_INDEX int = 1337
)

func initialModel() model {
	p, err := player.New(player.CLASS_ROGUE, "Josh")
	if err != nil {
		log.Fatal(err)
	}
	return model{
		player:       p,
		room:         room.STARTING_AREA,
		msg:          "Use arrow keys to move. q to quit.",
		view:         VIEW_MAIN,
		checkedIndex: NULL_INDEX,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) UpdateMain(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "up":
		if y > 0 {
			y--
		}
	case "down":
		if y < m.room.GetHeight()-1 {
			y++
		}
	case "left":
		if x > 0 {
			x--
		}
	case "right":
		if x < m.room.GetWidth()-1 {
			x++
		}
	case "b":
		m.view = VIEW_BAG
		return m, nil
	}

	if x < 0 || y < 0 {
		return m, nil
	}

	poi := m.room.UsePOI(x, y)
	if poi != nil {
		m.currentPOI = poi
		switch poi.GetType() {
		case "ENEMY":
			m.msg = "You encountered an enemy! Prepare to fight!"
			m.view = VIEW_BATTLE
		case "LOOT":
			m.msg = "Press E to open the chest, or S to skip!"
			m.view = VIEW_LOOT
		default:
			m.msg = poi.GetType()
			m.view = VIEW_MAIN
		}
	}
	m.room.SetPlayerLocation(x, y)
	return m, nil
}
func (m model) UpdateBag(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	bag := m.player.GetBag()
	items := bag.GetItems()
	switch msg.String() {
	case "b":
		m.view = VIEW_MAIN
		return m, nil
	case "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down":
		if m.cursor < len(items)-1 {
			m.cursor++
		}
	case " ":
		if m.checkedIndex == m.cursor {
			m.checkedIndex = NULL_INDEX
		} else {
			m.checkedIndex = m.cursor
		}
	case "enter":
		if m.checkedIndex == NULL_INDEX {
			m.msg = "You got nothing..."
			m.view = VIEW_MAIN
			return m, nil
		}
		if err := m.player.PickupItem(items[m.checkedIndex]); err != nil {
			m.msg = err.Error()
			return m, nil
		}
		m.msg = "You picked up " + items[m.checkedIndex].GetName()
		m.view = VIEW_MAIN
		m.checkedIndex = NULL_INDEX
	}

	return m, nil
}
func (m model) UpdateLoot(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "e":
		m.msg = ""
		m.view = VIEW_LOOT_LIST
	case "s":
		m.view = VIEW_MAIN
	}

	return m, nil
}
func (m model) UpdateLootList(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {
	loot := m.currentPOI.(*room.Loot)
	switch msg.String() {
	case "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down":
		if m.cursor < len(loot.GetItems())-1 {
			m.cursor++
		}
	case " ":
		if m.checkedIndex == m.cursor {
			m.checkedIndex = NULL_INDEX
		} else {
			m.checkedIndex = m.cursor
		}
	case "enter":

		if m.checkedIndex == NULL_INDEX {
			m.msg = "You got nothing..."
			m.view = VIEW_MAIN
			return m, nil
		}
		if err := m.player.PickupItem(loot.GetItems()[m.checkedIndex]); err != nil {
			m.msg = err.Error()
			return m, nil
		}
		m.msg = "You picked up " + loot.GetItems()[m.checkedIndex].GetName()
		m.view = VIEW_MAIN
		m.checkedIndex = NULL_INDEX
	}

	return m, nil
}
func (m model) UpdateBattle(msg tea.KeyMsg, x, y int) (tea.Model, tea.Cmd) {
	return m, nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		loc := m.room.GetPlayerLocation()
		x, y := loc.Get()

		switch m.view {
		case VIEW_BAG:
			m.msg = "Press space to toggle items, X to drop and ENTER to consume"
			return m.UpdateBag(msg)
		case VIEW_BATTLE:
			return m.UpdateBattle(msg, x, y)
		case VIEW_LOOT:
			return m.UpdateLoot(msg)
		case VIEW_LOOT_LIST:
			return m.UpdateLootList(msg, x, y)
		default:
			return m.UpdateMain(msg, x, y)
		}

	}
	return m, nil
}

func (m model) View() string {
	loc := m.room.GetPlayerLocation()
	x, y := loc.Get()
	s := fmt.Sprintf("Room: %s (%dx%d)\n", m.room.GetName(), m.room.GetWidth(), m.room.GetHeight())
	s += fmt.Sprintf("Player: %s (%s) HP:%d\n", m.player.GetName(), m.player.GetClass(), m.player.GetHealth())
	s += fmt.Sprintf("Location: (%d,%d)\n", x, y)
	s += m.msg + "\n"

	switch m.view {
	case VIEW_MAIN:
		s += "Press B to open bag\n"
		s += "\nMap:\n"
		for y := 0; y < m.room.GetHeight(); y++ {
			for x := 0; x < m.room.GetWidth(); x++ {
				locX, locY := loc.Get()
				if x == locX && y == locY {
					s += "@"
				} else {
					s += "."
				}
			}
			s += "\n"

		}
	case VIEW_LOOT:
		s += ascii.Chest()
	case VIEW_LOOT_LIST:
		loot := m.currentPOI.(*room.Loot)
		s += "Select item from (space to toggle, enter to confirm):\n\n"
		for i, item := range loot.GetItems() {
			cursor := " " // no cursor
			if i == m.cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if m.checkedIndex == i {
				checked = "[x]"
			}

			s += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case VIEW_BAG:
		s += "Inventory:\n\n"
		bag := m.player.GetBag()

		for i, item := range bag.GetItems() {
			cursor := " " // no cursor
			if i == m.cursor {
				cursor = ">" // cursor
			}
			checked := "[ ]"
			if m.checkedIndex == i {
				checked = "[x]"
			}

			s += fmt.Sprintf("%s %s %s\n", cursor, checked, item.ToString())
		}
	case VIEW_BATTLE:
		s += "\n BATTLE TIME:\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
