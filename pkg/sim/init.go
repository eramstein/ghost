package sim

func InitSim() *Sim {
	sim := Sim{
		Player:      InitPlayer(),
		Tiles:       InitRegion(),
		UI:          UIState{EditMode: false, SelectedCharacterIndex: -1, SelectedPlantIndex: -1, SelectedTileIndex: -1},
		ItemManager: NewItemManager(),
	}
	sim.InitItems()
	sim.InitCharacters()
	sim.InitPlants()
	return &sim
}
