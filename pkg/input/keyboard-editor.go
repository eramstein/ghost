package input

import (
	"fmt"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// handleKeyboard processes keyboard input events
func (m *Manager) HandleKeyboardEditor(deltaTime float32) {
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
}
