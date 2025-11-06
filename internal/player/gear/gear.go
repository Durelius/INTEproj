package gear

import "github.com/Durelius/INTEproj/internal/item"

type Gear struct {
	Head      item.Item
	Upperbody item.Item
	Legs      item.Item
	Weapon    item.Item
	Feet      item.Item
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
