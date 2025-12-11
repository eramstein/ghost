package sim

import "fmt"

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
	fmt.Printf("Initializing items\n")
	location := ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 1, Y: 1}}
	for i := 0; i < 8; i++ {
		sim.AddItem(Item{Type: ItemTypeFood, Efficiency: 30}, location)
	}
}
