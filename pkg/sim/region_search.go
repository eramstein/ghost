package sim

import "gociv/pkg/config"

func (sim *Sim) ScanForTile(position TilePosition, maxDistance int, terrain TileType) *TilePosition {
	// Check current tile
	tile := sim.GetTileAt(position)
	if tile.Type == terrain {
		return &position
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
			tile := sim.GetTileAt(TilePosition{X: x, Y: y})
			if tile.Type == terrain {
				return &TilePosition{X: x, Y: y}
			}
		}
		// bottom row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y + distance
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := sim.GetTileAt(TilePosition{X: x, Y: y})
			if tile.Type == terrain {
				return &TilePosition{X: x, Y: y}
			}
		}
		// left row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X - distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := sim.GetTileAt(TilePosition{X: x, Y: y})
			if tile.Type == terrain {
				return &TilePosition{X: x, Y: y}
			}
		}
		// right row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X + distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := sim.GetTileAt(TilePosition{X: x, Y: y})
			if tile.Type == terrain {
				return &TilePosition{X: x, Y: y}
			}
		}
	}
	return nil
}
