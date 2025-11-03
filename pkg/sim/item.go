package sim

type ItemType int

const (
	ItemTypeFood ItemType = iota
	ItemTypeTool
	ItemTypeWeapon
)

type ItemRef struct {
	Type  ItemType
	Index int
}

type ItemLocationType uint8

const (
	LocNone ItemLocationType = iota
	LocTile
	LocCharacter
)

func (sim *Sim) InitItems() {
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood}, ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}})
}
