package data

import (
	"fmt"
)

// LoadAllData loads all game data files
func LoadAllData() error {
	if err := LoadPlantDefinitions(); err != nil {
		return fmt.Errorf("failed to load plant definitions: %w", err)
	}
	if err := LoadItemDefinitions(); err != nil {
		return fmt.Errorf("failed to load item definitions: %w", err)
	}
	return nil
}
