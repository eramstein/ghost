package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// StructureDefinition represents a structure configuration loaded from JSON
type StructureDefinition struct {
	StructureType int    `json:"structureType"`
	Variant       int    `json:"variant"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

// StructureDataFile represents the structure of the JSON file
type StructureDataFile struct {
	Structures []StructureDefinition `json:"structures"`
}

// StructureDefinitionsMap maps (StructureType, Variant) -> StructureDefinition
var StructureDefinitionsMap map[int]map[int]StructureDefinition

// LoadStructureDefinitions loads structure definitions from the JSON file
func LoadStructureDefinitions() error {
	file, err := os.Open("pkg/data/structures.json")
	if err != nil {
		return fmt.Errorf("failed to open structures.json: %w", err)
	}
	defer file.Close()

	var data StructureDataFile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("failed to decode structures.json: %w", err)
	}

	// Initialize the map
	StructureDefinitionsMap = make(map[int]map[int]StructureDefinition)

	// Populate the map
	for _, structure := range data.Structures {
		if StructureDefinitionsMap[structure.StructureType] == nil {
			StructureDefinitionsMap[structure.StructureType] = make(map[int]StructureDefinition)
		}
		StructureDefinitionsMap[structure.StructureType][structure.Variant] = structure
	}

	fmt.Printf("Loaded %d structure definitions\n", len(data.Structures))
	return nil
}

// GetStructureDefinition retrieves a structure definition by type and variant
func GetStructureDefinition(structureType int, variant int) (*StructureDefinition, bool) {
	if StructureDefinitionsMap == nil {
		return nil, false
	}
	if variantMap, ok := StructureDefinitionsMap[structureType]; ok {
		if def, ok := variantMap[variant]; ok {
			return &def, true
		}
	}
	return nil, false
}
