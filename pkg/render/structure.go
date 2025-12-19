package render

import (
	"fmt"
	"gociv/pkg/config"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// StructureTypeColors maps each StructureType to its corresponding color
var StructureTypeColors = map[sim.StructureType]rl.Color{
	sim.Well:      {R: 104, G: 170, B: 214, A: 255}, // Blue (water-related)
	sim.Bed:       {R: 139, G: 115, B: 85, A: 255},  // Brown/beige (furniture)
	sim.Furniture: {R: 150, G: 150, B: 150, A: 255}, // Gray
	sim.Workshop:  {R: 255, G: 165, B: 0, A: 255},   // Orange (work/activity)
	sim.Storage:   {R: 72, G: 150, B: 72, A: 255},   // Green (storage/containers)
}

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

	// Get color for this structure type, default to ColorStructure if not found
	color, exists := StructureTypeColors[structure.StructureType]
	if !exists {
		color = ColorStructure
	}

	// Draw different shapes based on structure type
	switch structure.StructureType {
	case sim.Well:
		// Well: Circle (water source) with gray border (3px thick)
		borderColor := rl.Color{R: 150, G: 150, B: 150, A: 255}
		rl.DrawCircle(int32(centerX), int32(centerY), halfSize+3, borderColor)
		rl.DrawCircle(int32(centerX), int32(centerY), halfSize, color)
	case sim.Bed:
		// Bed: Rectangle (horizontal)
		rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize*0.6), int32(size), int32(size*0.6), color)
	case sim.Furniture:
		// Furniture: Square
		rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), color)
	case sim.Workshop:
		// Workshop: Triangle (pointing up)
		rl.DrawTriangle(
			rl.Vector2{X: centerX, Y: centerY - halfSize},            // Top point
			rl.Vector2{X: centerX - halfSize, Y: centerY + halfSize}, // Bottom left
			rl.Vector2{X: centerX + halfSize, Y: centerY + halfSize}, // Bottom right
			color,
		)
	case sim.Storage:
		// Storage: Rectangle with border (box-like)
		rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), color)
		rl.DrawRectangleLines(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), rl.Color{R: color.R - 30, G: color.G - 30, B: color.B - 30, A: 255})
	default:
		// Default: Rectangle
		rl.DrawRectangle(int32(centerX-halfSize), int32(centerY-halfSize), int32(size), int32(size), color)
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
