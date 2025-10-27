package render

import (
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawMap(renderer *Renderer, tiles []sim.Tile) {
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
	padding := float32(sim.TILE_SIZE)
	leftBound -= padding
	rightBound += padding
	topBound -= padding
	bottomBound += padding

	for i := range tiles {
		// Calculate tile world position
		tileWorldX := float32(tiles[i].Position.X * sim.TILE_SIZE)
		tileWorldY := float32(tiles[i].Position.Y * sim.TILE_SIZE)

		// Check if tile is within visible bounds
		if tileWorldX >= leftBound && tileWorldX <= rightBound &&
			tileWorldY >= topBound && tileWorldY <= bottomBound {

			rl.DrawRectangle(
				int32(tiles[i].Position.X*sim.TILE_SIZE),
				int32(tiles[i].Position.Y*sim.TILE_SIZE),
				sim.TILE_SIZE,
				sim.TILE_SIZE,
				TileTypeColors[tiles[i].Type],
			)

			// // Draw tile coordinates as text
			// coordText := fmt.Sprintf("%d,%d", tiles[i].Position.X, tiles[i].Position.Y)
			// rl.DrawText(coordText,
			// 	int32(tiles[i].Position.X*sim.TILE_SIZE)+2,
			// 	int32(tiles[i].Position.Y*sim.TILE_SIZE)+2,
			// 	8,
			// 	rl.White)

		}
	}

	// Draw region border as a single rectangle outline
	regionSize := float32(sim.REGION_SIZE * sim.TILE_SIZE)
	rl.DrawRectangleLinesEx(
		rl.Rectangle{X: 0, Y: 0, Width: regionSize, Height: regionSize},
		2.0,
		ColorBorder,
	)
}
