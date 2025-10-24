package sim

import (
	"encoding/gob"
	"fmt"
	"os"
)

const (
	REGION_SIZE = 50
	TILE_SIZE   = 30
)

func InitRegion() []Tile {
	region := make([]Tile, REGION_SIZE*REGION_SIZE)

	filename := "tiles.gob"
	file, err := os.Open(filename)
	if err != nil {
		for i := range region {
			region[i].Position = TilePosition{
				X: i % REGION_SIZE,
				Y: i / REGION_SIZE,
			}
		}
		return region
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&region)
	if err != nil {
		return region
	}
	fmt.Println("Loaded region from file")

	return region
}

func (s *Sim) GetTileAt(position TilePosition) *Tile {
	return &s.Tiles[position.Y*REGION_SIZE+position.X]
}
