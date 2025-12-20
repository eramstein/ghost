package sim

type Zone interface {
	Centroid() TilePosition
	Tiles() []TilePosition
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
