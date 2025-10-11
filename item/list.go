package item

var (
	NOTHING       = &BaseItem{weight: 0, name: "", itemType: EMPTY}
	IRON_SWORD    = Weapon{damage: 20, Item: &BaseItem{weight: 50, name: "iron sword", itemType: WEAPON}}
	STICK         = Weapon{damage: 10, Item: &BaseItem{weight: 20, name: "wooden stick", itemType: WEAPON}}
	IRON_HELMET   = Wearable{defense: 35, Item: &BaseItem{weight: 30, name: "iron helmet", itemType: WEARABLE}}
	CHAIN_MAIL    = Wearable{defense: 50, Item: &BaseItem{weight: 50, name: "chain mail", itemType: WEARABLE}}
	LEATHER_TUNIC = Wearable{defense: 15, Item: &BaseItem{weight: 10, name: "leather tunic", itemType: WEARABLE}}
)
