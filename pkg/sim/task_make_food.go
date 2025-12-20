package sim

func (sim *Sim) GetNextMakingFoodTask(character *Character, objective *Objective) (task *Task) {
	// for now we can only make food by planting seeds
	// it requires both a growable tile and a seed
	var newTask *Task

	// is there already something growing? if yes stuck objective

	// find

	// get closest field, create a one if needed
	field := sim.GetClosestField(character.TilePosition)
	if field == nil {
		// TODO: look for suitable spot anf fail objective if not found
		field = sim.CreateField([]TilePosition{character.TilePosition}, 0)
	}

	// create task to plant seeds
	//freeTiles := field.GetFreeTiles()

	return newTask
}
