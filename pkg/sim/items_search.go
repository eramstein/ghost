package sim

import (
	"fmt"
	"gociv/pkg/config"
)

// ScanForItem searches the closest item of a given type by looping tiles around a position
// if variant is irrelevant pass -1
func (sim *Sim) ScanForItem(position TilePosition, maxDistance int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	// Check current tile
	if position.X >= 0 && position.X < config.RegionSize && position.Y >= 0 && position.Y < config.RegionSize {
		tile := sim.GetTileAt(position)
		for _, itemID := range tile.Items {
			item := sim.GetItem(itemID)
			if item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1) {
				return &item
			}
		}
	}

	// Check tiles further and further away until maxDistance
	for distance := 1; distance <= maxDistance; distance++ {
		// top row: y = position.Y - distance, must be in bounds
		y := position.Y - distance
		if y >= 0 && y < config.RegionSize {
			// x = position.X + dx, must be in bounds: 0 <= position.X + dx < config.RegionSize
			dxMin := max(-distance, -position.X)
			dxMax := min(distance, config.RegionSize-1-position.X)
			for dx := dxMin; dx <= dxMax; dx++ {
				x := position.X + dx
				item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
				if item != nil {
					return item
				}
			}
		}
		// bottom row: y = position.Y + distance, must be in bounds
		y = position.Y + distance
		if y >= 0 && y < config.RegionSize {
			// x = position.X + dx, must be in bounds: 0 <= position.X + dx < config.RegionSize
			dxMin := max(-distance, -position.X)
			dxMax := min(distance, config.RegionSize-1-position.X)
			for dx := dxMin; dx <= dxMax; dx++ {
				x := position.X + dx
				item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
				if item != nil {
					return item
				}
			}
		}
		// left row: x = position.X - distance, must be in bounds
		x := position.X - distance
		if x >= 0 && x < config.RegionSize {
			// y = position.Y + dy, must be in bounds: 0 <= position.Y + dy < config.RegionSize
			dyMin := max(-distance, -position.Y)
			dyMax := min(distance, config.RegionSize-1-position.Y)
			for dy := dyMin; dy <= dyMax; dy++ {
				y := position.Y + dy
				item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
				if item != nil {
					return item
				}
			}
		}
		// right row: x = position.X + distance, must be in bounds
		x = position.X + distance
		if x >= 0 && x < config.RegionSize {
			// y = position.Y + dy, must be in bounds: 0 <= position.Y + dy < config.RegionSize
			dyMin := max(-distance, -position.Y)
			dyMax := min(distance, config.RegionSize-1-position.Y)
			for dy := dyMin; dy <= dyMax; dy++ {
				y := position.Y + dy
				item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
				if item != nil {
					return item
				}
			}
		}
	}
	return nil
}

func (sim *Sim) FindItemInTile(x int, y int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	if x == 1 && y == 1 {
		fmt.Printf("Finding item in tile %d, %d\n", x, y)
	}
	tile := sim.GetTileAt(TilePosition{X: x, Y: y})
	for _, itemID := range tile.Items {
		item := sim.GetItem(itemID)
		if item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1) {
			fmt.Printf("Found item %v ID %d %d\n", item, item.ID, itemID)
			return &item
		}
	}
	return nil
}

func (sim *Sim) FindInInventory(character *Character, itemType ItemType, variant int) *Item {
	for _, itemID := range character.Inventory {
		item := sim.GetItem(itemID)
		if item.Type == itemType && (item.Variant == variant || variant == -1) {
			return &item
		}
	}
	return nil
}
