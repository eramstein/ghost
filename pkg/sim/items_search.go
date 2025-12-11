package sim

import "gociv/pkg/config"

// ScanForItem searches the closest item of a given type by looping tiles around a position
// if variant is irrelevant pass -1
func (sim *Sim) ScanForItem(position TilePosition, maxDistance int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	// Check current tile
	if position.X >= 0 && position.X < config.RegionSize && position.Y >= 0 && position.Y < config.RegionSize {
		tile := sim.GetTileAt(position)
		for _, itemID := range tile.Items {
			item := sim.ItemManager.GetItem(itemID)
			if item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1) {
				return &item
			}
		}
	}

	// Check tiles further and further away until maxDistance
	for distance := 1; distance <= maxDistance; distance++ {
		// top row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y - distance
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// bottom row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y + distance
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// left row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X - distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// right row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X + distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			item := sim.FindItemInTile(x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
	}
	return nil
}

func (sim *Sim) FindItemInTile(x int, y int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
		return nil
	}
	tile := sim.GetTileAt(TilePosition{X: x, Y: y})
	for _, itemID := range tile.Items {
		item := sim.ItemManager.GetItem(itemID)
		if item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == -1) {
			return &item
		}
	}
	return nil
}

func (sim *Sim) FindInInventory(character *Character, itemType ItemType, variant int) *Item {
	for _, itemID := range character.Inventory {
		item := sim.ItemManager.GetItem(itemID)
		if item.Type == itemType && (item.Variant == variant || variant == -1) {
			return &item
		}
	}
	return nil
}
