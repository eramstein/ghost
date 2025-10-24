package render

import (
	"gociv/pkg/input"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawConsole renders the console overlay at the bottom
func DrawConsole(console *input.Console) {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	// Console height (single line)
	consoleHeight := int32(40)
	consoleY := int32(screenHeight) - consoleHeight

	// Console background (semi-transparent black overlay)
	rl.DrawRectangle(0, consoleY, int32(screenWidth), consoleHeight, rl.NewColor(0, 0, 0, 200))

	// Console border
	rl.DrawRectangleLines(0, consoleY, int32(screenWidth), consoleHeight, rl.NewColor(100, 100, 100, 255))

	// Console prompt
	prompt := "> "
	inputText := console.GetInputBuffer()
	fullText := prompt + inputText

	// Draw the input text
	rl.DrawText(fullText, 10, consoleY+10, 20, rl.NewColor(255, 255, 255, 255))

	// Draw blinking cursor
	if int(rl.GetTime()*2)%2 == 0 {
		cursorX := 10 + rl.MeasureText(fullText, 20)
		rl.DrawText("_", cursorX, consoleY+10, 20, rl.NewColor(255, 255, 255, 255))
	}
}
