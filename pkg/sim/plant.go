package sim

import (
	"gociv/pkg/data"
)

type PlantType int

const (
	PlantTypeTree PlantType = iota
	Tree
)

func (sim *Sim) UpdatePlants() {
	if sim.PlantManager == nil {
		return
	}
	sim.PlantManager.ForEach(func(id int, p *Plant) {
		sim.Update(p)
	})
}

func (sim *Sim) SpawnPlant(position TilePosition, variant int, plantType PlantType) int {
	plant, _ := data.GetPlantDefinition(int(plantType), variant)
	newPlant := Plant{
		Variant:    variant,
		Position:   position,
		PlantType:  plantType,
		GrowthRate: plant.GrowthRate,
		Produces: Production{
			Type:            ItemType(plant.Produces.Type),
			Variant:         plant.Produces.Variant,
			ProductionStage: 99,
			ProductionRate:  plant.Produces.ProductionRate,
		},
	}
	return sim.AddPlant(newPlant)
}

func (sim *Sim) Update(plant *Plant) {
	if plant.GrowthStage < 100 {
		plant.GrowthStage += plant.GrowthRate
	}
	if plant.GrowthStage >= 100 && plant.Produces.ProductionStage <= 100 {
		plant.Produces.ProductionStage += plant.Produces.ProductionRate
	}
	if plant.Produces.ProductionStage >= 100 {
		plant.Produces.ProductionStage = 0
		sim.AddItem(
			Item{
				Type:       plant.Produces.Type,
				Variant:    plant.Produces.Variant,
				Durability: 100,
				Efficiency: 100,
			},
			ItemLocation{LocationType: LocTile, TilePosition: plant.Position})
	}
}

func (sim *Sim) RemovePlantById(id int) {
	sim.RemovePlant(id)
}
