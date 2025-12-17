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
	location := ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 7, Y: 7}}
	for i := 0; i < 8; i++ {
		sim.AddItem(Item{Type: ItemTypeFood, Efficiency: 30}, location)
	}
	location = ItemLocation{LocationType: LocTile, TilePosition: TilePosition{X: 30, Y: 5}}
	for i := 0; i < 8; i++ {
		sim.AddItem(Item{Type: ItemTypeFood, Efficiency: 30}, location)
	}
}
