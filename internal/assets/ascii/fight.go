package ascii

import (
	"fmt"
	"strings"

	"github.com/Durelius/INTEproj/internal/battle"
	"github.com/Durelius/INTEproj/internal/enemy"
	"github.com/Durelius/INTEproj/internal/player"
)

// Fight returns an ASCII battle screen with proportional HP bars
func Fight(b *battle.Battle) string {
	enemy := b.Enemy
	player := b.Player

	enemyHealth := generateEnemyHealthBar(enemy)
	playerHealth := generatePlayerHealthBar(player)
	attackMessage := generateAttackString(player, enemy, b.PlayerTurn())

	out := fmt.Sprintf(
		`	
%s
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
%s

%s
Press "a" to progress fight.
`, enemyHealth, playerHealth, attackMessage)

	return out
}

func generateAttackString(p *player.Player, e enemy.Enemy, playerTurn bool) string {
	bothAreMaxHealth := p.GetCurrentHealth() == p.GetMaxHealth() && e.GetCurrentHealth() == e.GetMaxHealth()

	if bothAreMaxHealth {
		return	""
	}

	if !playerTurn {
		return fmt.Sprintf("You attack %s, dealing %d damage.", e.GetEnemyType(), p.GetDamage())
	} else {
		return fmt.Sprintf("%s attacks you for %d damage!", e.GetEnemyType(), int((float32(e.GetDamage()) * (1 - p.GetDamageReduction()))))
	}
}

func generateEnemyHealthBar(e enemy.Enemy) string {

	const barLength = 10
	eHP := e.GetCurrentHealth()
	eMaxHP := e.GetMaxHealth()

	enemyBlocks := int(float64(eHP) / float64(eMaxHP) * barLength)
	if enemyBlocks < 0 {
		enemyBlocks = 0
	}
	if enemyBlocks > barLength {
		enemyBlocks = barLength
	}
	bar := strings.Repeat("█", enemyBlocks) + strings.Repeat(" ", barLength-enemyBlocks)


	return fmt.Sprintf("%s HP: %d/%d [ %s ]", e.GetEnemyType(), eHP, eMaxHP, bar)
}


func generatePlayerHealthBar(p *player.Player) string {
	
	const barLength = 10

	pHP := p.GetCurrentHealth()
	pMaxHP := p.GetMaxHealth()

	playerBlocks := int(float64(pHP) / float64(pMaxHP) * barLength)
	if playerBlocks < 0 {
		playerBlocks = 0
	}
	if playerBlocks > barLength {
		playerBlocks = barLength
	}
	bar  := strings.Repeat("█", playerBlocks) + strings.Repeat(" ", barLength-playerBlocks)
	
 	return fmt.Sprintf("%s HP: %d/%d [ %s ]", p.GetName(), pHP, pMaxHP, bar)
}