package cli_test

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
)

// func testCLI() {
// 	gamestate := gamestate.GameState{}

// 	cli := cli.New(&gamestate)

// 	p := tea.NewProgram(cli)
// 	if _, err := p.Run(); err != nil {
// 		fmt.Println("Error running TUI:", err)
// 		os.Exit(1)
// 	}
// }

func runGameToMapState() gamestate.GameState {
	p := player.New("test", class.MAGE_STR)
	pois := make(map[room.Location]room.PointOfInterest)
	pois[room.NewLocation(5, 0)] = room.NewLoot()
	pois[room.NewLocation(0, 5)] = enemy.NewGoblin()

	r := room.NewCustomRoom(pois, 10, 10, 1, 1)
	gs := gamestate.New(p, r)
	return *gs
}

func TestCreateBothStates(t *testing.T) {
	runGameToMapState().Player.GetID()
}
