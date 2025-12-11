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

// SaveTiles saves only the Tiles data from a Sim to a file using gob encoding
func SaveTiles(tiles []sim.Tile, filename string) error {
	// Set all Items to nil before saving
	for i := range tiles {
		tiles[i].Items = nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(tiles)
	if err != nil {
		return fmt.Errorf("failed to encode tiles data: %w", err)
	}

	return nil
}

// LoadTiles loads only the Tiles data from a file using gob decoding
func LoadTiles(filename string) ([]sim.Tile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var tiles []sim.Tile
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&tiles)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tiles data: %w", err)
	}

	return tiles, nil
}
