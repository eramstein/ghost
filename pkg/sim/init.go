package sim

func InitSim() *Sim {
	// region includes plants and structures
	regionData := InitRegion()

	sim := Sim{
		Player:           InitPlayer(),
		Tiles:            regionData.Tiles,
		UI:               UIState{EditMode: false, EditorMode: EditorModeTiles, EditorTileType: TileTypeEmpty, EditorPlantType: PlantTypeTree, EditorPlantVariant: 0, EditorStructureType: Well, SelectedCharacterIndex: -1, SelectedPlantIndex: -1, SelectedTileIndex: -1, SelectedStructureIndex: -1},
		ItemManager:      NewItemManager(),
		PlantManager:     regionData.PlantManager,
		StructureManager: regionData.StructureManager,
	}

	sim.InitItems()
	sim.InitCharacters()
	return &sim
}
