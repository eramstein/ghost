package sim

import (
	"fmt"
	"gociv/pkg/config"
	"math"
)

// EightDirections represents all 8-directional movement (orthogonal + diagonal)
// Used for pathfinding, BFS searches, and other tile exploration algorithms
var EightDirections = [][2]int{
	{0, 1},   // up
	{1, 1},   // up-right
	{1, 0},   // right
	{1, -1},  // down-right
	{0, -1},  // down
	{-1, -1}, // down-left
	{-1, 0},  // left
	{-1, 1},  // up-left
}

func (sim *Sim) advanceToNextTile(c *Character, nextTile TilePosition, nextTileWorldPosition WorldPosition) {
	c.WorldPosition = nextTileWorldPosition
	c.Path = c.Path[1:]
	c.TilePosition = nextTile
}

func (sim *Sim) MoveForTask(character *Character) {
	task := character.CurrentTask
	if task == nil {
		return
	}
	target := task.TargetTile
	if target == nil {
		fmt.Printf("Task target is not a *TilePosition: %v\n", task.TargetTile)
		return
	}

	// are we there yet?
	if character.TilePosition.X == target.X && character.TilePosition.Y == target.Y {
		sim.CompleteTask(character)
		return
	}

	// if the character has not set its path yet, or is headed in the wrong direction, find a path to the target
	if len(character.Path) == 0 || character.Path[len(character.Path)-1] != *target {
		path := sim.FindPath(character.TilePosition, *target, 0)
		if path == nil {
			fmt.Printf("No path found for %v to %v\n", character.Name, target)
			sim.CancelTask(character)
			return
		}
		character.Path = path
	}
}

func (sim *Sim) Move(c *Character, deltaTime float32) {
	if len(c.Path) == 0 {
		return
	}
	nextTile := c.Path[0]
	nextTileWorldPosition := WorldPosition{
		X: float32(nextTile.X*config.TileSize + config.TileSize/2),
		Y: float32(nextTile.Y*config.TileSize + config.TileSize/2),
	}

	// Calculate direction vector from current position to target
	dx := nextTileWorldPosition.X - c.WorldPosition.X
	dy := nextTileWorldPosition.Y - c.WorldPosition.Y
	distance := dx*dx + dy*dy

	// If we're already very close, snap to target and move to next tile
	if distance < (0.1*config.TileSize)*(0.1*config.TileSize) {
		sim.advanceToNextTile(c, nextTile, nextTileWorldPosition)
		return
	}

	// Calculate distance and normalize direction vector
	distanceSqrt := float32(math.Sqrt(float64(distance)))
	if distanceSqrt <= 0 {
		return
	}

	// Normalize direction
	direction := WorldPosition{
		X: dx / distanceSqrt,
		Y: dy / distanceSqrt,
	}

	// Calculate movement this frame
	moveDistance := CHARACTER_SPEED * deltaTime
	remainingDistance := distanceSqrt

	// If we would overshoot, snap to target instead
	if moveDistance >= remainingDistance {
		sim.advanceToNextTile(c, nextTile, nextTileWorldPosition)
	} else {
		// Move towards target
		c.WorldPosition.X += direction.X * moveDistance
		c.WorldPosition.Y += direction.Y * moveDistance
	}
}
