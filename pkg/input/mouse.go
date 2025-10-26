package input

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Manager) HandleMouse(deltaTime float32) {
	// Handle mouse clicks (only in EditMode)
	if rl.IsMouseButtonDown(rl.MouseLeftButton) && m.sim.UI.EditMode {
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) && m.sim.UI.EditMode {
	}

	// Handle mouse wheel
	wheelMove := rl.GetMouseWheelMove()
	if wheelMove != 0 {
		fmt.Printf("Mouse wheel: %f\n", wheelMove)
		// Add mouse wheel handling logic here
	}
}
