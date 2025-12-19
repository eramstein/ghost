package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// structureTypeString converts StructureType to a readable string
func structureTypeString(st sim.StructureType) string {
	switch st {
	case sim.Well:
		return "Well"
	case sim.Bed:
		return "Bed"
	case sim.Furniture:
		return "Furniture"
	case sim.Workshop:
		return "Workshop"
	case sim.Storage:
		return "Storage"
	default:
		return fmt.Sprintf("Unknown (%d)", int(st))
	}
}

func DrawStructure(renderer *Renderer, structure sim.Structure) {
	centerX := float32(structure.Position.X*config.TileSize + config.TileSize/2)
	centerY := float32(structure.Position.Y*config.TileSize + config.TileSize/2)
	size := float32(config.TileSize) / 2
	halfSize := size / 2
	rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), ColorStructure)
}

// DrawStructureDetails renders structure info starting at (x, y) and returns
// the updated y position after drawing.
func DrawStructureDetails(renderer *Renderer, structure *sim.Structure, x, y int) int {
	if structure == nil {
		return y
	}

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)

	// Title - Structure Type
	titleText := structureTypeString(structure.StructureType)
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)

	// ID, type & variant
	renderer.RenderTextWithColor(fmt.Sprintf("ID: %d", structure.ID), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("Type: %d", structure.StructureType), x, y, rl.NewColor(200, 200, 200, 255))
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
