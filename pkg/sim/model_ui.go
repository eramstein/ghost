package sim

type EditorMode int

const (
	EditorModeTiles EditorMode = iota
	EditorModePlants
	EditorModeStructures
)

func (em EditorMode) String() string {
	switch em {
	case EditorModeTiles:
		return "Tiles"
	case EditorModePlants:
		return "Plants"
	case EditorModeStructures:
		return "Structures"
	default:
		return "Unknown"
	}
}

type UIState struct {
	EditMode               bool
	Pause                  bool
	EditorMode             EditorMode
	EditorTileType         TileType
	EditorPlantType        PlantType
	EditorPlantVariant     int16
	EditorStructureType    StructureType
	SelectedTileIndex      int
	SelectedCharacterIndex int
	SelectedPlantIndex     int
	SelectedStructureIndex int
}
