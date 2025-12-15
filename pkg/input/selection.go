package input

func (m *Manager) SelectCharacter(characterID int) {
	m.sim.UI.SelectedCharacterIndex = characterID
}

func (m *Manager) SelectTile(tileID int) {
	m.sim.UI.SelectedTileIndex = tileID
}

// SelectPlant now takes a plant ID (managed by PlantManager) rather than a slice index.
func (m *Manager) SelectPlant(plantID int) {
	m.sim.UI.SelectedPlantIndex = plantID
}

func (m *Manager) ClearSelections() {
	m.sim.UI.SelectedTileIndex = -1
	m.sim.UI.SelectedCharacterIndex = -1
	m.sim.UI.SelectedPlantIndex = -1
}
