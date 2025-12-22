package sim

import (
	"gociv/pkg/config"
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
				sim.AddItem(Item{
					Type:    ItemTypeFood,
					Variant: field.SeedVariant,
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
		for _, dir := range EightDirections {
			adjX := closestDirt.X + dir[0]
			adjY := closestDirt.Y + dir[1]
			tile := sim.GetTileAt(TilePosition{X: adjX, Y: adjY})
			if tile.Type == TileTypeDirt {
				suitableTiles = append(suitableTiles, TilePosition{X: adjX, Y: adjY})
			}
		}
	}
	return suitableTiles
}
