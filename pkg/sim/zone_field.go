package sim

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/data"
	"math"
)

func (f Field) GetCentroid() TilePosition {
	return f.Centroid
}

func (f Field) GetTiles() []TilePosition {
	return f.Tiles
}

func (sim *Sim) CreateField(tiles []TilePosition, seedVariant int) *Field {
	newField := Field{
		Centroid:    GetZoneCentroid(tiles),
		Tiles:       tiles,
		SeedVariant: seedVariant,
		TileStatus:  make([]FieldTileStatus, len(tiles)),
	}
	sim.Fields = append(sim.Fields, newField)
	for _, tile := range tiles {
		sim.Tiles[sim.GetTileIDFromPosition(tile)].ZoneType = ZoneTypeField
		sim.Tiles[sim.GetTileIDFromPosition(tile)].ZoneIndex = int8(len(sim.Fields) - 1)
	}
	return &sim.Fields[len(sim.Fields)-1]
}

func (sim *Sim) UpdateFields() {
	for i := range sim.Fields {
		sim.UpdateField(&sim.Fields[i])
	}
}

func (sim *Sim) UpdateField(field *Field) {
	for i, tile := range field.TileStatus {
		if tile.Seeded {
			field.TileStatus[i].GrowthStage += config.FieldGrowthRate
			if field.TileStatus[i].GrowthStage >= 100 {
				field.TileStatus[i].Seeded = false
				field.TileStatus[i].GrowthStage = 0
				foodItem, _ := data.GetItemDefinition(int(ItemTypeFood), field.SeedVariant)
				if foodItem.ItemType != int(ItemTypeFood) {
					fmt.Printf("Food item type is not Food: %v\n", foodItem.ItemType)
					continue
				}
				sim.AddItem(Item{
					Type:       ItemTypeFood,
					Variant:    foodItem.Variant,
					Efficiency: foodItem.Efficiency,
				}, ItemLocation{LocationType: LocTile, TilePosition: field.Tiles[i]})
			}
		}
	}
}

func (sim *Sim) GetClosestField(tilePosition TilePosition) *Field {
	var closestField *Field
	var minDistance = -1
	for i := range sim.Fields {
		centroid := sim.Fields[i].GetCentroid()
		distance := math.Sqrt(float64((tilePosition.X-centroid.X)*(tilePosition.X-centroid.X) + (tilePosition.Y-centroid.Y)*(tilePosition.Y-centroid.Y)))
		if minDistance == -1 || distance < float64(minDistance) {
			minDistance = int(distance)
			closestField = &sim.Fields[i]
		}
	}
	return closestField
}

func (field *Field) GetFreeTiles() []TilePosition {
	var freeTiles []TilePosition
	for i, tile := range field.Tiles {
		if !field.TileStatus[i].Seeded {
			freeTiles = append(freeTiles, tile)
		}
	}
	return freeTiles
}

func (field *Field) GetGrowingTiles() []TilePosition {
	var growingTiles []TilePosition
	for i, tile := range field.Tiles {
		if field.TileStatus[i].Seeded && field.TileStatus[i].GrowthStage < 100 {
			growingTiles = append(growingTiles, tile)
		}
	}
	return growingTiles
}

func (sim *Sim) GetGrowingTilesCount() int {
	count := 0
	for _, field := range sim.Fields {
		count += len(field.GetGrowingTiles())
	}
	return count
}

func (sim *Sim) GetSuitableFieldTiles(character *Character) []TilePosition {
	var suitableTiles []TilePosition
	closestDirt := sim.ScanForTile(character.TilePosition, -1, TileTypeDirt)
	if closestDirt != nil {
		suitableTiles = append(suitableTiles, *closestDirt)
		// BFS: visit all dirt tiles by order of distance
		visited := make(map[string]bool)
		visited[getPositionKey(*closestDirt)] = true
		distance := 0
		levelStart := 0
		levelEnd := len(suitableTiles)

		for levelStart < len(suitableTiles) && len(suitableTiles) <= config.FieldDefaultSize {
			// Process all tiles at current distance level
			for i := levelStart; i < levelEnd; i++ {
				current := suitableTiles[i]

				// Check all neighbors
				for _, dir := range EightDirections {
					newX, newY := current.X+dir[0], current.Y+dir[1]

					// Check bounds
					if newX < 0 || newX >= config.RegionSize || newY < 0 || newY >= config.RegionSize {
						continue
					}

					// Check if already visited
					key := getPositionKey(TilePosition{X: newX, Y: newY})
					if visited[key] {
						continue
					}

					// Check if tile is dirt
					tileIndex := newY*config.RegionSize + newX
					if sim.Tiles[tileIndex].Type != TileTypeDirt {
						// Mark as visited but don't explore further
						visited[key] = true
						continue
					}

					// Mark as visited and add to suitableTiles
					visited[key] = true
					neighborPos := TilePosition{X: newX, Y: newY}
					suitableTiles = append(suitableTiles, neighborPos)
				}
			}

			// Move to next distance level
			distance++
			levelStart = levelEnd
			levelEnd = len(suitableTiles)
		}
	}
	return suitableTiles
}
