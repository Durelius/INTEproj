package item

type Weapon struct {
	Item
	damage int
}

func (w *Weapon) GetDamage() int {
	return w.damage
}
func (w *Weapon) GetWearPosition() WearPosition {
	return WEAR_POSITION_WEAPON
}
