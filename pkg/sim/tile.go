package sim

import (
	"fmt"
	"math"
)

type TileType int

const (
	TileTypeEmpty TileType = iota
	TileTypeWall
	TileTypeFloor
	TileTypeDirt
	TileTypeWater
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
	case TileTypeWater:
		return "Water"
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

func (t *Tile) AddItem(itemID int) {
	t.Items = append(t.Items, itemID)
}

func (t *Tile) RemoveItem(itemID int) {
	fmt.Printf("Removing item %d from tile %v with items %v\n", itemID, t.Position, t.Items)
	for i, item := range t.Items {
		if item == itemID {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}
	}
}

func IsAdjacent(x1, y1, x2, y2 int) bool {
	return math.Abs(float64(x1-x2)) <= 1 && math.Abs(float64(y1-y2)) <= 1
}
