package sim

type StructureType int

const (
	Well StructureType = iota
	Furniture
	Workshop
	Storage
)

func (sim *Sim) InitStructures() {
	sim.SpawnStructure(TilePosition{X: 21, Y: 15}, 0, Well)
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
	sim.AddStructure(newStructure)
}
