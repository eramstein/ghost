package input

import (
	"fmt"
	"gociv/pkg/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Manager) HandleMouse(deltaTime float32) {
	// Update mouse state
	m.mousePosition = rl.GetMousePosition()
	m.leftPressed = rl.IsMouseButtonPressed(rl.MouseLeftButton)
	m.rightPressed = rl.IsMouseButtonPressed(rl.MouseRightButton)

	// Handle mouse clicks
	if m.leftPressed && !m.sim.UI.EditMode {
		worldX, worldY := m.ScreenToWorld(rl.GetMouseX(), rl.GetMouseY())
		m.sim.UI.SelectedCharacterID = -1
		for _, character := range m.sim.Characters {
			if character.WorldPosition.X >= worldX-config.TileSize/2 && character.WorldPosition.X <= worldX+config.TileSize/2 && character.WorldPosition.Y >= worldY-config.TileSize/2 && character.WorldPosition.Y <= worldY+config.TileSize/2 {
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
