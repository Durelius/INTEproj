package ascii

import "strings"

// Encounter returns an ASCII battle screen with proportional HP bars
func Encounter(playerHP, playerMaxHP, enemyHP, enemyMaxHP int) string {
	out := `
───────────────────────────────
|                        (◕ ◕)  |
|                         /▌\   |
|                        /  \   |
|                               |
|                               |
|                               |
|                               |
|                               |
|   (\_/)                       |
|   (o_o)                       |
|   /| |\                       |
|  / | | \                      |
───────────────────────────────
`

	const barLength = 10

	// Player HP bar
	playerBlocks := int(float64(playerHP) / float64(playerMaxHP) * barLength)
	if playerBlocks < 0 {
		playerBlocks = 0
	}
	if playerBlocks > barLength {
		playerBlocks = barLength
	}
	playerBar := strings.Repeat("█", playerBlocks) + strings.Repeat(" ", barLength-playerBlocks)
	out += "\nYour HP:  [" + playerBar + "]"

	// Enemy HP bar
	enemyBlocks := int(float64(enemyHP) / float64(enemyMaxHP) * barLength)
	if enemyBlocks < 0 {
		enemyBlocks = 0
	}
	if enemyBlocks > barLength {
		enemyBlocks = barLength
	}
	enemyBar := strings.Repeat("█", enemyBlocks) + strings.Repeat(" ", barLength-enemyBlocks)
	out += "\nEnemy HP: [" + enemyBar + "]"

	return out
}
