package data

import (
	"fmt"
)

// LoadAllData loads all game data files
func LoadAllData() error {
	if err := LoadPlantDefinitions(); err != nil {
		return fmt.Errorf("failed to load plant definitions: %w", err)
	}
	if err := LoadStructureDefinitions(); err != nil {
		return fmt.Errorf("failed to load structure definitions: %w", err)
	}
	return nil
}
