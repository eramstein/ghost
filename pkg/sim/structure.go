package sim

type StructureType int

const (
	Well StructureType = iota
	Bed
	Furniture
	Workshop
	Storage
)

func (sim *Sim) SpawnStructure(position TilePosition, structureType StructureType) {
	newStructure := Structure{
		Position:      position,
		StructureType: structureType,
		Condition:     100,
		Owner:         -1,
	}
	sim.AddStructure(newStructure)
}
