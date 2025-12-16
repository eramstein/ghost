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

// DrawCharacterDetails renders character info starting at (x, y) and returns
// the updated y position after drawing.
func DrawCharacterDetails(renderer *Renderer, character *sim.Character, x, y int) int {
	if character == nil {
		return y
	}

	// Text settings
	lineHeight := int32(renderer.DefaultFont.BaseSize + 6)

	// Title - Character Name
	titleText := character.Name
	renderer.RenderTextWithColor(titleText, x, y, rl.NewColor(255, 255, 255, 255))
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
	if character.CurrentTask != nil && character.CurrentTask.Type != sim.Move && character.CurrentTask.Type != sim.NoTaskType {
		taskTypeStr := character.CurrentTask.Type.String()
		renderer.RenderTextWithColor(fmt.Sprintf("  Type: %s", taskTypeStr), x, y, rl.NewColor(200, 200, 200, 255))
		y += int(lineHeight)
		renderer.RenderTextWithColor(fmt.Sprintf("  Progress: %.1f%%", character.CurrentTask.Progress), x, y, rl.NewColor(200, 200, 200, 255))
		y += int(lineHeight)
		if character.CurrentTask.TargetTile != nil {
			renderer.RenderTextWithColor(fmt.Sprintf("  Target: (%d, %d)", character.CurrentTask.TargetTile.X, character.CurrentTask.TargetTile.Y), x, y, rl.NewColor(200, 200, 200, 255))
			y += int(lineHeight)
		}
		if character.CurrentTask.TargetItem != nil {
			renderer.RenderTextWithColor(fmt.Sprintf("  Target Item: %d", character.CurrentTask.TargetItem.Type), x, y, rl.NewColor(200, 200, 200, 255))
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
			objTypeStr := objective.Type.String()
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

	return y
}
