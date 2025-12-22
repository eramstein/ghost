package sim

import (
	"fmt"
	"gociv/pkg/config"
)

type ObjectiveType int

// Sorted by priority (lowest number is highest priority)
const (
	NoObjective ObjectiveType = iota
	DrinkObjective
	EatObjective
	SleepObjective
	MakeFoodObjective
	BuildObjective
)

func (ot ObjectiveType) String() string {
	switch ot {
	case NoObjective:
		return "No Objective"
	case DrinkObjective:
		return "Drink"
	case EatObjective:
		return "Eat"
	case SleepObjective:
		return "Sleep"
	case MakeFoodObjective:
		return "Make Food"
	case BuildObjective:
		return "Build"
	}
	return "Unknown"
}

func (sim *Sim) UpdateObjectives(character *Character) {
	// periodically un-stuck all objectives to try again
	if sim.Time%config.CharacterObjectiveResetInterval == 0 {
		for i := range character.Objectives {
			character.Objectives[i].Stuck = false
		}
	}

	if character.Needs.Food >= config.NeedFoodMax && !character.HasObjective(EatObjective) {
		sim.AddObjective(character, EatObjective, 0)
	}

	if character.Needs.Water >= config.NeedWaterMax && !character.HasObjective(DrinkObjective) {
		sim.AddObjective(character, DrinkObjective, 0)
	}

	if character.Needs.Sleep >= config.NeedSleepMax && !character.HasObjective(SleepObjective) {
		sim.AddObjective(character, SleepObjective, 0)
	}

	if character.Needs.Food >= config.NeedFoodMax && sim.GetGrowingTilesCount() == 0 && !character.HasObjective(MakeFoodObjective) {
		sim.AddObjective(character, MakeFoodObjective, 0)
	}
}

func (sim *Sim) AddObjective(character *Character, objectiveType ObjectiveType, variant int) (createdObjective Objective) {
	objective := Objective{
		Type:    objectiveType,
		Variant: variant,
		Plan:    []Task{},
	}
	character.Objectives = append(character.Objectives, objective)
	return objective
}

func (character *Character) HasObjective(objectiveType ObjectiveType) bool {
	for _, objective := range character.Objectives {
		if objective.Type == objectiveType {
			return true
		}
	}
	return false
}

func (character *Character) CompleteObjective(objective *Objective) {
	fmt.Printf("Completing objective %v %v\n", character.Name, objective.Type)
	for i := len(character.Objectives) - 1; i >= 0; i-- {
		charObjective := character.Objectives[i]
		if charObjective.Type == objective.Type && charObjective.Variant == objective.Variant {
			character.Objectives = append(character.Objectives[:i], character.Objectives[i+1:]...)
		}
	}
}

func (sim *Sim) CheckIfObjectiveIsAchieved(character *Character, objective *Objective) {
	switch objective.Type {
	case EatObjective:
		if character.Needs.Food < 40 {
			character.CompleteObjective(objective)
		}
	case DrinkObjective:
		if character.Needs.Water < 40 {
			character.CompleteObjective(objective)
		}
	case SleepObjective:
		if character.Needs.Sleep < 10 {
			character.CompleteObjective(objective)
		}
	case MakeFoodObjective:
		if sim.GetGrowingTilesCount() >= config.PlantSeedsAtLeast {
			character.CompleteObjective(objective)
		}
	}
}

// Get the top non-stuck priority objective (lowest ObjectiveType is highest priority)
func (sim *Sim) GetTopPriorityObjective(character *Character) *Objective {
	if len(character.Objectives) == 0 {
		return nil
	}
	lowestIndex := -1
	for i := range character.Objectives {
		if !character.Objectives[i].Stuck && (lowestIndex == -1 || character.Objectives[i].Type < character.Objectives[lowestIndex].Type) {
			lowestIndex = i
		}
	}
	if lowestIndex == -1 {
		return nil
	}
	return &character.Objectives[lowestIndex]
}
