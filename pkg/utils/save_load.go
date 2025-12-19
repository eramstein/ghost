package utils

import (
	"encoding/gob"
	"fmt"
	"gociv/pkg/sim"
	"os"
)

// SaveSim saves the current Sim state to a quicksave file using gob encoding
func SaveSim(s *sim.Sim, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(s)
	if err != nil {
		return fmt.Errorf("failed to encode sim data: %w", err)
	}

	return nil
}

// LoadSim loads a Sim state from a quicksave file using gob decoding
func LoadSim(filename string) (*sim.Sim, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var sim sim.Sim
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&sim)
	if err != nil {
		return nil, fmt.Errorf("failed to decode sim data: %w", err)
	}

	return &sim, nil
}

// SaveRegion saves the complete region data (tiles, plants, structures) to a file using gob encoding
func SaveRegion(s *sim.Sim, filename string) error {
	// Set all Items to nil before saving
	tiles := make([]sim.Tile, len(s.Tiles))
	copy(tiles, s.Tiles)
	for i := range tiles {
		tiles[i].Items = nil
	}

	regionData := sim.RegionData{
		Tiles:            tiles,
		PlantManager:     s.PlantManager,
		StructureManager: s.StructureManager,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(regionData)
	if err != nil {
		return fmt.Errorf("failed to encode region data: %w", err)
	}

	return nil
}
