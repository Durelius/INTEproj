package enemy

import (
	"INTE/projekt/character"
)

type BaseEnemy struct {
	character.Character
	class Class
}

type Enemy interface {
	GetClass() Class
	character.Character
}

type Class string

const (
	CLASS_SKELETON Class = "SKELETON"
	CLASS_GOBLIN   Class = "GOBLIN"
	CLASS_SPIDER   Class = "SPIDER"
)

func New(class Class, name string) (Enemy, error) {
	char, err := character.New(name)
	if err != nil {
		return nil, err
	}
	return &BaseEnemy{class: class, Character: char}, nil
}

func (e *BaseEnemy) GetClass() Class {
	return e.class
}
func (e *BaseEnemy) Fight(rec character.Fightable) {

}
