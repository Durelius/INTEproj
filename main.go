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

	"INTE/projekt/player"
	"INTE/projekt/room"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	player player.Player
	room   *room.Room
	msg    string
}

func initialModel() model {
	p, err := player.New(player.CLASS_ROGUE, "Josh")
	if err != nil {
		log.Fatal(err)
	}
	return model{
		player: p,
		room:   room.STARTING_AREA,
		msg:    "Use arrow keys to move. q to quit.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		loc := m.room.GetPlayerLocation()
		locX, locY := loc.Get()

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if locY > 0 {
				locY--
			}
		case "down":
			if locY < m.room.GetHeight()-1 {
				locY++
			}
		case "left":
			if locX > 0 {
				locX--
			}
		case "right":
			if locX < m.room.GetWidth()-1 {
				locX++
			}
		}
		if locX < 0 || locY < 0 {
			break
		}
		// Update location with new coordinates!
		m.room.SetPlayerLocation(room.NewLocation(locX, locY))
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
		if y > 20 { // Only show a small part for performance
			break
		}
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
