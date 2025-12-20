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
	// location := ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 7, Y: 7}}
	// for i := 0; i < 8; i++ {
	// 	sim.AddItem(Item{Type: ItemTypeFood, Efficiency: 30}, location)
	// }
}
