package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawStructures(renderer *Renderer, structures []sim.Structure) {
	for _, structure := range structures {
		centerX := float32(structure.Position.X*config.TileSize + config.TileSize/2)
		centerY := float32(structure.Position.Y*config.TileSize + config.TileSize/2)
		size := float32(config.TileSize) / 2
		halfSize := size / 2

		rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), ColorStructure)
	}
}

// DrawStructureDetails renders structure info starting at (x, y) and returns
// the updated y position after drawing.
func DrawStructureDetails(renderer *Renderer, structure *sim.Structure, x, y int) int {
	if structure == nil {
		return y
	}

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)

	// Title - Structure Name (from data definition if available)
	titleText := "Structure"
	if def, ok := data.GetStructureDefinition(int(structure.StructureType), structure.Variant); ok && def != nil && def.Name != "" {
		titleText = def.Name
	}
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)

	// ID, type & variant
	renderer.RenderTextWithColor(fmt.Sprintf("ID: %d", structure.ID), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("Type: %d", structure.StructureType), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("Variant: %d", structure.Variant), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Position
	renderer.RenderTextWithColor("Position:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Tile: (%d, %d)", structure.Position.X, structure.Position.Y), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Condition
	renderer.RenderTextWithColor("Condition:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  %d%%", structure.Condition), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Build Progress
	if structure.BuildProgress < 100 {
		renderer.RenderTextWithColor("Build Progress:", x, y, rl.NewColor(255, 255, 255, 255))
		y += int(lineHeight)
		renderer.RenderTextWithColor(fmt.Sprintf("  %d%%", structure.BuildProgress), x, y, rl.NewColor(200, 200, 200, 255))
		y += int(lineHeight)
	}

	// Owner
	renderer.RenderTextWithColor("Owner:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	if structure.Owner >= 0 {
		renderer.RenderTextWithColor(fmt.Sprintf("  Character ID: %d", structure.Owner), x, y, rl.NewColor(200, 200, 200, 255))
	} else {
		renderer.RenderTextWithColor("  None", x, y, rl.NewColor(200, 200, 200, 255))
	}
	y += int(lineHeight)

	return y
}
