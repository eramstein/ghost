package render

import (
	"gociv/pkg/input"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	Camera  rl.Camera2D
	Console *input.Console
}

func NewRenderer(simData *sim.Sim, console *input.Console) *Renderer {
	r := &Renderer{
		Camera: rl.Camera2D{
			Target:   rl.Vector2{X: simData.Player.WorldPosition.X, Y: simData.Player.WorldPosition.Y},
			Offset:   rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
			Rotation: 0,
			Zoom:     1,
		},
		Console: console,
	}
	return r
}

func (r *Renderer) Render(simData *sim.Sim) {

	// Update camera to follow player
	r.Camera.Target = rl.Vector2{X: simData.Player.WorldPosition.X, Y: simData.Player.WorldPosition.Y}

	rl.ClearBackground(ColorBackground)
	rl.BeginMode2D(r.Camera)
	DrawMap(r, simData.Tiles)
	//DrawMapDebug(r, simData.Tiles, simData.Characters)
	DrawPlayer(r, simData.Player)
	DrawCharacters(r, simData.Characters)
	rl.EndMode2D()

	// Draw UI elements (outside of 2D mode)
	DrawUI(simData.UI)

	// Draw console if open
	if r.Console != nil && r.Console.IsOpen() {
		DrawConsole(r.Console)
	}
}

// DrawUI renders UI elements like EditMode indicator
func DrawUI(ui sim.UIState) {
	if ui.EditMode {
		// Draw EditMode indicator in top-left corner
		rl.DrawText("EDIT MODE", 10, 10, 20, ColorEditMode)
	}
}

// todo - unload textures
func (r *Renderer) Close() {
}
