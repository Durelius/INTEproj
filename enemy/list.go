package enemy

var ENEMY_LIST = []Enemy{
	func() Enemy { e, _ := New(CLASS_SKELETON, "Bonecrusher"); return e }(),
	func() Enemy { e, _ := New(CLASS_SKELETON, "Rattler"); return e }(),
	func() Enemy { e, _ := New(CLASS_GOBLIN, "Goblin Grunt"); return e }(),
	func() Enemy { e, _ := New(CLASS_GOBLIN, "Sneaky Goblin"); return e }(),
	func() Enemy { e, _ := New(CLASS_SPIDER, "Web Spinner"); return e }(),
	func() Enemy { e, _ := New(CLASS_SPIDER, "Night Crawler"); return e }(),
}
