package input

import (
	"fmt"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Manager) HandleMouseEditor(deltaTime float32) {
	// Handle mouse clicks (only in EditMode)
	if rl.IsMouseButtonDown(rl.MouseLeftButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.ScreenToWorld(mouseX, mouseY)
		tile := m.sim.GetTileAt(m.WorldToTile(worldX, worldY))
		fmt.Printf("Clicked tile position: (%d, %d)\n", tile.Position.X, tile.Position.Y)
		tile.UpdateType(m.sim.UI.EditorTileType)
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.ScreenToWorld(mouseX, mouseY)
		tile := m.sim.GetTileAt(m.WorldToTile(worldX, worldY))
		fmt.Printf("Clicked tile position: (%d, %d)\n", tile.Position.X, tile.Position.Y)
		tile.UpdateType(sim.TileTypeEmpty)
	}

	// Handle mouse wheel
	wheelMove := rl.GetMouseWheelMove()
	if wheelMove != 0 {
		fmt.Printf("Mouse wheel: %f\n", wheelMove)
		// Add mouse wheel handling logic here
	}
}
