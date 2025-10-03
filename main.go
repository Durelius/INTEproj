package main

import (
	"INTE/projekt/character"
	"log"
)

func main() {
	char, err := character.New("josh")
	if err != nil {
		log.Fatal(err)
	}
	id := char.GetID()
	log.Println(id)

}
