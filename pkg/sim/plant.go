package sim

import (
	"fmt"
	"gociv/pkg/data"
)

type PlantType int

const (
	PlantTypeTree PlantType = iota
	Tree
)

func (sim *Sim) InitPlants() {
	// apple tree
	sim.SpawnPlant(TilePosition{X: 1, Y: 3}, 0, PlantTypeTree)
}

func (sim *Sim) UpdatePlants() {
	if sim.PlantManager == nil {
		return
	}
	sim.PlantManager.ForEach(func(id int, p *Plant) {
		sim.Update(p)
	})
}

func (sim *Sim) SpawnPlant(position TilePosition, variant int, plantType PlantType) {
	plant, _ := data.GetPlantDefinition(int(plantType), variant)
	newPlant := Plant{
		Position:   position,
		PlantType:  plantType,
		GrowthRate: plant.GrowthRate,
		Produces: PlantProduction{
			Type:            ItemType(plant.Produces.Type),
			Variant:         plant.Produces.Variant,
			ProductionStage: 99,
			ProductionRate:  plant.Produces.ProductionRate,
		},
	}
	sim.PlantManager.addPlant(newPlant)
}

func (sim *Sim) Update(plant *Plant) {
	if plant.GrowthStage < 100 {
		plant.GrowthStage += plant.GrowthRate
		fmt.Printf("Plant growth: %d/%d\n", plant.GrowthRate, plant.GrowthStage)
	}
	if plant.GrowthStage >= 100 && plant.Produces.ProductionStage <= 100 {
		plant.Produces.ProductionStage += plant.Produces.ProductionRate
	}
	if plant.Produces.ProductionStage >= 100 {
		fmt.Printf("Plant produced: %v\n", plant.Produces)
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
	sim.PlantManager.removePlant(id)
}
