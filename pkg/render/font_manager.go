package render

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// FontManager handles font loading and management
type FontManager struct {
	fonts map[string]rl.Font
}

// NewFontManager creates a new font manager
func NewFontManager() *FontManager {
	return &FontManager{
		fonts: make(map[string]rl.Font),
	}
}

// LoadFont loads a font from a TTF/OTF file
func (fm *FontManager) LoadFont(name, path string, fontSize int32) error {
	font := rl.LoadFontEx(path, fontSize, nil, 0)
	if font.BaseSize == 0 {
		return fmt.Errorf("failed to load font: %s", path)
	}
	fm.fonts[name] = font
	return nil
}

// GetFont returns a loaded font by name
func (fm *FontManager) GetFont(name string) rl.Font {
	font, exists := fm.fonts[name]
	if !exists {
		// Fallback to default raylib font if custom font not found
		font = rl.GetFontDefault()
	}
	return font
}

// UnloadAll unloads all loaded fonts
func (fm *FontManager) UnloadAll() {
	for _, font := range fm.fonts {
		rl.UnloadFont(font)
	}
	fm.fonts = make(map[string]rl.Font)
}
