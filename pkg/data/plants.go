package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// PlantDefinition represents a plant configuration loaded from JSON
type PlantDefinition struct {
	PlantType  int           `json:"plantType"`
	Variant    int16         `json:"variant"`
	Name       string        `json:"name"`
	GrowthRate uint8         `json:"growthRate"`
	Produces   ProductionDef `json:"produces"`
}

// ProductionDef represents what a plant produces
type ProductionDef struct {
	Type           int   `json:"type"`
	Variant        int16 `json:"variant"`
	ProductionRate uint8 `json:"productionRate"`
}

// PlantDataFile represents the structure of the JSON file
type PlantDataFile struct {
	Plants []PlantDefinition `json:"plants"`
}

// PlantDefinitionsMap maps (PlantType, Variant) -> PlantDefinition
var PlantDefinitionsMap map[int]map[int16]PlantDefinition

// LoadPlantDefinitions loads plant definitions from the JSON file
func LoadPlantDefinitions() error {
	file, err := os.Open("pkg/data/plants.json")
	if err != nil {
		return fmt.Errorf("failed to open plants.json: %w", err)
	}
	defer file.Close()

	var data PlantDataFile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("failed to decode plants.json: %w", err)
	}

	// Initialize the map
	PlantDefinitionsMap = make(map[int]map[int16]PlantDefinition)

	// Populate the map
	for _, plant := range data.Plants {
		if PlantDefinitionsMap[plant.PlantType] == nil {
			PlantDefinitionsMap[plant.PlantType] = make(map[int16]PlantDefinition)
		}
		PlantDefinitionsMap[plant.PlantType][plant.Variant] = plant
	}

	fmt.Printf("Loaded %d plant definitions\n", len(data.Plants))
	return nil
}

// GetPlantDefinition retrieves a plant definition by type and variant
func GetPlantDefinition(plantType int, variant int16) (*PlantDefinition, bool) {
	if PlantDefinitionsMap == nil {
		return nil, false
	}
	if variantMap, ok := PlantDefinitionsMap[plantType]; ok {
		if def, ok := variantMap[variant]; ok {
			return &def, true
		}
	}
	return nil, false
}
