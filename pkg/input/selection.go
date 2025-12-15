package input

func (m *Manager) SelectCharacter(characterID int) {
	m.sim.UI.SelectedCharacterIndex = characterID
}

func (m *Manager) SelectTile(tileID int) {
	m.sim.UI.SelectedTileIndex = tileID
}

func (m *Manager) SelectPlant(plantID int) {
	plantIndex := -1
	for i, plant := range m.sim.Plants {
		if plant.ID == plantID {
			plantIndex = i
			break
		}
	}
	m.sim.UI.SelectedPlantIndex = plantIndex
}

func (m *Manager) ClearSelections() {
	m.sim.UI.SelectedTileIndex = -1
	m.sim.UI.SelectedCharacterIndex = -1
	m.sim.UI.SelectedPlantIndex = -1
}
