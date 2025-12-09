package input

import (
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ScreenToTileCoordinates converts screen coordinates to tile coordinates
func (m *Manager) ScreenToTileCoordinates(screenPos rl.Vector2) sim.TilePosition {
	if m.camera == nil {
		return sim.TilePosition{}
	}
	worldPos := rl.GetScreenToWorld2D(screenPos, *m.camera)
	tileX := int(worldPos.X) / sim.TILE_SIZE
	tileY := int(worldPos.Y) / sim.TILE_SIZE
	return sim.TilePosition{X: tileX, Y: tileY}
}
