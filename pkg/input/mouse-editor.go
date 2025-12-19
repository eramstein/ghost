package input

import (
	"fmt"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Manager) HandleMouseEditor(deltaTime float32) {
	// Handle mouse clicks (only in EditMode)
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.ScreenToWorld(mouseX, mouseY)
		tilePos := m.WorldToTile(worldX, worldY)
		tile := m.sim.GetTileAt(tilePos)

		switch m.sim.UI.EditorMode {
		case sim.EditorModeTiles:
			fmt.Printf("Clicked tile position: (%d, %d)\n", tile.Position.X, tile.Position.Y)
			tile.UpdateType(m.sim.UI.EditorTileType)

		case sim.EditorModePlants:
			// Check if tile already has a plant
			if tile.Plant >= 0 {
				fmt.Printf("Tile already has a plant (ID: %d), removing it\n", tile.Plant)
				m.sim.RemovePlant(tile.Plant)
			}
			// Add new plant
			plantDef, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), m.sim.UI.EditorPlantVariant)
			if ok && plantDef != nil {
				plantID := m.sim.SpawnPlant(tilePos, m.sim.UI.EditorPlantVariant, m.sim.UI.EditorPlantType)
				fmt.Printf("Added plant (ID: %d, Type: %d, Variant: %d, Name: %s) at (%d, %d)\n",
					plantID, m.sim.UI.EditorPlantType, m.sim.UI.EditorPlantVariant, plantDef.Name, tilePos.X, tilePos.Y)
			} else {
				fmt.Printf("Invalid plant definition (Type: %d, Variant: %d)\n", m.sim.UI.EditorPlantType, m.sim.UI.EditorPlantVariant)
			}

		case sim.EditorModeStructures:
			// Check if tile already has a structure
			if tile.Structure >= 0 {
				fmt.Printf("Tile already has a structure (ID: %d), removing it\n", tile.Structure)
				m.sim.RemoveStructure(tile.Structure)
			}
			// Add new structure
			structureID := m.sim.SpawnStructure(tilePos, m.sim.UI.EditorStructureType)
			fmt.Printf("Added structure (ID: %d, Type: %d) at (%d, %d)\n",
				structureID, m.sim.UI.EditorStructureType, tilePos.X, tilePos.Y)
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.ScreenToWorld(mouseX, mouseY)
		tilePos := m.WorldToTile(worldX, worldY)
		tile := m.sim.GetTileAt(tilePos)

		switch m.sim.UI.EditorMode {
		case sim.EditorModeTiles:
			fmt.Printf("Right-clicked tile position: (%d, %d)\n", tile.Position.X, tile.Position.Y)
			tile.UpdateType(sim.TileTypeEmpty)

		case sim.EditorModePlants:
			// Remove plant if present
			if tile.Plant >= 0 {
				fmt.Printf("Removing plant (ID: %d) from tile (%d, %d)\n", tile.Plant, tilePos.X, tilePos.Y)
				m.sim.RemovePlant(tile.Plant)
			}

		case sim.EditorModeStructures:
			// Remove structure if present
			if tile.Structure >= 0 {
				fmt.Printf("Removing structure (ID: %d) from tile (%d, %d)\n", tile.Structure, tilePos.X, tilePos.Y)
				m.sim.RemoveStructure(tile.Structure)
			}
		}
	}

	// Handle mouse wheel
	wheelMove := rl.GetMouseWheelMove()
	if wheelMove != 0 {
		fmt.Printf("Mouse wheel: %f\n", wheelMove)
		// Add mouse wheel handling logic here
	}
}
