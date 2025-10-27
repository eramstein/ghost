package sim

type TileType int

const (
	TileTypeEmpty TileType = iota
	TileTypeWall
	TileTypeFloor
	TileTypeDirt
)

func (tt TileType) String() string {
	switch tt {
	case TileTypeEmpty:
		return "Empty"
	case TileTypeWall:
		return "Wall"
	case TileTypeFloor:
		return "Floor"
	case TileTypeDirt:
		return "Dirt"
	default:
		return "Unknown"
	}
}

type MoveCost float64

const (
	DefaultMoveCost   MoveCost = 1.0  // Normal movement cost
	DifficultMoveCost MoveCost = 2.0  // Increased cost for difficult terrain
	ImpassableCost    MoveCost = -1.0 // Represents an impassable tile
)

func (t *Tile) UpdateType(newType TileType) {
	t.Type = newType
	switch newType {
	case TileTypeWall:
		t.MoveCost = ImpassableCost
	case TileTypeEmpty, TileTypeFloor, TileTypeDirt:
		t.MoveCost = DefaultMoveCost
	default:
		t.MoveCost = DefaultMoveCost
	}
}
