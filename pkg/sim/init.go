package sim

func InitSim() *Sim {
	sim := Sim{
		Player:           InitPlayer(),
		Tiles:            InitRegion(),
		UI:               UIState{EditMode: false, SelectedCharacterIndex: -1, SelectedPlantIndex: -1, SelectedTileIndex: -1, SelectedStructureIndex: -1},
		ItemManager:      NewItemManager(),
		PlantManager:     NewPlantManager(),
		StructureManager: NewStructureManager(),
	}
	sim.InitItems()
	sim.InitCharacters()
	sim.InitPlants()
	sim.InitStructures()
	return &sim
}
