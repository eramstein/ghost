package render

import (
	"fmt"
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DisplayTime shows the current time
func DisplayTime(r *Renderer, calendar *sim.Calendar) {
	timeText := fmt.Sprintf("Day %d, Hour %d, Minute %d", calendar.Day, calendar.Hour, calendar.Minute)

	// Draw white background
	rl.DrawRectangle(8, 8, 185, 24, rl.White)

	// Draw text
	r.RenderTextWithColor(timeText, 20, 13, rl.Black)
}
