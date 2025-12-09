package sim

func InitSim() *Sim {
	sim := Sim{
		Player:      InitPlayer(),
		Tiles:       InitRegion(),
		UI:          UIState{EditMode: false},
		ItemManager: NewItemManager(),
	}
	sim.InitItems()
	sim.InitCharacters()
	return &sim
}
