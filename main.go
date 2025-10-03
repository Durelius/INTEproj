package main

import (
	"INTE/projekt/player"
	"log"
)

func main() {
	player, err := player.New(player.CLASS_MAGE, "josh")
	if err != nil {
		log.Fatal(err)
	}

	id := player.GetID()
	log.Println(id)

}
