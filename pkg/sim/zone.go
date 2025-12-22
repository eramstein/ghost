package sim

type ZoneType int

const (
	ZoneTypeNone ZoneType = iota
	ZoneTypeField
	ZoneTypeRoom
)

type Zone interface {
	GetCentroid() TilePosition
	GetTiles() []TilePosition
}

func GetZoneCentroid(tiles []TilePosition) TilePosition {
	var centroid TilePosition
	for _, tile := range tiles {
		centroid.X += tile.X
		centroid.Y += tile.Y
	}
	centroid.X /= len(tiles)
	centroid.Y /= len(tiles)
	return centroid
}

func GetZoneTileIndex(zone Zone, position TilePosition) int {
	for i, tile := range zone.GetTiles() {
		if tile.IsSameAs(position) {
			return i
		}
	}
	return -1
}
