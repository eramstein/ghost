package input

import (
	"fmt"
	"gociv/pkg/sim"
	"gociv/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Manager struct {
	sim           *sim.Sim
	camera        *rl.Camera2D
	console       *Console
	mousePosition rl.Vector2
	leftPressed   bool
	rightPressed  bool
}

// NewManager creates a new input manager
func NewManager(sim *sim.Sim, camera *rl.Camera2D) *Manager {
	return &Manager{
		sim:     sim,
		camera:  camera,
		console: NewConsole(sim),
	}
}

// HandleInputs processes all input events including keyboard and mouse
func (m *Manager) HandleInputs(deltaTime float32) {
	// F11 - Toggle Console
	if rl.IsKeyPressed(rl.KeyF11) {
		m.console.Toggle()
	}
	// F12 - Toggle EditMode
	if rl.IsKeyPressed(rl.KeyF12) {
		m.sim.UI.EditMode = !m.sim.UI.EditMode
	}

	// F5 - Save quicksave
	if rl.IsKeyPressed(rl.KeyF5) {
		err := utils.SaveSim(m.sim, "quicksave")
		if err != nil {
			fmt.Printf("Error saving game: %v\n", err)
		} else {
			fmt.Println("Game saved successfully!")
		}
	}

	// F4 - Load quicksave
	if rl.IsKeyPressed(rl.KeyF4) {
		loadedSim, err := utils.LoadSim("quicksave")
		if err != nil {
			fmt.Printf("Error loading game: %v\n", err)
		} else {
			// Replace current sim with loaded sim
			*m.sim = *loadedSim
			fmt.Println("Game loaded successfully!")
		}
	}

	// If console is open, handle console input
	if m.console.IsOpen() {
		m.console.HandleInput()
		return // Don't process other inputs when console is open
	}
	// Edit Mode
	if m.sim.UI.EditMode {
		m.HandleKeyboardEditor(deltaTime)
		m.HandleMouseEditor(deltaTime)
		// Play Mode
	} else {
		m.HandleKeyboard(deltaTime)
		m.HandleMouse(deltaTime)
	}

}

// GetConsole returns the console instance for rendering
func (m *Manager) GetConsole() *Console {
	return m.console
}
