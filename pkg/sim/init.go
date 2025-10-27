package sim

func InitSim() Sim {
	return Sim{
		Player:      InitPlayer(),
		Tiles:       InitRegion(),
		UI:          UIState{EditMode: false},
		Characters:  InitCharacters(),
		ItemManager: NewItemManager(),
	}
}
