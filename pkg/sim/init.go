package sim

func InitSim() *Sim {
	sim := Sim{
		Player:      InitPlayer(),
		Tiles:       InitRegion(),
		UI:          UIState{EditMode: false},
		Characters:  InitCharacters(),
		ItemManager: NewItemManager(),
	}
	sim.InitItems()
	return &sim
}
