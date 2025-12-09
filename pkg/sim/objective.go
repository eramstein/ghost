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
	BuildObjective
)

func (sim *Sim) UpdateObjectives(character *Character) {
	// periodically un-stuck all objectives to try again
	if sim.Time%config.CharacterObjectiveResetInterval == 0 {
		for _, objective := range character.Objectives {
			objective.Stuck = false
		}
	}

	if character.Needs.Food >= 50 && !character.HasObjective(EatObjective) {
		sim.AddObjective(character, EatObjective, 0)
	}

	if character.Needs.Water >= 50 && !character.HasObjective(DrinkObjective) {
		sim.AddObjective(character, DrinkObjective, 0)
	}

	if character.Needs.Sleep >= 50 && !character.HasObjective(SleepObjective) {
		sim.AddObjective(character, SleepObjective, 0)
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
	for i, charObjective := range character.Objectives {
		if &charObjective == objective {
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
	}
}

// Get the top non-stuck priority objective (lowest ObjectiveType is highest priority)
func (sim *Sim) GetTopPriorityObjective(character *Character) *Objective {
	if len(character.Objectives) == 0 {
		return nil
	}
	lowestObjective := character.Objectives[0]
	for _, objective := range character.Objectives {
		if !objective.Stuck && objective.Type < lowestObjective.Type {
			lowestObjective = objective
		}
	}
	return &lowestObjective
}
