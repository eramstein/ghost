package render

import (
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// Main game color theme
	ColorBackground = rl.Color{R: 17, G: 21, B: 28, A: 255}
	ColorBorder     = rl.Color{R: 25, G: 32, B: 36, A: 255}
	ColorWall       = rl.Color{R: 33, G: 45, B: 64, A: 255}
	ColorFloor      = rl.Color{R: 54, G: 65, B: 86, A: 255}
	ColorDirt       = rl.Color{R: 125, G: 78, B: 87, A: 255}
	ColorLife       = rl.Color{R: 214, G: 104, B: 83, A: 255}

	// Debug or Editor colors
	ColorEditMode = rl.Color{R: 255, G: 255, B: 0, A: 255} // Yellow for EditMode indicator
	ColorPath     = rl.Color{R: 255, G: 165, B: 0, A: 180} // Orange with transparency for path highlighting
)

// TileTypeColors maps each TileType to its corresponding color
var TileTypeColors = map[sim.TileType]rl.Color{
	sim.TileTypeEmpty: ColorBackground,
	sim.TileTypeWall:  ColorWall,
	sim.TileTypeFloor: ColorFloor,
	sim.TileTypeDirt:  ColorDirt,
}

var ItemTypeColors = map[sim.ItemType]rl.Color{
	sim.ItemTypeFood: ColorLife,
}
