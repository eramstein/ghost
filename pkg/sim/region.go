package sim

import (
	"encoding/gob"
	"fmt"
	"gociv/pkg/config"
	"math/rand"
	"os"
)

// RegionData contains all region-related data (tiles, plants, structures)
type RegionData struct {
	Tiles            []Tile
	PlantManager     *PlantManager
	StructureManager *StructureManager
}

// RegionInitResult contains the loaded region data
type RegionInitResult struct {
	Tiles            []Tile
	PlantManager     *PlantManager
	StructureManager *StructureManager
}

// LoadRegion loads the complete region data (tiles, plants, structures) from a file using gob decoding
func LoadRegion(filename string) (*RegionData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var regionData RegionData
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&regionData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode region data: %w", err)
	}

	return &regionData, nil
}

func InitRegion() RegionInitResult {
	result := RegionInitResult{
		Tiles:            make([]Tile, config.RegionSize*config.RegionSize),
		PlantManager:     NewPlantManager(),
		StructureManager: NewStructureManager(),
	}

	filename := "tiles.gob"
	regionData, err := LoadRegion(filename)

	// if no saved region data, initialize the region with empty tiles
	if err != nil {
		for i := range result.Tiles {
			result.Tiles[i].Position = TilePosition{
				X: i % config.RegionSize,
				Y: i / config.RegionSize,
			}
			result.Tiles[i].Structure = -1
			result.Tiles[i].Plant = -1
		}
		return result
	}

	// if saved region data, load it
	result.Tiles = regionData.Tiles
	if regionData.PlantManager != nil {
		result.PlantManager = regionData.PlantManager
	}
	if regionData.StructureManager != nil {
		result.StructureManager = regionData.StructureManager
	}
	fmt.Println("Loaded region from file")

	return result
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
