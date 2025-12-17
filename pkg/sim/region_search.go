package sim

import "gociv/pkg/config"

// ScanForTile searches the closest reachable tile of a given terrain type using BFS
// Only explores passable tiles, so it respects walls and obstacles
func (sim *Sim) ScanForTile(position TilePosition, maxDistance int, terrain TileType) *TilePosition {
	// Check current tile first
	if position.X >= 0 && position.X < config.RegionSize && position.Y >= 0 && position.Y < config.RegionSize {
		tile := sim.GetTileAt(position)
		if tile.Type == terrain {
			return &position
		}
	}

	// BFS: explore tiles in order of distance, only through passable tiles
	visited := make(map[string]bool)
	queue := []TilePosition{position}
	visited[getPositionKey(position)] = true

	distance := 0
	// Process level by level to track distance
	levelStart := 0
	levelEnd := len(queue)

	for levelStart < len(queue) && distance < maxDistance {
		// Process all tiles at current distance level
		for i := levelStart; i < levelEnd; i++ {
			current := queue[i]

			// Check all neighbors
			for _, dir := range EightDirections {
				newX, newY := current.X+dir[0], current.Y+dir[1]

				// Check bounds
				if newX < 0 || newX >= config.RegionSize || newY < 0 || newY >= config.RegionSize {
					continue
				}

				// Check if already visited
				key := getPositionKey(TilePosition{X: newX, Y: newY})
				if visited[key] {
					continue
				}

				// Check if tile is passable
				tileIndex := newY*config.RegionSize + newX
				if sim.Tiles[tileIndex].MoveCost == ImpassableCost {
					// Mark as visited but don't explore further
					visited[key] = true
					continue
				}

				// Mark as visited and add to queue
				visited[key] = true
				neighborPos := TilePosition{X: newX, Y: newY}
				queue = append(queue, neighborPos)

				// Check if this tile matches the terrain type
				tile := sim.GetTileAt(neighborPos)
				if tile.Type == terrain {
					return &neighborPos
				}
			}
		}

		// Move to next distance level
		distance++
		levelStart = levelEnd
		levelEnd = len(queue)
	}

	return nil
}
