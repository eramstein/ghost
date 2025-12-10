package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawMapDebug(renderer *Renderer, tiles []sim.Tile, characters []sim.Character) {
	// Calculate visible area based on camera
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	// Calculate world coordinates of screen bounds
	// Account for camera zoom
	halfScreenWidth := screenWidth / (2 * renderer.Camera.Zoom)
	halfScreenHeight := screenHeight / (2 * renderer.Camera.Zoom)

	// Calculate visible bounds in world coordinates
	leftBound := renderer.Camera.Target.X - halfScreenWidth
	rightBound := renderer.Camera.Target.X + halfScreenWidth
	topBound := renderer.Camera.Target.Y - halfScreenHeight
	bottomBound := renderer.Camera.Target.Y + halfScreenHeight

	// Add some padding to avoid edge artifacts
	padding := float32(config.TileSize)
	leftBound -= padding
	rightBound += padding
	topBound -= padding
	bottomBound += padding

	for i := range tiles {
		// Calculate tile world position
		tileWorldX := float32(tiles[i].Position.X * config.TileSize)
		tileWorldY := float32(tiles[i].Position.Y * config.TileSize)

		// Check if tile is within visible bounds
		if tileWorldX >= leftBound && tileWorldX <= rightBound &&
			tileWorldY >= topBound && tileWorldY <= bottomBound {

			color := ColorBackground
			if tiles[i].Type == sim.TileTypeWall {
				color = ColorWall
			}
			rl.DrawRectangle(
				int32(tiles[i].Position.X*config.TileSize),
				int32(tiles[i].Position.Y*config.TileSize),
				config.TileSize,
				config.TileSize,
				color,
			)
		}
	}

	// Collect all path tiles from all characters
	pathTiles := make(map[string]bool)
	for _, character := range characters {
		for _, pathTile := range character.Path {
			key := fmt.Sprintf("%d,%d", pathTile.X, pathTile.Y)
			pathTiles[key] = true
		}
	}

	// Draw path highlights
	for i := range tiles {
		// Calculate tile world position
		tileWorldX := float32(tiles[i].Position.X * config.TileSize)
		tileWorldY := float32(tiles[i].Position.Y * config.TileSize)

		// Check if tile is within visible bounds
		if tileWorldX >= leftBound && tileWorldX <= rightBound &&
			tileWorldY >= topBound && tileWorldY <= bottomBound {

			// Check if this tile is in any character's path
			key := fmt.Sprintf("%d,%d", tiles[i].Position.X, tiles[i].Position.Y)
			if pathTiles[key] {
				// Draw path highlight overlay
				rl.DrawRectangle(
					int32(tiles[i].Position.X*config.TileSize),
					int32(tiles[i].Position.Y*config.TileSize),
					config.TileSize,
					config.TileSize,
					ColorPath,
				)
			}
		}
	}

	// Draw region border as a single rectangle outline
	regionSize := float32(config.RegionSize * config.TileSize)
	rl.DrawRectangleLinesEx(
		rl.Rectangle{X: 0, Y: 0, Width: regionSize, Height: regionSize},
		2.0,
		ColorBorder,
	)
}
