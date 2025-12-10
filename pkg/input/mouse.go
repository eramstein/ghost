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
		tilePos := m.ScreenToTileCoordinates(m.mousePosition)
		for _, character := range m.sim.Characters {
			if character.TilePosition.X == tilePos.X && character.TilePosition.Y == tilePos.Y {
				m.sim.UI.SelectedCharacterID = character.ID
			}
		}
		fmt.Printf("Char clicked: %d\n", m.sim.UI.SelectedCharacterID)
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
