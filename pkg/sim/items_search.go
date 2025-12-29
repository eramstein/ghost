package sim

import (
	"fmt"
	"gociv/pkg/config"
)

// ScanForItem searches the closest reachable item of a given type using BFS
// Only explores passable tiles, so it respects walls and obstacles
// if variant is irrelevant pass -1
func (sim *Sim) ScanForItem(characterID int, position TilePosition, maxDistance int, itemType ItemType, variant int16, unclaimedOnly bool) *Item {
	// Check current tile first
	if position.X >= 0 && position.X < config.RegionSize && position.Y >= 0 && position.Y < config.RegionSize {
		tile := sim.GetTileAt(position)
		for _, itemID := range tile.Items {
			item := sim.GetItemPtr(itemID)
			if item != nil && item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1 || item.OwnedBy == characterID) {
				return item
			}
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

	for levelStart < len(queue) && (distance < maxDistance || maxDistance == -1) {
		// Process all tiles at current distance level
		for i := levelStart; i < levelEnd; i++ {
			current := queue[i]

			// Check all neighbors
			for _, dir := range EightDirections {
				newX, newY := current.X+int16(dir[0]), current.Y+int16(dir[1])

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

				// Check for items in this tile
				item := sim.FindItemInTile(characterID, neighborPos, itemType, variant, unclaimedOnly)
				if item != nil {
					return item
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

func (sim *Sim) FindItemInTile(characterID int, position TilePosition, itemType ItemType, variant int16, unclaimedOnly bool) *Item {
	if position.X == 1 && position.Y == 1 {
		fmt.Printf("Finding item in tile %d, %d\n", position.X, position.Y)
	}
	tile := sim.GetTileAt(position)
	for _, itemID := range tile.Items {
		item := sim.GetItemPtr(itemID)
		if item != nil && item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1 || item.OwnedBy == characterID) {
			fmt.Printf("Found item %v ID %d %d for character %d\n", item, item.ID, itemID, characterID)
			return item
		}
	}
	return nil
}

func (sim *Sim) FindInInventory(character *Character, itemType ItemType, variant int16) *Item {
	for _, itemID := range character.Inventory {
		item := sim.GetItemPtr(itemID)
		if item != nil && item.Type == itemType && (item.Variant == variant || variant == -1) {
			return item
		}
	}
	return nil
}
