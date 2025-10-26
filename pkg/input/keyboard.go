package input

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// handleKeyboard processes keyboard input events
func (m *Manager) HandleKeyboard(deltaTime float32) {

	// Spacebar - Toggle Pause
	if rl.IsKeyPressed(rl.KeySpace) {
		m.sim.UI.Pause = !m.sim.UI.Pause
		if m.sim.UI.Pause {
			fmt.Println("Game paused")
		} else {
			fmt.Println("Game resumed")
		}
	}

	// Handle WASD movement
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
