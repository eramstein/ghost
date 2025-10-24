package render

import (
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawPlayer(renderer *Renderer, player sim.Player) {
	rl.DrawCircleV(
		rl.Vector2{X: player.WorldPosition.X, Y: player.WorldPosition.Y},
		10,
		rl.White,
	)
}
