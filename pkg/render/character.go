package render

import (
	"gociv/pkg/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCharacters(renderer *Renderer, characters []sim.Character) {
	for _, character := range characters {
		rl.DrawCircleV(
			rl.Vector2{X: character.WorldPosition.X, Y: character.WorldPosition.Y},
			10,
			ColorLife,
		)
	}
}
