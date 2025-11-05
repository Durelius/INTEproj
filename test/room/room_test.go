package roomtest

import (
	"testing"
	"time"

	"github.com/Durelius/INTEproj/internal/gamestate"
	"github.com/Durelius/INTEproj/internal/room"
)

// We should add this to a profiler somehow and do some additional stress testing
func TestRoomStress(t *testing.T) {
	gs := gamestate.NewDefault()
	start := time.Now()
	gs.Room = room.NewRandomRoom(room.NewLocation(0, 0), 10000, 10000, 5000)
	duration := time.Since(start)
	maxTime := time.Second * 10
	if duration > maxTime {
		t.Errorf("Execution took longer than expected %v, Actual: %v", maxTime, duration)
	}
	t.Logf("Execution time to create a big room: %v seconds", duration.Seconds())

}
