package sim

import (
	"fmt"
	"gociv/pkg/config"
)

const CHARACTER_SPEED = 100

func (sim *Sim) InitCharacters() {
	sim.MakeCharacter("Henry", TilePosition{
		X: 25,
		Y: 25,
	})
	// sim.MakeCharacter("Emma", TilePosition{
	// 	X: 11,
	// 	Y: 13,
	// })
	// sim.MakeCharacter("Lise", TilePosition{
	// 	X: 11,
	// 	Y: 14,
	// })
	// sim.MakeCharacter("Ousmane", TilePosition{
	// 	X: 12,
	// 	Y: 10,
	// })
	// sim.MakeCharacter("Molly", TilePosition{
	// 	X: 12,
	// 	Y: 12,
	// })
	// sim.MakeCharacter("Robert", TilePosition{
	// 	X: 20,
	// 	Y: 14,
	// })
	// sim.MakeCharacter("Didier", TilePosition{
	// 	X: 20,
	// 	Y: 10,
	// })
	// sim.MakeCharacter("Morgane", TilePosition{
	// 	X: 20,
	// 	Y: 12,
	// })
}

func (sim *Sim) MakeCharacter(name string, pos TilePosition) {
	character := Character{
		ID:           int8(len(sim.Characters)),
		Name:         name,
		TilePosition: pos,
		WorldPosition: WorldPosition{
			X: float32(pos.X*config.TileSize + config.TileSize/2),
			Y: float32(pos.Y*config.TileSize + config.TileSize/2),
		},
		Needs: Needs{
			Food:  100,
			Water: 0,
			Sleep: 0,
		},
	}
	sim.Characters = append(sim.Characters, character)
}

func (sim *Sim) UpdateCharacters() {
	if sim.Time%config.CharacterNeedsUpdateInterval == 0 {
		for i := range sim.Characters {
			sim.Characters[i].UpdateNeeds()
		}
	}
	if sim.Time%config.CharacterObjectiveUpdateInterval == 0 {
		for i := range sim.Characters {
			sim.UpdateObjectives(&sim.Characters[i])
		}
	}
	if sim.Time%config.CharacterTaskUpdateInterval == 0 {
		for i := range sim.Characters {
			if sim.Characters[i].CurrentTask == nil {
				sim.SetCurrentTask(&sim.Characters[i])
			}
			sim.WorkOnCurrentTask(&sim.Characters[i])
		}
	}
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food++
	character.Needs.Water++
	character.Needs.Sleep++
}

// GetInventoryItems returns all items of a specific type and variant in the character's inventory
// if variant is -1, all variants are returned
func (sim *Sim) GetInventoryItems(character *Character, itemType ItemType, variant int16) []*Item {
	var items []*Item
	for _, itemID := range character.Inventory {
		item := sim.GetItemPtr(itemID)
		if item != nil && item.Type == itemType && (item.Variant == variant || variant == -1) {
			items = append(items, item)
		}
	}
	return items
}

func (sim *Sim) PickUp(character *Character) {
	task := character.CurrentTask
	item := task.TargetItem
	tile := sim.GetTileAt(item.Location.TilePosition)
	if item.Location.LocationType != LocTile || !IsAdjacent(character.TilePosition.X, character.TilePosition.Y, item.Location.TilePosition.X, item.Location.TilePosition.Y) {
		fmt.Printf("WARNING: Item %v to PICKUP is not on a tile or not adjacent\n", item)
		sim.CancelTask(character)
		return
	}
	fmt.Printf("Picking up %v\n", item)
	character.Inventory = append(character.Inventory, item.ID)
	item.Location.CharacterID = character.ID
	tile.RemoveItem(item.ID)
	task.Progress = 100
}
