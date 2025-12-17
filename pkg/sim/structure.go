package sim

type StructureType int

const (
	Extractor StructureType = iota
	Furniture
	Workshop
	Storage
)

func (sim *Sim) InitStructures() {
	sim.SpawnStructure(TilePosition{X: 2, Y: 2}, 0, Extractor)
}

func (sim *Sim) SpawnStructure(position TilePosition, variant int, structureType StructureType) {
	//structure, _ := data.GetStructureDefinition(int(structureType), variant)
	newStructure := Structure{
		Variant:       variant,
		Position:      position,
		StructureType: structureType,
		Condition:     100,
		Owner:         -1,
	}
	sim.StructureManager.AddStructure(newStructure)
}
