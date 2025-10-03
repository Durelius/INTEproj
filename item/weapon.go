package item

type Weapon struct {
	Item
	damage int
}

func (w *Weapon) GetDamage() int {
	return w.damage
}
