package sim

import "fmt"

func (sim *Sim) GetNextMakingFoodTask(character *Character, objective *Objective) (task *Task) {
	// for now we can only make food by planting seeds
	// it requires both a growable tile and a seed
	var newTask *Task

	// does the character have a seed?
	if len(sim.GetInventoryItems(character, ItemTypeSeed, -1)) > 0 {
		seed := sim.GetInventoryItems(character, ItemTypeSeed, -1)[0]
		// if yes, go plant it
		// get closest field, create one if needed
		field := sim.GetClosestField(character.TilePosition)
		if field == nil {
			suitableTiles := sim.GetSuitableFieldTiles(character)
			if len(suitableTiles) == 0 {
				ObjectiveFailed(character, objective)
				fmt.Printf("No suitable field tiles found for %v\n", character.Name)
				return nil
			}
			field = sim.CreateField(suitableTiles, 0)
			fmt.Printf("Created field with %v tiles\n", len(field.Tiles))
		}

		// plant seeds on the closest free tile
		freeTiles := field.GetFreeTiles()
		if len(freeTiles) > 0 {
			if freeTiles[0].IsSameAs(character.TilePosition) {
				// if the closest free tile is 	the character's current tile, create task to plant seeds
				newTask = &Task{
					Objective:      objective,
					Type:           PlantSeed,
					TargetTile:     &freeTiles[0],
					MaterialSource: seed,
				}
			} else {
				// if yes, go to the closest free tile
				newTask = &Task{
					Objective:  objective,
					Type:       Move,
					TargetTile: &freeTiles[0],
				}
			}
		} else {
			ObjectiveFailed(character, objective)
			fmt.Printf("No free field tiles found for %v\n", character.Name)
			return nil
		}
	} else {
		// if no, is there a seed on the tile?
		seedOnTile := sim.FindItemInTile(character.ID, character.TilePosition, ItemTypeSeed, -1, true)
		if seedOnTile != nil {
			// if yes, pick it up
			newTask = &Task{
				Objective:  objective,
				Type:       PickUp,
				TargetItem: seedOnTile,
			}
		} else {
			// is there a seed somewhere else?
			closestSeed := sim.ScanForItem(character.ID, character.TilePosition, -1, ItemTypeSeed, -1, true)
			if closestSeed != nil {
				// if yes, go to it
				newTask = &Task{
					Objective:  objective,
					Type:       Move,
					TargetTile: &closestSeed.Location.TilePosition,
				}
			} else {
				// if no, stuck objective (TODO: get a way to provide seeds)
				ObjectiveFailed(character, objective)
				fmt.Printf("No seed found for %v\n", character.Name)
				return nil
			}
		}
	}

	return newTask
}

func (sim *Sim) PlantSeed(character *Character) {
	task := character.CurrentTask
	tile := sim.GetTileAt(*task.TargetTile)
	if tile.ZoneType != ZoneTypeField {
		fmt.Printf("Tile %v is not a field\n", tile.Position)
		return
	}
	field := sim.Fields[tile.ZoneIndex]
	materialSource := task.MaterialSource
	if materialSource == nil {
		fmt.Printf("No material source found for %v\n", character.Name)
		return
	}
	if materialSource.Type != ItemTypeSeed {
		fmt.Printf("Material source %v is not a seed\n", materialSource.Type)
		return
	}
	tileFieldIndex := GetZoneTileIndex(field, tile.Position)
	if tileFieldIndex == -1 {
		fmt.Printf("Tile %v is not in field %v\n", tile.Position, field.GetTiles())
		return
	}
	task.Progress += 20
	fmt.Println("Planting seed on", character.Name, tile)
	if task.Progress >= 100 {
		field.TileStatus[tileFieldIndex].Seeded = true
		field.TileStatus[tileFieldIndex].GrowthStage = 0
		field.TileStatus[tileFieldIndex].SeedVariant = materialSource.Variant
		sim.DecreaseItemStackCount(materialSource.ID)
		fmt.Printf("Planted seed on %v with variant %v\n", tile.Position, materialSource.Variant)
	}
}
