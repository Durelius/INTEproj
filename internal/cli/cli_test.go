package cli

import (
	"testing"

	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/item"
	"github.com/Durelius/INTEproj/internal/player"
	"github.com/Durelius/INTEproj/internal/player/class"
	"github.com/Durelius/INTEproj/internal/room"
	tea "github.com/charmbracelet/bubbletea"
)

func setupTestCLI() *CLI {
	// Initialize a dummy player
	p := player.New("TestHero", class.MAGE_STR)

	// Initialize a dummy room (10x10)
	pois := make(map[room.Location]room.PointOfInterest)
	pois[room.NewLocation(5, 3)] = room.NewLoot()
	pois[room.NewLocation(7, 5)] = enemy.NewGoblin()

	r := room.NewCustomRoom(pois, 10, 10, 1, 1)
	r.SetPlayerLocation(5, 5)
	// Initialize GameState
	gs := gamestate.New(p, r)

	// Return the CLI
	return New(gs)
}

func assertState(t *testing.T, cli *CLI, expectedType string) {
	t.Helper()
	var currentType string
	switch cli.view.(type) {
	case *initialState:
		currentType = "initialState"
	case *mainState:
		currentType = "mainState"
	case *inventoryState:
		currentType = "inventoryState"
	case *lootState:
		currentType = "lootState"
	case *battleState:
		currentType = "battleState"
	default:
		currentType = "unknown"
	}

	if currentType != expectedType {
		t.Errorf("Expected state %s, got %s", expectedType, currentType)
	}
}

//Walking states

func TestWalkingRight(t *testing.T) {
	cli := setupTestCLI()
	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRight})
	loc := cli.game.Room.GetPlayerLocation()
	x, _ := loc.Get()
	if x != 6 {
		t.Errorf("Expected player X to be 6, got %d - Cannot move right", x)
	}
}
func TestWalkingLeft(t *testing.T) {
	cli := setupTestCLI()
	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyLeft})
	loc := cli.game.Room.GetPlayerLocation()
	x, _ := loc.Get()
	if x != 4 {
		t.Errorf("Expected player X to be 4, got %d - Cannot move Left", x)
	}
}

func TestWalkingDown(t *testing.T) {
	cli := setupTestCLI()
	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	loc := cli.game.Room.GetPlayerLocation()
	_, y := loc.Get()
	if y != 6 {
		t.Errorf("Expected player Y to be 4, got %d - Cannot move Down", y)
	}
}
func TestWalkingUp(t *testing.T) {
	cli := setupTestCLI()
	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyUp})
	loc := cli.game.Room.GetPlayerLocation()
	_, y := loc.Get()
	if y != 4 {
		t.Errorf("Expected player Y to be 4, got %d - Cannot move Up", y)
	}
}

//Loot tests

func TestFlow_Loot_PickupAny(t *testing.T) {
	cli := setupTestCLI() // Ensure this setup places a Loot box at (5,4)
	cli.view = &mainState{}

	initialCount := len(cli.game.Player.GetItems())

	// 1. Walk Up into Loot
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyUp})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyUp})

	// 2. Open Chest ('e')
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})

	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyEnter})

	finalCount := len(cli.game.Player.GetItems())
	if finalCount <= initialCount {
		t.Errorf("Expected inventory count to increase, went from %d to %d", initialCount, finalCount)
	}
	assertState(t, cli, "mainState")
}
func TestInventory_CursorBounds(t *testing.T) {
	cli := setupTestCLI()
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})

	//Stress testa och gå ner massa
	for i := 0; i < 10; i++ {
		cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	}

	// Max index for equipment is 4 (5 items: 0,1,2,3,4).
	if cli.cursor > 4 {
		t.Errorf("Cursor went out of bounds: %d", cli.cursor)
	}

	//Stress testa och gå upp massa
	for i := 0; i < 10; i++ {
		cli.view.update(cli, tea.KeyMsg{Type: tea.KeyUp})
	}
	if cli.cursor < 0 {
		t.Errorf("Cursor went out of bounds: %d", cli.cursor)
	}

}
func TestEquipItems(t *testing.T) {
	cli := setupTestCLI()

	cli.game.Player.PickupItem(item.GetItemByName("Leather Cap"))
	cli.game.Player.PickupItem(item.GetItemByName("Padded Vest"))
	cli.game.Player.PickupItem(item.GetItemByName("Cloth Trousers"))
	cli.game.Player.PickupItem(item.GetItemByName("Worn Boots"))
	cli.game.Player.PickupItem(item.GetItemByName("common sword"))

	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	for i := 0; i < 5; i++ {
		cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})
	}
	if cli.cursor != 0 {
		t.Error("All items were not equipped")
	}

}

func TestUnEquipItems(t *testing.T) {
	cli := setupTestCLI()

	cli.game.Player.EquipItem(item.GetItemByName("Leather Cap"))
	cli.game.Player.EquipItem(item.GetItemByName("Padded Vest"))
	cli.game.Player.EquipItem(item.GetItemByName("Cloth Trousers"))
	cli.game.Player.EquipItem(item.GetItemByName("Worn Boots"))
	cli.game.Player.EquipItem(item.GetItemByName("common sword"))
	cli.view = &mainState{}
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})

	//Head
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	//Torso
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	//legs
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	//Feet
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyDown})
	//Weapon
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("e")})

	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("b")})
	if cli.game.Player.GetTotalDefense() != 0 {
		t.Error("Player did not unequip all armour.")
	}
	if cli.game.Player.GetDamage() != 10 {
		t.Error("Player did not unequip their weapon.")
	}
}

//Battle tests

func TestFlow_Battle_EncounterToFight(t *testing.T) {
	cli := setupTestCLI()
	cli.view = &mainState{}

	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRight})
	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRight})

	bs, ok := cli.view.(*battleState)
	if !ok {
		t.Fatalf("Failed to trigger battle state")
	}
	if bs.stage != encounter {
		t.Errorf("Expected Battle Stage 'encounter', got %v", bs.stage)
	}

	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("f")})

	if bs.stage != fight {
		t.Errorf("Expected Battle Stage 'fight' after pressing F, got %v", bs.stage)
	}

	cli.view.update(cli, tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("a")})

	// Assert we are still in battleState (either fighting, victory, or defeat)
	assertState(t, cli, "battleState")
}
