package gear

import "github.com/Durelius/INTEproj/internal/item"

type Gear struct {
	Head      item.Item
	Upperbody item.Item
	Legs      item.Item
	Weapon    item.Item
	Feet      item.Item
}


func (g *Gear) Equip(it item.Item) {
	// Should also check the level requirements of items and if something is currently equipped in the slot.

	switch it.GetType() {
	case "WEAPON":	
		// do something
	case "WEARABLE":
		// do something
	}	
}

func (g *Gear) GetTotalWeight() int {
	total := 0
	for _, it := range []item.Item{g.Head, g.Upperbody, g.Legs, g.Weapon, g.Feet} {
		if it != nil {
			total += it.GetWeight()
		}
	}
	return total
}