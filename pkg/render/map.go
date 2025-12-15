package render

import (
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawMap(renderer *Renderer, simData *sim.Sim) {
	tiles := simData.Tiles

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

			DrawTile(renderer, simData, tiles[i])

			// // Draw tile coordinates as text
			// coordText := fmt.Sprintf("%d,%d", tiles[i].Position.X, tiles[i].Position.Y)
			// rl.DrawText(coordText,
			// 	int32(tiles[i].Position.X*config.TileSize)+2,
			// 	int32(tiles[i].Position.Y*config.TileSize)+2,
			// 	8,
			// 	rl.White)

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

func DrawTile(renderer *Renderer, simData *sim.Sim, tile sim.Tile) {
	// terrain
	rl.DrawRectangle(
		int32(tile.Position.X*config.TileSize),
		int32(tile.Position.Y*config.TileSize),
		config.TileSize,
		config.TileSize,
		TileTypeColors[tile.Type],
	)
	// items
	for i, itemID := range tile.Items {
		baseX := float32(tile.Position.X * config.TileSize)
		baseY := float32(tile.Position.Y * config.TileSize)

		item := simData.GetItem(itemID)

		itemSize := float32(5)
		// Distribute items on a 3x3 grid within the tile
		step := float32(config.TileSize) / 3.0
		col := i % 3
		row := (i / 3) % 3
		centerX := baseX + (step / 2.0) + float32(col)*step
		centerY := baseY + (step / 2.0) + float32(row)*step

		rl.DrawRectanglePro(
			rl.Rectangle{X: centerX, Y: centerY, Width: itemSize, Height: itemSize},
			rl.Vector2{X: itemSize / 2.0, Y: itemSize / 2.0},
			45.0,
			ItemTypeColors[item.Type],
		)
	}
}
