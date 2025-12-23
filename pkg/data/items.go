package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// ItemDefinition represents an item configuration loaded from JSON
type ItemDefinition struct {
	ItemType   int    `json:"itemType"`
	Variant    int    `json:"variant"`
	Name       string `json:"name"`
	Efficiency int    `json:"efficiency"` // e.g. nutrition value for food
	StackSize  int    `json:"stackSize"`
}

// ItemDataFile represents the structure of the JSON file
type ItemDataFile struct {
	Items []ItemDefinition `json:"items"`
}

// ItemDefinitionsMap maps (ItemType, Variant) -> ItemDefinition
var ItemDefinitionsMap map[int]map[int]ItemDefinition

// LoadItemDefinitions loads item definitions from the JSON file
func LoadItemDefinitions() error {
	file, err := os.Open("pkg/data/items.json")
	if err != nil {
		return fmt.Errorf("failed to open items.json: %w", err)
	}
	defer file.Close()

	var data ItemDataFile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("failed to decode items.json: %w", err)
	}

	// Initialize the map
	ItemDefinitionsMap = make(map[int]map[int]ItemDefinition)

	// Populate the map
	for _, item := range data.Items {
		if ItemDefinitionsMap[item.ItemType] == nil {
			ItemDefinitionsMap[item.ItemType] = make(map[int]ItemDefinition)
		}
		ItemDefinitionsMap[item.ItemType][item.Variant] = item
	}

	fmt.Printf("Loaded %d item definitions\n", len(data.Items))
	return nil
}

// GetItemDefinition retrieves an item definition by type and variant
func GetItemDefinition(itemType int, variant int) (*ItemDefinition, bool) {
	if ItemDefinitionsMap == nil {
		return nil, false
	}
	if variantMap, ok := ItemDefinitionsMap[itemType]; ok {
		if def, ok := variantMap[variant]; ok {
			return &def, true
		}
	}
	return nil, false
}


