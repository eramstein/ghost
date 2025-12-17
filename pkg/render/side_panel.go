package render

import (
	"fmt"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawSidePanel renders a single side panel on the right of the screen and,
// if present, shows details about the selected tile, character, plant and structure
// stacked one below the other.
func DrawSidePanel(renderer *Renderer, simData *sim.Sim) {
	hasTile := simData.UI.SelectedTileIndex != -1 &&
		simData.UI.SelectedTileIndex >= 0 &&
		simData.UI.SelectedTileIndex < len(simData.Tiles)
	hasCharacter := simData.UI.SelectedCharacterIndex != -1 &&
		simData.UI.SelectedCharacterIndex >= 0 &&
		simData.UI.SelectedCharacterIndex < len(simData.Characters)
	hasPlant := simData.UI.SelectedPlantIndex != -1 && simData.PlantManager != nil
	hasStructure := simData.UI.SelectedStructureIndex != -1 && simData.StructureManager != nil

	// Nothing selected: don't draw a panel at all.
	if !hasTile && !hasCharacter && !hasPlant && !hasStructure {
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
	// Make the panel fully opaque for maximum readability
	rl.DrawRectangle(panelX, panelY, panelWidth, panelHeight, rl.NewColor(20, 25, 30, 255))

	// Panel border
	rl.DrawRectangleLines(panelX, panelY, panelWidth, panelHeight, ColorBorder)

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)
	padding := int32(10)
	x := int(panelX + padding)
	y := int(panelY + padding)

	// Helper for section separators (clear visual delimitation, no titles)
	drawSectionSeparator := func() {
		// Small top margin before the separator
		y += int(lineHeight / 2)

		// Thicker, more visible separator bar
		separatorHeight := int32(3)
		rl.DrawRectangle(
			int32(x),
			int32(y),
			panelX+panelWidth-padding-int32(x),
			separatorHeight,
			ColorBorder,
		)

		// Space between the separator and the section content
		y += int(separatorHeight) + int(lineHeight/2)
	}

	// Tile details
	if hasTile {
		tile := &simData.Tiles[simData.UI.SelectedTileIndex]
		drawSectionSeparator()
		y = DrawTileDetails(renderer, tile, x, y)

		y += int(lineHeight) // Extra spacing after section
	}

	// Character details
	if hasCharacter {
		character := &simData.Characters[simData.UI.SelectedCharacterIndex]
		drawSectionSeparator()
		y = DrawCharacterDetails(renderer, character, x, y)

		y += int(lineHeight) // Extra spacing after section
	}

	// Plant details
	if hasPlant {
		plant := simData.GetPlantByID(simData.UI.SelectedPlantIndex)
		if plant != nil {
			drawSectionSeparator()
			y = DrawPlantDetails(renderer, plant, x, y)
		}
	}

	// Structure details
	if hasStructure {
		structure := simData.GetStructureByID(simData.UI.SelectedStructureIndex)
		if structure != nil {
			drawSectionSeparator()
			y = DrawStructureDetails(renderer, structure, x, y)
		}
	}
}

// DrawTileDetails renders tile info starting at (x, y) and returns
// the updated y position after drawing.
func DrawTileDetails(renderer *Renderer, tile *sim.Tile, x, y int) int {
	if tile == nil {
		return y
	}

	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)

	renderer.RenderTextWithColor(
		fmt.Sprintf("Position: (%d, %d)", tile.Position.X, tile.Position.Y),
		x, y, rl.NewColor(200, 200, 200, 255),
	)
	y += int(lineHeight)

	renderer.RenderTextWithColor(
		fmt.Sprintf("Type: %s", tile.Type.String()),
		x, y, rl.NewColor(200, 200, 200, 255),
	)
	y += int(lineHeight)

	renderer.RenderTextWithColor(
		fmt.Sprintf("Move cost: %.1f", tile.MoveCost),
		x, y, rl.NewColor(200, 200, 200, 255),
	)
	y += int(lineHeight)

	if len(tile.Items) > 0 {
		renderer.RenderTextWithColor(
			fmt.Sprintf("Items on tile: %d", len(tile.Items)),
			x, y, rl.NewColor(200, 200, 200, 255),
		)
		y += int(lineHeight)
	}

	return y
}
