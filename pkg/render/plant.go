package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawPlant(renderer *Renderer, plant sim.Plant) {
	centerX := float32(plant.Position.X*config.TileSize + config.TileSize/2)
	centerY := float32(plant.Position.Y*config.TileSize + config.TileSize/2)
	size := float32(10)

	// Create an upward-pointing triangle
	v1 := rl.Vector2{X: centerX, Y: centerY - size}        // Top vertex
	v2 := rl.Vector2{X: centerX - size, Y: centerY + size} // Bottom left
	v3 := rl.Vector2{X: centerX + size, Y: centerY + size} // Bottom right

	rl.DrawTriangle(v1, v2, v3, ColorPlant)
}

// DrawPlantDetails renders plant info starting at (x, y) and returns
// the updated y position after drawing.
func DrawPlantDetails(renderer *Renderer, plant *sim.Plant, x, y int) int {
	if plant == nil {
		return y
	}

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)

	// Title - Plant Name (from data definition if available)
	titleText := "Plant"
	if def, ok := data.GetPlantDefinition(int(plant.PlantType), plant.Variant); ok && def != nil && def.Name != "" {
		titleText = def.Name
	}
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)

	// ID, type & variant
	renderer.RenderTextWithColor(fmt.Sprintf("ID: %d", plant.ID), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("Type: %d", plant.PlantType), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("Variant: %d", plant.Variant), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Position
	renderer.RenderTextWithColor("Position:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Tile: (%d, %d)", plant.Position.X, plant.Position.Y), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Growth
	renderer.RenderTextWithColor("Growth:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Stage: %d%%", plant.GrowthStage), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Rate: %d / update", plant.GrowthRate), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Production
	renderer.RenderTextWithColor("Production:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Item Type: %d", plant.Produces.Type), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Variant: %d", plant.Produces.Variant), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Stage: %d%%", plant.Produces.ProductionStage), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Rate: %d / update", plant.Produces.ProductionRate), x, y, rl.NewColor(200, 200, 200, 255))

	return y
}
