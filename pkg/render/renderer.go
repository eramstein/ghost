package render

import (
	"gociv/pkg/input"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	Camera      rl.Camera2D
	Console     *input.Console
	FontManager *FontManager
	DefaultFont rl.Font
}

func NewRenderer(simData *sim.Sim, console *input.Console) *Renderer {
	r := &Renderer{
		Camera: rl.Camera2D{
			Target:   rl.Vector2{X: simData.Player.WorldPosition.X, Y: simData.Player.WorldPosition.Y},
			Offset:   rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
			Rotation: 0,
			Zoom:     1,
		},
		Console:     console,
		FontManager: NewFontManager(),
		DefaultFont: rl.GetFontDefault(),
	}
	return r
}

func (r *Renderer) Render(simData *sim.Sim) {

	// Update camera to follow player
	r.Camera.Target = rl.Vector2{X: simData.Player.WorldPosition.X, Y: simData.Player.WorldPosition.Y}

	rl.ClearBackground(ColorBackground)
	rl.BeginMode2D(r.Camera)
	DrawMap(r, simData)
	DrawPlants(r, simData.GetPlants())
	DrawStructures(r, simData.GetStructures())
	DrawPlayer(r, simData.Player)
	DrawCharacters(r, simData.Characters)
	rl.EndMode2D()

	// Draw UI elements (outside of 2D mode)
	r.DrawUI(simData)
}

// DrawUI renders UI elements like EditMode indicator
func (r *Renderer) DrawUI(simData *sim.Sim) {
	DisplayTime(r, &simData.Calendar)

	if simData.UI.EditMode {
		// Draw EditMode indicator in top-left corner
		rl.DrawText("EDIT MODE", 10, 10, 20, ColorEditMode)

		// Draw current EditorTileType below EditMode indicator
		tileTypeText := "Tile Type: " + simData.UI.EditorTileType.String()
		rl.DrawText(tileTypeText, 10, 35, 16, ColorEditMode)
	}

	// Unified side panel with stacked tile/character/plant details
	DrawSidePanel(r, simData)

	// Draw console if open
	if r.Console != nil && r.Console.IsOpen() {
		DrawConsole(r.Console)
	}
}

// todo - unload textures
func (r *Renderer) Close() {
}

// RenderText renders text at a specific position
func (r *Renderer) RenderText(text string, x, y int) {
	rl.DrawTextEx(r.DefaultFont, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(r.DefaultFont.BaseSize), 1.0, rl.Black)
}

// RenderTextWithColor renders text at a specific position with a specific color
func (r *Renderer) RenderTextWithColor(text string, x, y int, color rl.Color) {
	rl.DrawTextEx(r.DefaultFont, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(r.DefaultFont.BaseSize), 1.0, color)
}
