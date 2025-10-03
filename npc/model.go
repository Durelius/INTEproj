package npc

import (
	"INTE/projekt/character"
)

type BaseNPC struct {
	character.Character
}
type NPC interface {
	character.Character
}

func New(name string) (NPC, error) {
	char, err := character.New(name)
	if err != nil {
		return nil, err
	}
	return &BaseNPC{Character: char}, nil
}
