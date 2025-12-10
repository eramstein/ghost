package render

import (
	"fmt"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCharacters(renderer *Renderer, characters []sim.Character) {
	for _, character := range characters {
		rl.DrawCircleV(
			rl.Vector2{X: character.WorldPosition.X, Y: character.WorldPosition.Y},
			10,
			ColorLife,
		)
	}
}

func DrawCharacterDetails(renderer *Renderer, character *sim.Character) {
	if character == nil {
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

	// Title - Character Name
	titleText := character.Name
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight) + 5

	// Separator line
	rl.DrawLine(int32(x), int32(y), panelX+panelWidth-padding, int32(y), ColorBorder)
	y += int(lineHeight)

	// ID
	renderer.RenderTextWithColor(fmt.Sprintf("ID: %d", character.ID), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Position
	renderer.RenderTextWithColor("Position:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  World: (%.1f, %.1f)", character.WorldPosition.X, character.WorldPosition.Y), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Tile: (%d, %d)", character.TilePosition.X, character.TilePosition.Y), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Needs
	renderer.RenderTextWithColor("Needs:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Food: %d", character.Needs.Food), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Water: %d", character.Needs.Water), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)
	renderer.RenderTextWithColor(fmt.Sprintf("  Sleep: %d", character.Needs.Sleep), x, y, rl.NewColor(200, 200, 200, 255))
	y += int(lineHeight)

	// Current Task
	renderer.RenderTextWithColor("Current Task:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	if character.Task.Type != sim.NoTaskType {
		taskTypeStr := taskTypeToString(character.Task.Type)
		renderer.RenderTextWithColor(fmt.Sprintf("  Type: %s", taskTypeStr), x, y, rl.NewColor(200, 200, 200, 255))
		y += int(lineHeight)
		renderer.RenderTextWithColor(fmt.Sprintf("  Progress: %.1f%%", character.Task.Progress*100), x, y, rl.NewColor(200, 200, 200, 255))
		y += int(lineHeight)
		if character.Task.TargetTile.X >= 0 && character.Task.TargetTile.Y >= 0 {
			renderer.RenderTextWithColor(fmt.Sprintf("  Target: (%d, %d)", character.Task.TargetTile.X, character.Task.TargetTile.Y), x, y, rl.NewColor(200, 200, 200, 255))
			y += int(lineHeight)
		}
		if character.Task.TargetItemID >= 0 {
			renderer.RenderTextWithColor(fmt.Sprintf("  Target Item: %d", character.Task.TargetItemID), x, y, rl.NewColor(200, 200, 200, 255))
			y += int(lineHeight)
		}
	} else {
		renderer.RenderTextWithColor("  None", x, y, rl.NewColor(150, 150, 150, 255))
		y += int(lineHeight)
	}
	y += int(lineHeight)

	// Objectives
	renderer.RenderTextWithColor("Objectives:", x, y, rl.NewColor(255, 255, 255, 255))
	y += int(lineHeight)
	if len(character.Objectives) > 0 {
		for i, objective := range character.Objectives {
			objTypeStr := objectiveTypeToString(objective.Type)
			stuckStr := ""
			if objective.Stuck {
				stuckStr = " [STUCK]"
			}
			renderer.RenderTextWithColor(fmt.Sprintf("  %d. %s%s", i+1, objTypeStr, stuckStr), x, y, rl.NewColor(200, 200, 200, 255))
			y += int(lineHeight)
			if len(objective.Plan) > 0 {
				renderer.RenderTextWithColor(fmt.Sprintf("     Plan: %d tasks", len(objective.Plan)), x, y, rl.NewColor(150, 150, 150, 255))
				y += int(lineHeight)
			}
		}
	} else {
		renderer.RenderTextWithColor("  None", x, y, rl.NewColor(150, 150, 150, 255))
		y += int(lineHeight)
	}
	y += int(lineHeight)

	// Ambitions
	if len(character.Ambitions) > 0 {
		renderer.RenderTextWithColor("Ambitions:", x, y, rl.NewColor(255, 255, 255, 255))
		y += int(lineHeight)
		for i, ambition := range character.Ambitions {
			renderer.RenderTextWithColor(fmt.Sprintf("  %d. %s", i+1, ambition.Description), x, y, rl.NewColor(200, 200, 200, 255))
			y += int(lineHeight)
		}
		y += int(lineHeight)
	}

	// Path
	if len(character.Path) > 0 {
		renderer.RenderTextWithColor("Path:", x, y, rl.NewColor(255, 255, 255, 255))
		y += int(lineHeight)
		pathLength := len(character.Path)
		if pathLength > 5 {
			// Show first few and last few path points
			for i := 0; i < 3; i++ {
				renderer.RenderTextWithColor(fmt.Sprintf("  (%d, %d)", character.Path[i].X, character.Path[i].Y), x, y, rl.NewColor(200, 200, 200, 255))
				y += int(lineHeight)
			}
			renderer.RenderTextWithColor("  ...", x, y, rl.NewColor(150, 150, 150, 255))
			y += int(lineHeight)
			for i := pathLength - 2; i < pathLength; i++ {
				renderer.RenderTextWithColor(fmt.Sprintf("  (%d, %d)", character.Path[i].X, character.Path[i].Y), x, y, rl.NewColor(200, 200, 200, 255))
				y += int(lineHeight)
			}
			renderer.RenderTextWithColor(fmt.Sprintf("  Total: %d steps", pathLength), x, y, rl.NewColor(150, 150, 150, 255))
		} else {
			for _, pos := range character.Path {
				renderer.RenderTextWithColor(fmt.Sprintf("  (%d, %d)", pos.X, pos.Y), x, y, rl.NewColor(200, 200, 200, 255))
				y += int(lineHeight)
			}
		}
	}
}

// Helper function to convert TaskType to string
func taskTypeToString(taskType sim.TaskType) string {
	switch taskType {
	case sim.NoTaskType:
		return "None"
	case sim.Move:
		return "Move"
	case sim.Eat:
		return "Eat"
	case sim.Drink:
		return "Drink"
	case sim.Sleep:
		return "Sleep"
	case sim.PickUp:
		return "PickUp"
	case sim.Build:
		return "Build"
	case sim.Chop:
		return "Chop"
	default:
		return "Unknown"
	}
}

// Helper function to convert ObjectiveType to string
func objectiveTypeToString(objType sim.ObjectiveType) string {
	switch objType {
	case sim.NoObjective:
		return "None"
	case sim.DrinkObjective:
		return "Drink"
	case sim.EatObjective:
		return "Eat"
	case sim.SleepObjective:
		return "Sleep"
	case sim.BuildObjective:
		return "Build"
	default:
		return "Unknown"
	}
}
