package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawPlants(renderer *Renderer, plants []sim.Plant) {
	for _, plant := range plants {
		centerX := float32(plant.Position.X*config.TileSize + config.TileSize/2)
		centerY := float32(plant.Position.Y*config.TileSize + config.TileSize/2)
		size := float32(10)

		// Create an upward-pointing triangle
		v1 := rl.Vector2{X: centerX, Y: centerY - size}        // Top vertex
		v2 := rl.Vector2{X: centerX - size, Y: centerY + size} // Bottom left
		v3 := rl.Vector2{X: centerX + size, Y: centerY + size} // Bottom right

		rl.DrawTriangle(v1, v2, v3, ColorPlant)
	}
}

func DrawPlantDetails(renderer *Renderer, plant *sim.Plant) {
	if plant == nil {
		return
	}

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	// Panel dimensions
	panelWidth := int32(300)
	panelX := int32(screenWidth) - panelWidth
	panelY := int32(0)
	panelHeight := int32(screenHeight)

	// Panel background (semi-transparent dark overlay)
	rl.DrawRectangle(panelX, panelY, panelWidth, panelHeight, rl.NewColor(20, 25, 30, 240))

	// Panel border
	rl.DrawRectangleLines(panelX, panelY, panelWidth, panelHeight, ColorBorder)

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)
	padding := int32(10)
	x := int(panelX + padding)
	y := int(panelY + padding)

	// Title - Plant Name (from data definition if available)
	titleText := "Plant"
	if def, ok := data.GetPlantDefinition(int(plant.PlantType), plant.Variant); ok && def != nil && def.Name != "" {
		titleText = def.Name
	}
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight) + 5

	// Separator line
	rl.DrawLine(int32(x), int32(y), panelX+panelWidth-padding, int32(y), ColorBorder)
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
}
