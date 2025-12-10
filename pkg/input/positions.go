package input

import (
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ScreenToTileCoordinates converts screen coordinates to tile coordinates
func (m *Manager) ScreenToTileCoordinates(screenPos rl.Vector2) sim.TilePosition {
	if m.camera == nil {
		return sim.TilePosition{}
	}
	worldPos := rl.GetScreenToWorld2D(screenPos, *m.camera)
	tileX := int(worldPos.X) / config.TileSize
	tileY := int(worldPos.Y) / config.TileSize
	return sim.TilePosition{X: tileX, Y: tileY}
}
