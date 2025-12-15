package input

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Manager) HandleMouse(deltaTime float32) {
	// Update mouse state
	m.mousePosition = rl.GetMousePosition()
	m.leftPressed = rl.IsMouseButtonPressed(rl.MouseLeftButton)
	m.rightPressed = rl.IsMouseButtonPressed(rl.MouseRightButton)

	// Handle mouse clicks
	if m.leftPressed && !m.sim.UI.EditMode {
		m.ClearSelections()
		worldX, worldY := m.ScreenToWorld(rl.GetMouseX(), rl.GetMouseY())
		tilePosition := m.WorldToTile(worldX, worldY)
		tileID := m.sim.GetTileIDFromPosition(tilePosition)
		m.SelectTile(tileID)
		for _, character := range m.sim.Characters {
			if character.TilePosition.X == tilePosition.X && character.TilePosition.Y == tilePosition.Y {
				m.SelectCharacter(character.ID)
				break
			}
		}
		for _, plant := range m.sim.GetPlants() {
			if plant.Position.X == tilePosition.X && plant.Position.Y == tilePosition.Y {
				// store the plant ID in SelectedPlantIndex
				m.SelectPlant(plant.ID)
				break
			}
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) && !m.sim.UI.EditMode {
	}

	// Handle mouse wheel
	wheelMove := rl.GetMouseWheelMove()
	if wheelMove != 0 {
		fmt.Printf("Mouse wheel: %f\n", wheelMove)
		// Add mouse wheel handling logic here
	}
}
