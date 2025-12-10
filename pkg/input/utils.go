package input

import (
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// screenToWorld converts screen coordinates to world coordinates using camera
func (m *Manager) ScreenToWorld(screenX, screenY int32) (float32, float32) {
	// Convert screen coordinates to world coordinates using camera
	screenPos := rl.Vector2{X: float32(screenX), Y: float32(screenY)}
	worldPos := rl.GetScreenToWorld2D(screenPos, *m.camera)

	return worldPos.X, worldPos.Y
}

// worldToTile converts world coordinates to tile coordinates
func (m *Manager) WorldToTile(worldX, worldY float32) sim.TilePosition {
	tileX := int(worldX / float32(config.TileSize))
	tileY := int(worldY / float32(config.TileSize))

	return sim.TilePosition{
		X: tileX,
		Y: tileY,
	}
}
