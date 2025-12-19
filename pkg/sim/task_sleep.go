package sim

import (
	"fmt"
	"gociv/pkg/config"
)

func (sim *Sim) GetNextSleepingTask(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// If the character is in their bed, add a task to sleep
	if bed := sim.FindStructureInTile(character.ID, character.TilePosition, Bed, -1, false); bed != nil {
		newTask = &Task{
			Objective: objective,
			Type:      Sleep,
		}
	} else {
		// Else, go to their bed if they have one or claim one if they don't have one
		characterBeds := sim.StructureManager.GetStructuresByOwnerAndType(character.ID, Bed)
		if len(characterBeds) > 0 {
			newTask = &Task{
				Objective:  objective,
				Type:       Move,
				TargetTile: &characterBeds[0].Position,
			}
		} else {
			// Claim the closest bed
			closestBed := sim.ScanForStructure(character.ID, character.TilePosition, config.RegionSize, Bed, -1, true)
			if closestBed != nil {
				closestBed.Owner = character.ID
				newTask = &Task{
					Objective:  objective,
					Type:       Move,
					TargetTile: &closestBed.Position,
				}
			} else {
				// If no bed found, add an objective to build one
				fmt.Printf("No bed found for %v, adding objective to build one\n", character.Name)
			}
		}
	}
	return newTask
}

func (sim *Sim) Sleep(character *Character) {
	task := character.CurrentTask
	fmt.Println("Sleeping", character.Name)
	character.Needs.Sleep -= 5
	if character.Needs.Sleep <= 0 {
		character.Needs.Sleep = 0
		task.Progress = 100
	}
}
