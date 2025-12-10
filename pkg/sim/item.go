package sim

type ItemType int

const (
	ItemTypeFood ItemType = iota
	ItemTypeTool
	ItemTypeWeapon
)

type ItemLocationType uint8

const (
	LocNone ItemLocationType = iota
	LocTile
	LocCharacter
)

func (sim *Sim) InitItems() {
	location := ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}}
	for i := 0; i < 15; i++ {
		sim.AddItem(Item{Type: ItemTypeFood}, location)
	}
}
