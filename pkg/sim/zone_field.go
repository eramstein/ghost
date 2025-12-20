package sim

import "math"

func (f *Field) GetCentroid() TilePosition {
	return f.Centroid
}

func (f *Field) GetTiles() []TilePosition {
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
	return &sim.Fields[len(sim.Fields)-1]
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
