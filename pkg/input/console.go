package input

import (
	"fmt"
	"gociv/pkg/sim"
	"gociv/pkg/utils"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Console struct {
	isOpen         bool
	inputBuffer    string
	commandHistory []string
	historyIndex   int
	sim            *sim.Sim
}

// NewConsole creates a new console instance
func NewConsole(sim *sim.Sim) *Console {
	return &Console{
		isOpen:         false,
		inputBuffer:    "",
		commandHistory: make([]string, 0),
		historyIndex:   -1,
		sim:            sim,
	}
}

// IsOpen returns whether the console is currently open
func (c *Console) IsOpen() bool {
	return c.isOpen
}

// Toggle toggles the console open/closed state
func (c *Console) Toggle() {
	c.isOpen = !c.isOpen
	if c.isOpen {
		c.inputBuffer = ""
		c.historyIndex = -1
	}
}

// HandleInput processes console input when the console is open
func (c *Console) HandleInput() {
	if !c.isOpen {
		return
	}

	// Handle text input using Raylib's simple approach
	key := rl.GetCharPressed()
	if key != 0 {
		// Only allow printable characters
		if key >= 32 && key <= 126 && len(c.inputBuffer) < 256 {
			c.inputBuffer += string(rune(key))
		}
	}

	// Handle backspace
	if rl.IsKeyPressed(rl.KeyBackspace) && len(c.inputBuffer) > 0 {
		c.inputBuffer = c.inputBuffer[:len(c.inputBuffer)-1]
	}

	// Handle Enter to execute command
	if rl.IsKeyPressed(rl.KeyEnter) {
		c.executeCommand(c.inputBuffer)
		c.addToHistory(c.inputBuffer)
		c.inputBuffer = ""
		c.historyIndex = -1
	}

	// Handle Escape to close console
	if rl.IsKeyPressed(rl.KeyEscape) {
		c.isOpen = false
		c.inputBuffer = ""
		c.historyIndex = -1
	}

	// Handle command history navigation
	if rl.IsKeyPressed(rl.KeyUp) {
		c.navigateHistory(-1)
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		c.navigateHistory(1)
	}
}

// executeCommand processes and executes console commands
func (c *Console) executeCommand(command string) {
	command = strings.TrimSpace(command)
	if command == "" {
		return
	}

	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	switch cmd {
	case "clear":
		// Clear console output (this would need to be implemented in the renderer)
		fmt.Println("Console cleared")
	case "echo":
		if len(args) > 0 {
			fmt.Println(strings.Join(args, " "))
		} else {
			fmt.Println("Usage: echo <message>")
		}
	case "save-tiles":
		c.handleSaveCommand(args)
	case "load-tiles":
		c.handleLoadCommand(args)
	default:
		fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", cmd)
	}
}

// handleSaveCommand handles save-related commands
func (c *Console) handleSaveCommand(args []string) {
	filename := "tiles.gob"
	if len(args) > 0 {
		filename = args[0]
	}

	err := utils.SaveRegion(c.sim, filename)
	if err != nil {
		fmt.Printf("Error saving region: %v\n", err)
	} else {
		fmt.Printf("Region saved successfully to %s!\n", filename)
	}
}

// handleLoadCommand handles load-related commands
func (c *Console) handleLoadCommand(args []string) {
	filename := "tiles.gob"
	if len(args) > 0 {
		filename = args[0]
	}

	regionData, err := sim.LoadRegion(filename)
	if err != nil {
		fmt.Printf("Error loading region: %v\n", err)
	} else {
		// Replace current sim region data with loaded data
		c.sim.Tiles = regionData.Tiles
		if regionData.PlantManager != nil {
			c.sim.PlantManager = regionData.PlantManager
		}
		if regionData.StructureManager != nil {
			c.sim.StructureManager = regionData.StructureManager
		}
		fmt.Printf("Region loaded successfully from %s!\n", filename)
	}
}

// addToHistory adds a command to the history
func (c *Console) addToHistory(command string) {
	if command == "" {
		return
	}

	// Don't add duplicate consecutive commands
	if len(c.commandHistory) == 0 || c.commandHistory[len(c.commandHistory)-1] != command {
		c.commandHistory = append(c.commandHistory, command)
	}

	// Limit history size
	if len(c.commandHistory) > 50 {
		c.commandHistory = c.commandHistory[1:]
	}
}

// navigateHistory navigates through command history
func (c *Console) navigateHistory(direction int) {
	if len(c.commandHistory) == 0 {
		return
	}

	c.historyIndex += direction

	if c.historyIndex < 0 {
		c.historyIndex = -1
		c.inputBuffer = ""
		return
	}

	if c.historyIndex >= len(c.commandHistory) {
		c.historyIndex = len(c.commandHistory) - 1
	}

	if c.historyIndex >= 0 && c.historyIndex < len(c.commandHistory) {
		c.inputBuffer = c.commandHistory[c.historyIndex]
	}
}

// GetInputBuffer returns the current input buffer
func (c *Console) GetInputBuffer() string {
	return c.inputBuffer
}

// GetHistory returns the command history
func (c *Console) GetHistory() []string {
	return c.commandHistory
}
