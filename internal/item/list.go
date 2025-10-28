package item

// Carls
var (
	NOTHING       = &BaseItem{weight: 0, name: ""}
	IRON_SWORD    = Weapon{damage: 20, Item: &BaseItem{weight: 50, name: "iron sword"}}
	STICK         = Weapon{damage: 10, Item: &BaseItem{weight: 20, name: "wooden stick"}}
	IRON_HELMET   = Head{defense: 35, Item: &BaseItem{weight: 30, name: "iron helmet"}}
	CHAIN_MAIL    = UpperBody{defense: 50, Item: &BaseItem{weight: 50, name: "chain mail"}}
	LEATHER_TUNIC = UpperBody{defense: 15, Item: &BaseItem{weight: 10, name: "leather tunic"}}
	ADIDAS_PANTS  = LowerBody{defense: 13, Item: &BaseItem{weight: 10, name: "Three stripes"}}

	COMMON_SWORD    = Weapon{damage: 15, Item: &BaseItem{weight: 40, name: "Common Sword", rarity: Common}}
	RARE_SWORD      = Weapon{damage: 25, Item: &BaseItem{weight: 45, name: "Rare Sword", rarity: Rare}}
	EPIC_SWORD      = Weapon{damage: 40, Item: &BaseItem{weight: 50, name: "Epic Sword", rarity: Epic}}
	LEGENDARY_SWORD = Weapon{damage: 60, Item: &BaseItem{weight: 55, name: "Legendary Sword", rarity: Legendary}}

	COMMON_BOW    = Weapon{damage: 10, Item: &BaseItem{weight: 30, name: "Common Bow", rarity: Common}}
	RARE_BOW      = Weapon{damage: 20, Item: &BaseItem{weight: 35, name: "Rare Bow", rarity: Rare}}
	EPIC_BOW      = Weapon{damage: 35, Item: &BaseItem{weight: 38, name: "Epic Bow", rarity: Epic}}
	LEGENDARY_BOW = Weapon{damage: 50, Item: &BaseItem{weight: 40, name: "Legendary Bow", rarity: Legendary}}
)

// Depricated, används inte. Använd *_ITEMS listorna
var (
	ITEM_LIST_DROPPABLE = []Item{
		&IRON_SWORD, &STICK, &IRON_HELMET, &STICK, &IRON_HELMET, &CHAIN_MAIL, &LEATHER_TUNIC, &ADIDAS_PANTS,
	}
)
var (
	COMMON_ITEMS = []Item{&COMMON_SWORD, &COMMON_BOW}
)
var (
	RARE_ITEMS = []Item{&RARE_BOW, &RARE_BOW}
)
var (
	EPIC_ITEMS = []Item{&EPIC_BOW, &EPIC_SWORD}
)
var (
	LEGENDARY_ITEMS = []Item{&LEGENDARY_SWORD, &LEGENDARY_BOW}
)
