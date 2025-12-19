package sim

import (
	"fmt"
	"gociv/pkg/config"
)

func (sim *Sim) Eat(character *Character) {
	task := character.CurrentTask
	item := task.TargetItem
	task.Progress += 10
	fmt.Println("Eating", character.Name, item.Type, item.Efficiency)
	if task.Progress >= 100 {
		character.Needs.Food -= item.Efficiency
		if character.Needs.Food < 0 {
			character.Needs.Food = 0
		}
		sim.RemoveItem(item.ID)
	}
}

// Set next task required to achieve an eat objective
func (sim *Sim) GetNextEatingTask(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// Check if the character has the item in their inventory
	itemInInventory := sim.FindInInventory(character, ItemTypeFood, -1)
	// If the character has the item in their inventory, add a task to eat it
	if itemInInventory != nil {
		newTask = &Task{
			Objective:  objective,
			Type:       Eat,
			TargetItem: itemInInventory,
		}
		// If the character is on a tile with a food item, add a task to eat it
	} else if itemOnTile := sim.FindItemInTile(character.ID, character.TilePosition, ItemTypeFood, -1, true); itemOnTile != nil {
		// claim item
		itemOnTile.OwnedBy = character.ID
		// eat it
		newTask = &Task{
			Objective:  objective,
			Type:       Eat,
			TargetItem: itemOnTile,
		}
	} else {
		// If no food on tile, find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.ID, character.TilePosition, config.RegionSize/2-1, ItemTypeFood, -1, true)
		if closestItem != nil {
			// claim item
			closestItem.OwnedBy = character.ID
			// go to it
			newTask = &Task{
				Objective:  objective,
				Type:       Move,
				TargetTile: &closestItem.Location.TilePosition,
			}
		} else {
			fmt.Printf("No food found for %v\n", character.Name)
		}
	}
	return newTask
}
