package enemy

import (
	"fmt"

	"github.com/Durelius/INTEproj/internal/character"
)

type BaseEnemy struct {
	character.Character
	class Class
}

type Enemy interface {
	GetClass() Class
	character.Character
	character.Fightable
	GetType() string
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
func (e *BaseEnemy) GetDamage() int {
	return 100
}
func (e *BaseEnemy) IsFightable() (fightable character.Fightable, ok bool) {
	return e, true
}
func (e *BaseEnemy) Attack(rec character.Character) (int, error) {
	eFightable, ok := e.IsFightable()
	if !ok {
		return 0, fmt.Errorf("Attacker can't fight")
	}
	pFightable, ok := rec.IsFightable()
	if !ok {
		return 0, fmt.Errorf("Receiver can't fight")
	}

	return pFightable.ReceiveDamage(eFightable.GetDamage()), nil
}
func (e *BaseEnemy) ReceiveDamage(damage int) int {
	e.SetHealth(e.GetHealth() - damage)

	return e.GetHealth()
}
func (e *BaseEnemy) GetType() string {
	return "ENEMY"
}
