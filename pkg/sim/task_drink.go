package sim

import (
	"fmt"
	"gociv/pkg/config"
)

func (sim *Sim) Drink(character *Character) {
	task := character.CurrentTask
	position := task.TargetTile
	tile := sim.GetTileAt(*position)
	if tile.Type != TileTypeWater && sim.FindStructureInTile(character.ID, *position, Well, -1, true) == nil {
		return
	}
	task.Progress += 50
	fmt.Println("Drinking", character.Name)
	if task.Progress >= 100 {
		character.Needs.Water = 0
		task.Progress = 100
	}
}

func (sim *Sim) GetNextDrinkingTask(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// Go to the closest water tile if needed, then drink
	var closestWater *TilePosition
	closestWell := sim.ScanForStructure(character.ID, character.TilePosition, config.RegionSize/2-1, Well, -1, true)
	if closestWell != nil {
		closestWater = &closestWell.Position
	} else {
		closestWater = sim.ScanForTile(character.TilePosition, config.RegionSize/2-1, TileTypeWater)
	}
	fmt.Printf("closestWater: %v\n", closestWater)
	if closestWater == nil {
		return
	}
	if IsAdjacent(character.TilePosition.X, character.TilePosition.Y, closestWater.X, closestWater.Y) {
		newTask = &Task{
			Objective:  objective,
			Type:       Drink,
			TargetTile: closestWater,
		}
	} else {
		// stop one tile before the water tile
		// problem: if closestWater is not accessible, there will be no path found and no task added
		path := sim.FindPath(character.TilePosition, *closestWater, 1)
		if len(path) > 0 {
			newTask = &Task{
				Objective:  objective,
				Type:       Move,
				TargetTile: &(path[len(path)-1]),
			}
		} else {
			fmt.Printf("No drinking path found for %v %v\n", character.Name, closestWater)
		}
	}
	return newTask
}
