package sim

import "fmt"

type ItemType int

const (
	ItemTypeNone ItemType = iota
	ItemTypeFood
	ItemTypeTool
	ItemTypeWeapon
	ItemTypeSeed
)

type ItemLocationType uint8

const (
	LocNone ItemLocationType = iota
	LocTile
	LocCharacter
)

func (sim *Sim) InitItems() {
	fmt.Printf("Initializing items\n")
	location := ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 16, Y: 16}}
	sim.AddItem(Item{Type: ItemTypeSeed, Variant: 2, StackCount: 8}, location)
}
