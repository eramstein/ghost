package sim

import (
	"encoding/gob"
	"fmt"
	"gociv/pkg/config"
	"math/rand"
	"os"
)

func InitRegion() []Tile {
	region := make([]Tile, config.RegionSize*config.RegionSize)

	filename := "tiles.gob"
	file, err := os.Open(filename)

	// if no saved region data, initialize the region with empty tiles
	if err != nil {
		for i := range region {
			region[i].Position = TilePosition{
				X: i % config.RegionSize,
				Y: i / config.RegionSize,
			}
			region[i].Structure = -1
			region[i].Plant = -1
		}
		return region
	}
	defer file.Close()

	// if saved region data, load it
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&region)
	if err != nil {
		return region
	}
	fmt.Println("Loaded region from file")

	return region
}

func (s *Sim) GetTileAt(position TilePosition) *Tile {
	return &s.Tiles[s.GetTileIDFromPosition(position)]
}

func (s *Sim) GetTileIDFromPosition(position TilePosition) int {
	return position.Y*config.RegionSize + position.X
}

func (s *Sim) GetRandomEmptyTile() *Tile {
	var emptyTiles []*Tile
	for i := range s.Tiles {
		if s.Tiles[i].Type != TileTypeWall {
			emptyTiles = append(emptyTiles, &s.Tiles[i])
		}
	}
	if len(emptyTiles) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(emptyTiles))
	return emptyTiles[randomIndex]
}
