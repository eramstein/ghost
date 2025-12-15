package sim

type UIState struct {
	EditMode               bool
	Pause                  bool
	EditorTileType         TileType
	SelectedTileIndex      int
	SelectedCharacterIndex int
	SelectedPlantIndex     int
}
