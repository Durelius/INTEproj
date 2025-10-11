package main

import (
	"INTE/projekt/enemy"
	"INTE/projekt/player"
	"log"
)

func main() {
	player, err := player.New(player.CLASS_MAGE, "josh")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(player.GetHealth())

	enemy, err := enemy.New(enemy.CLASS_GOBLIN, "gobgosh")
	if err != nil {
		log.Fatal(err)
	}
	if fightable, ok := player.IsFightable(); ok {
		if enemyFightable, ok := enemy.IsFightable(); ok {
			fightable.Attack(enemyFightable)
		}
	}

}
