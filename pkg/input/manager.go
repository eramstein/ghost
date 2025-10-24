package input

import (
	"fmt"
	"gociv/pkg/sim"
	"gociv/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Manager struct {
	sim     *sim.Sim
	camera  *rl.Camera2D
	console *Console
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
	// Handle console toggle
	if rl.IsKeyPressed(rl.KeyF11) {
		m.console.Toggle()
	}

	// If console is open, handle console input
	if m.console.IsOpen() {
		m.console.HandleInput()
		return // Don't process other inputs when console is open
	}

	m.handleKeyboard(deltaTime)
	m.handleMouse(deltaTime)
}

// handleKeyboard processes keyboard input events
func (m *Manager) handleKeyboard(deltaTime float32) {
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

	// F12 - Toggle EditMode
	if rl.IsKeyPressed(rl.KeyF12) {
		m.sim.UI.EditMode = !m.sim.UI.EditMode
		if m.sim.UI.EditMode {
			fmt.Println("EditMode enabled")
		} else {
			fmt.Println("EditMode disabled")
		}
	}

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

// screenToWorld converts screen coordinates to world coordinates using camera
func (m *Manager) screenToWorld(screenX, screenY int32) (float32, float32) {
	// Convert screen coordinates to world coordinates using camera
	screenPos := rl.Vector2{X: float32(screenX), Y: float32(screenY)}
	worldPos := rl.GetScreenToWorld2D(screenPos, *m.camera)

	return worldPos.X, worldPos.Y
}

// worldToTile converts world coordinates to tile coordinates
func (m *Manager) worldToTile(worldX, worldY float32) sim.TilePosition {
	tileX := int(worldX / float32(sim.TILE_SIZE))
	tileY := int(worldY / float32(sim.TILE_SIZE))

	return sim.TilePosition{
		X: tileX,
		Y: tileY,
	}
}

// handleMouse processes mouse input events
func (m *Manager) handleMouse(deltaTime float32) {
	// Handle mouse clicks (only in EditMode)
	if rl.IsMouseButtonDown(rl.MouseLeftButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.screenToWorld(mouseX, mouseY)
		tile := m.sim.GetTileAt(m.worldToTile(worldX, worldY))
		fmt.Printf("Clicked tile position: (%d, %d)\n", tile.Position.X, tile.Position.Y)
		tile.UpdateType(sim.TileTypeWall)
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) && m.sim.UI.EditMode {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		worldX, worldY := m.screenToWorld(mouseX, mouseY)
		tile := m.sim.GetTileAt(m.worldToTile(worldX, worldY))
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

// GetConsole returns the console instance for rendering
func (m *Manager) GetConsole() *Console {
	return m.console
}
