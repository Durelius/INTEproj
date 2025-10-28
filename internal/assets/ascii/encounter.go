package ascii

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/enemy"
)

// Encounter returns an ASCII battle screen
func Encounter(e enemy.Enemy) string {
	return fmt.Sprintf(`
A %s has appeared! Press R to attempt to run away, or F to accept your destiny.
───────────────────────────────
|                               |
|                               |
|                               |
|                               |
|                               |
|            (◕ ◕)              |
|              /▌\              |
|             /  \              |
|                               |
|                               |
|                               |
|                               |
───────────────────────────────
`, e.GetEnemyType())
}
