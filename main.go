package main

import (
	"fmt"
	"gociv/pkg/data"
	"gociv/pkg/input"
	"gociv/pkg/render"
	"gociv/pkg/sim"
	"gociv/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var tickTime = float32(0.0)

func main() {
	rl.SetTraceLogLevel(rl.LogWarning)

	// Enable 4x MSAA anti-aliasing for smoother graphics
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(1600, 900, "Ghost")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Load all game data
	if err := data.LoadAllData(); err != nil {
		fmt.Printf("Error loading game data: %v\n", err)
		panic(err)
	}

	// Try to load quicksave first, fallback to InitSim if it fails
	var simData *sim.Sim
	loadedSim, err := utils.LoadSim("quicksave")
	if err != nil {
		fmt.Printf("No quicksave found or error loading: %v\n", err)
		fmt.Println("Starting new game...")
		simData = sim.InitSim()
	} else {
		fmt.Println("Loaded quicksave successfully!")
		simData = loadedSim
	}

	renderer := render.NewRenderer(simData, nil)
	inputManager := input.NewManager(simData, &renderer.Camera)
	renderer.Console = inputManager.GetConsole()
	if err := renderer.FontManager.LoadFont("default", "assets/fonts/Roboto-Regular.ttf", 16); err != nil {
		fmt.Printf("failed to load default font: %v\n", err)
	}
	renderer.DefaultFont = renderer.FontManager.GetFont("default")

	defer renderer.Close()

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()
		tickTime += deltaTime

		// Handle input events (keyboard, mouse)
		inputManager.HandleInputs(deltaTime)

		if !simData.UI.Pause {
			// Frame updates (things needed to be done every frame)
			simData.FrameUpdate(deltaTime)
			// Logic updates (long term simulation)
			if tickTime >= sim.SIM_STEP {
				simData.LogicUpdate()
				tickTime = 0.0
			}
		}

		// Draw
		rl.BeginDrawing()
		renderer.Render(simData)

		// Draw white background for FPS
		rl.DrawFPS(700, 10)
		rl.EndDrawing()
	}
}
