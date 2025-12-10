package sim

import (
	"fmt"
	"gociv/pkg/config"
)

// func (sim *Sim) FollowPath(character *Character, extraMove bool) {
// 	fmt.Printf("Following path for %v: %v\n", character.Name, character.Path)
// 	if character.Path == nil {
// 		return
// 	}
// 	task := character.CurrentTask
// 	path := character.Path
// 	if len(path) > 0 {
// 		nextTile := path[0]
// 		nextTileMoveCost := sim.GetTileAt(nextTile).MoveCost
// 		if nextTileMoveCost == -1 {
// 			return
// 		}
// 		if task.Progress >= float32(nextTileMoveCost) {
// 			character.TilePosition = nextTile
// 			character.Path = path[1:]
// 			task.Progress = task.Progress - float32(nextTileMoveCost)
// 			if len(character.Path) > 0 {
// 				sim.FollowPath(character, true)
// 			} else {
// 				sim.CompleteTask(character)
// 			}
// 		}
// 	}
// }

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
	if len(character.Path) == 0 || character.Path[len(character.Path)-1] != *target {
		path := sim.FindPath(character.TilePosition, *target, 0)
		if path == nil {
			fmt.Printf("No path found for %v to %v\n", character.Name, target)
			sim.CancelTask(character)
			return
		}
		character.Path = path
	}
	// sim.FollowPath(character, false)
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
	var direction WorldPosition
	if nextTileWorldPosition.X-c.WorldPosition.X > 0 {
		direction.X = 1
	} else {
		direction.X = -1
	}
	if nextTileWorldPosition.Y-c.WorldPosition.Y > 0 {
		direction.Y = 1
	} else {
		direction.Y = -1
	}

	c.WorldPosition.X += direction.X * CHARACTER_SPEED * deltaTime
	c.WorldPosition.Y += direction.Y * CHARACTER_SPEED * deltaTime

	dx := c.WorldPosition.X - nextTileWorldPosition.X
	dy := c.WorldPosition.Y - nextTileWorldPosition.Y
	distance := dx*dx + dy*dy

	if distance < (0.2*config.TileSize)*(0.2*config.TileSize) {
		if ((direction.X == 1 && c.WorldPosition.X >= nextTileWorldPosition.X) ||
			(direction.X == -1 && c.WorldPosition.X <= nextTileWorldPosition.X)) &&
			((direction.Y == 1 && c.WorldPosition.Y >= nextTileWorldPosition.Y) ||
				(direction.Y == -1 && c.WorldPosition.Y <= nextTileWorldPosition.Y)) {
			c.Path = c.Path[1:]
			if len(c.Path) == 0 {
				c.Path = nil
				sim.CompleteTask(c)
				return
			}
			c.TilePosition = nextTile
			fmt.Printf("Moved to %v\n", c.TilePosition)
		}
	}
}
