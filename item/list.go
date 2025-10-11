package item

var (
	NOTHING       = &BaseItem{weight: 0, name: ""}
	IRON_SWORD    = Weapon{damage: 20, Item: &BaseItem{weight: 50, name: "iron sword"}}
	STICK         = Weapon{damage: 10, Item: &BaseItem{weight: 20, name: "wooden stick"}}
	IRON_HELMET   = Head{defense: 35, Item: &BaseItem{weight: 30, name: "iron helmet"}}
	CHAIN_MAIL    = UpperBody{defense: 50, Item: &BaseItem{weight: 50, name: "chain mail"}}
	LEATHER_TUNIC = UpperBody{defense: 15, Item: &BaseItem{weight: 10, name: "leather tunic"}}
	ADIDAS_PANTS  = LowerBody{defense: 13, Item: &BaseItem{weight: 10, name: "Three stripes"}}
)
