package cli

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestMoveLeft(t *testing.T) {
	gs := RunGameToInitialState()
	model := CreateCharacterInMainState(t, gs)
	startLoc := gs.Room.GetPlayerLocation()
	startX, _ := startLoc.Get()

	expectedX := startX - 1

	msg := tea.KeyMsg{Type: tea.KeyLeft}
	model.Update(msg)
	var endLocation = gs.Room.GetPlayerLocation()
	endX, _ := endLocation.Get()

	if expectedX == endX {
		t.Errorf("The player did not move to the left, expected Location: %d. Was %d ", expectedX, endX)
	}
}
func TestMoveRight(t *testing.T) {
	gs := RunGameToInitialState()
	model := CreateCharacterInMainState(t, gs)
	startLoc := gs.Room.GetPlayerLocation()
	startX, _ := startLoc.Get()

	expectedX := startX + 1

	msg := tea.KeyMsg{Type: tea.KeyRight}
	model.Update(msg)
	var endLocation = gs.Room.GetPlayerLocation()
	endX, _ := endLocation.Get()

	if expectedX == endX {
		t.Errorf("The player did not move to the Right, expected Location: %d. Was %d ", expectedX, endX)
	}
}
func TestMoveUp(t *testing.T) {
	gs := RunGameToInitialState()
	model := CreateCharacterInMainState(t, gs)
	startLoc := gs.Room.GetPlayerLocation()
	_, startY := startLoc.Get()

	expectedY := startY - 1

	msg := tea.KeyMsg{Type: tea.KeyUp}
	model.Update(msg)
	var endLocation = gs.Room.GetPlayerLocation()
	_, endY := endLocation.Get()

	if expectedY == endY {
		t.Errorf("The player did not move to the left, expected Location: %d. Was %d ", expectedY, endY)
	}
}
func TestMoveDown(t *testing.T) {
	gs := RunGameToInitialState()
	model := CreateCharacterInMainState(t, gs)
	startLoc := gs.Room.GetPlayerLocation()
	_, startY := startLoc.Get()

	expectedY := startY + 1

	msg := tea.KeyMsg{Type: tea.KeyDown}
	model.Update(msg)
	var endLocation = gs.Room.GetPlayerLocation()
	_, endY := endLocation.Get()

	if expectedY == endY {
		t.Errorf("The player did not move to the left, expected Location: %d. Was %d ", expectedY, endY)
	}
}
