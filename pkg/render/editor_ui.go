package render

import (
	"fmt"
	"gociv/pkg/data"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Editor UI colors
var (
	ColorEditorBackground = rl.Color{R: 20, G: 20, B: 20, A: 230}    // Dark semi-transparent background
	ColorEditorText       = rl.Color{R: 255, G: 255, B: 255, A: 255} // White text
	ColorEditorTitle      = rl.Color{R: 255, G: 255, B: 0, A: 255}   // Yellow for title
	ColorEditorLabel      = rl.Color{R: 200, G: 200, B: 200, A: 255} // Light gray for labels
	ColorEditorValue      = rl.Color{R: 100, G: 200, B: 255, A: 255} // Light blue for values
)

// DrawEditorUI renders the editor mode UI panel with background
func DrawEditorUI(renderer *Renderer, simData *sim.Sim) {
	if !simData.UI.EditMode {
		return
	}

	// Panel dimensions and position
	panelX := int32(10)
	panelY := int32(100)
	panelWidth := int32(350)
	padding := int32(15)

	// Text position (inside panel with padding)
	textX := panelX + padding
	startY := panelY + padding

	// Use default font
	font := renderer.DefaultFont
	fontSize := float32(font.BaseSize) * 1.3 // Increase base font size by 30%
	titleSize := fontSize * 1.25

	// Calculate content height based on mode
	contentHeight := padding * 2           // Top and bottom padding
	contentHeight += int32(titleSize) + 10 // Title
	contentHeight += int32(fontSize) + 8   // Mode line

	// Add mode-specific height
	switch simData.UI.EditorMode {
	case sim.EditorModeTiles:
		contentHeight += int32(fontSize) + 8      // Tile type
		contentHeight += int32(fontSize*0.85) + 8 // Help text
	case sim.EditorModePlants:
		contentHeight += int32(fontSize) + 8 // Plant type
		contentHeight += int32(fontSize) + 8 // Variant
		// Check if plant name exists
		if def, ok := data.GetPlantDefinition(int(simData.UI.EditorPlantType), simData.UI.EditorPlantVariant); ok && def != nil {
			contentHeight += int32(fontSize) + 8 // Name
		}
		contentHeight += int32(fontSize*0.85) + 8 // Help text
	case sim.EditorModeStructures:
		contentHeight += int32(fontSize) + 8      // Structure type
		contentHeight += int32(fontSize*0.85) + 8 // Help text
	}

	contentHeight += 10 + int32(fontSize*0.85) + 8 // Mode switching instructions

	panelHeight := contentHeight

	// Draw background panel
	rl.DrawRectangle(panelX, panelY, panelWidth, panelHeight, ColorEditorBackground)
	// Draw border
	rl.DrawRectangleLines(panelX, panelY, panelWidth, panelHeight, ColorBorder)

	yPos := startY

	// Draw title
	titleText := "EDIT MODE"
	rl.DrawTextEx(font, titleText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, titleSize, 1.0, ColorEditorTitle)
	yPos += int32(titleSize) + 10

	// Draw current editor mode
	modeText := "Mode: " + simData.UI.EditorMode.String()
	rl.DrawTextEx(font, modeText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorText)
	yPos += int32(fontSize) + 8

	// Draw mode-specific selection info
	switch simData.UI.EditorMode {
	case sim.EditorModeTiles:
		labelText := "Tile Type:"
		rl.DrawTextEx(font, labelText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorLabel)
		valueText := simData.UI.EditorTileType.String()
		valueX := textX + int32(fontSize*float32(len(labelText)+1))
		rl.DrawTextEx(font, valueText, rl.Vector2{X: float32(valueX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorValue)
		yPos += int32(fontSize) + 8

		helpText := "Keys: 1-5 to select tile type"
		rl.DrawTextEx(font, helpText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize*0.85, 1.0, ColorEditorLabel)
		yPos += int32(fontSize*0.85) + 8

	case sim.EditorModePlants:
		labelText := "Plant Type:"
		rl.DrawTextEx(font, labelText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorLabel)
		valueText := fmt.Sprintf("%d", simData.UI.EditorPlantType)
		valueX := textX + int32(fontSize*float32(len(labelText)+1))
		rl.DrawTextEx(font, valueText, rl.Vector2{X: float32(valueX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorValue)
		yPos += int32(fontSize) + 8

		labelText2 := "Variant:"
		rl.DrawTextEx(font, labelText2, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorLabel)
		valueText2 := fmt.Sprintf("%d", simData.UI.EditorPlantVariant)
		valueX2 := textX + int32(fontSize*float32(len(labelText2)+1))
		rl.DrawTextEx(font, valueText2, rl.Vector2{X: float32(valueX2), Y: float32(yPos)}, fontSize, 1.0, ColorEditorValue)
		yPos += int32(fontSize) + 8

		// Try to get plant name from data
		if def, ok := data.GetPlantDefinition(int(simData.UI.EditorPlantType), simData.UI.EditorPlantVariant); ok && def != nil {
			labelText3 := "Name:"
			rl.DrawTextEx(font, labelText3, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorLabel)
			valueText3 := def.Name
			valueX3 := textX + int32(fontSize*float32(len(labelText3)+1))
			rl.DrawTextEx(font, valueText3, rl.Vector2{X: float32(valueX3), Y: float32(yPos)}, fontSize, 1.0, ColorEditorValue)
			yPos += int32(fontSize) + 8
		}

		helpText := "Keys: 1 to select type, Q/W to cycle variants"
		rl.DrawTextEx(font, helpText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize*0.85, 1.0, ColorEditorLabel)
		yPos += int32(fontSize*0.85) + 8

	case sim.EditorModeStructures:
		labelText := "Structure Type:"
		rl.DrawTextEx(font, labelText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorLabel)
		valueText := structureTypeString(simData.UI.EditorStructureType)
		valueX := textX + int32(fontSize*float32(len(labelText)+1))
		rl.DrawTextEx(font, valueText, rl.Vector2{X: float32(valueX), Y: float32(yPos)}, fontSize, 1.0, ColorEditorValue)
		yPos += int32(fontSize) + 8

		helpText := "Keys: 1-5 to select structure type"
		rl.DrawTextEx(font, helpText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize*0.85, 1.0, ColorEditorLabel)
		yPos += int32(fontSize*0.85) + 8
	}

	// Draw mode switching instructions
	yPos += 10
	helpText := "T = Tiles, P = Plants, B = Structures"
	rl.DrawTextEx(font, helpText, rl.Vector2{X: float32(textX), Y: float32(yPos)}, fontSize*0.85, 1.0, ColorEditorLabel)
}
