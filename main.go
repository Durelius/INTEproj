package main

import (
	"INTE/projekt/enemy"
	"INTE/projekt/player"
	"log"
)

func main() {
	player, err := player.New(player.CLASS_PALADIN, "josh")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(player.GetHealth())

	enemy, err := enemy.New(enemy.CLASS_GOBLIN, "gobgosh")
	if err != nil {
		log.Fatal(err)
	}
	newHP, err := player.Attack(enemy)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(newHP)

}
