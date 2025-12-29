package input

import (
	"fmt"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// handleKeyboard processes keyboard input events
func (m *Manager) HandleKeyboardEditor(deltaTime float32) {
	// Mode switching: T for Tiles, P for Plants, B for Structures
	if rl.IsKeyPressed(rl.KeyT) {
		m.sim.UI.EditorMode = sim.EditorModeTiles
		fmt.Println("Editor mode: Tiles")
	}
	if rl.IsKeyPressed(rl.KeyP) {
		m.sim.UI.EditorMode = sim.EditorModePlants
		fmt.Println("Editor mode: Plants")
	}
	if rl.IsKeyPressed(rl.KeyB) {
		m.sim.UI.EditorMode = sim.EditorModeStructures
		fmt.Println("Editor mode: Structures")
	}

	// Handle input based on current editor mode
	switch m.sim.UI.EditorMode {
	case sim.EditorModeTiles:
		// Map number keys to TileType values
		if rl.IsKeyPressed(rl.KeyOne) {
			m.sim.UI.EditorTileType = sim.TileTypeEmpty
			fmt.Println("Editor tile type set to: Empty")
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			m.sim.UI.EditorTileType = sim.TileTypeWall
			fmt.Println("Editor tile type set to: Wall")
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			m.sim.UI.EditorTileType = sim.TileTypeFloor
			fmt.Println("Editor tile type set to: Floor")
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			m.sim.UI.EditorTileType = sim.TileTypeDirt
			fmt.Println("Editor tile type set to: Dirt")
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			m.sim.UI.EditorTileType = sim.TileTypeWater
			fmt.Println("Editor tile type set to: Water")
		}

	case sim.EditorModePlants:
		// Number keys select plant type (0 = PlantTypeTree)
		if rl.IsKeyPressed(rl.KeyOne) {
			m.sim.UI.EditorPlantType = sim.PlantTypeTree
			m.sim.UI.EditorPlantVariant = 0
			fmt.Println("Editor plant type set to: Tree, Variant: 0")
		}
		// Use Q/E to cycle through variants for the current plant type
		if rl.IsKeyPressed(rl.KeyQ) {
			// Decrease variant
			if m.sim.UI.EditorPlantVariant > 0 {
				m.sim.UI.EditorPlantVariant--
			} else {
				// Try to find the max variant for this plant type
				if def, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), m.sim.UI.EditorPlantVariant+1); ok && def != nil {
					// Variant exists, but we're at 0, so wrap to max
					maxVariant := 0
					for i := 0; i < 10; i++ {
						if def, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), int16(i)); ok && def != nil {
							maxVariant = i
						}
					}
					m.sim.UI.EditorPlantVariant = int16(maxVariant)
				}
			}
			if def, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), m.sim.UI.EditorPlantVariant); ok && def != nil {
				fmt.Printf("Editor plant variant set to: %d (%s)\n", m.sim.UI.EditorPlantVariant, def.Name)
			}
		}
		if rl.IsKeyPressed(rl.KeyE) {
			// Increase variant
			m.sim.UI.EditorPlantVariant++
			if def, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), m.sim.UI.EditorPlantVariant); ok && def != nil {
				fmt.Printf("Editor plant variant set to: %d (%s)\n", m.sim.UI.EditorPlantVariant, def.Name)
			} else {
				// Variant doesn't exist, wrap to 0
				m.sim.UI.EditorPlantVariant = 0
				if def, ok := data.GetPlantDefinition(int(m.sim.UI.EditorPlantType), 0); ok && def != nil {
					fmt.Printf("Editor plant variant set to: 0 (%s)\n", def.Name)
				}
			}
		}

	case sim.EditorModeStructures:
		// Number keys select structure type
		if rl.IsKeyPressed(rl.KeyOne) {
			m.sim.UI.EditorStructureType = sim.Well
			fmt.Println("Editor structure type set to: Well")
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			m.sim.UI.EditorStructureType = sim.Bed
			fmt.Println("Editor structure type set to: Bed")
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			m.sim.UI.EditorStructureType = sim.Furniture
			fmt.Println("Editor structure type set to: Furniture")
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			m.sim.UI.EditorStructureType = sim.Workshop
			fmt.Println("Editor structure type set to: Workshop")
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			m.sim.UI.EditorStructureType = sim.Storage
			fmt.Println("Editor structure type set to: Storage")
		}
	}

	// Handle WASD movement (works in all modes)
	if rl.IsKeyDown(rl.KeyW) {
		m.sim.Player.MoveUp(deltaTime)
	}
	if rl.IsKeyDown(rl.KeyS) {
		m.sim.Player.MoveDown(deltaTime)
	}
	if rl.IsKeyDown(rl.KeyA) {
		m.sim.Player.MoveLeft(deltaTime)
	}
	if rl.IsKeyDown(rl.KeyD) {
		m.sim.Player.MoveRight(deltaTime)
	}
}
