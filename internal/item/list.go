package item

var AllItems = []Item{
	&Weapon{damage: 20, weight: 50, name: "Iron Sword"},
	&Weapon{damage: 10, weight: 20, name: "Wooden Stick"},
	&Wearable{defense: 35, weight: 30, name: "Iron Helmet", slot: WEAR_POSITION_HEAD},
	&Wearable{defense: 50, weight: 50, name: "Chain Mail", slot: WEAR_POSITION_UPPER_BODY},
	&Wearable{defense: 15, weight: 10, name: "Leather Tunic", slot: WEAR_POSITION_UPPER_BODY},
	&Wearable{defense: 13, weight: 10, name: "Three Stripes", slot: WEAR_POSITION_LOWER_BODY},

	&Weapon{damage: 15, weight: 40, name: "Common Sword", rarity: Common},
	&Weapon{damage: 25, weight: 45, name: "Rare Sword", rarity: Rare},
	&Weapon{damage: 40, weight: 50, name: "Epic Sword", rarity: Epic},
	&Weapon{damage: 60, weight: 55, name: "Legendary Sword", rarity: Legendary},

	&Weapon{damage: 10, weight: 30, name: "Common Bow", rarity: Common},
	&Weapon{damage: 20, weight: 35, name: "Rare Bow", rarity: Rare},
	&Weapon{damage: 35, weight: 38, name: "Epic Bow", rarity: Epic},
	&Weapon{damage: 50, weight: 40, name: "Legendary Bow", rarity: Legendary},
}
