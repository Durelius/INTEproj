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
	character.Fightable
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
func (p *BaseEnemy) GetDamage() int {
	return 100
}
func (p *BaseEnemy) IsFightable() (fightable character.Fightable, ok bool) {
	return p, true
}
func (e *BaseEnemy) Attack(rec character.Fightable) error {
	return nil
}
