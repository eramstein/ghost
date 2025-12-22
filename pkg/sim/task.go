package sim

import (
	"fmt"
)

type TaskType int

const (
	NoTaskType TaskType = iota
	Move
	Eat
	Drink
	Sleep
	PickUp
	PlantSeed
)

func (tt TaskType) String() string {
	switch tt {
	case NoTaskType:
		return "None"
	case Move:
		return "Move"
	case Eat:
		return "Eat"
	case Drink:
		return "Drink"
	case Sleep:
		return "Sleep"
	case PickUp:
		return "Pick up"
	case PlantSeed:
		return "Plant seed"
	default:
		return "Unknown"
	}
}

func (sim *Sim) SetCurrentTask(character *Character) {
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective != nil {
		nextTask := sim.CreateNextTask(character, topObjective)
		if nextTask != nil {
			character.CurrentTask = nextTask
		} else {
			topObjective.Stuck = true
			fmt.Printf("Objective stuck because no task: %v\n", topObjective)
		}
	}
}

func (sim *Sim) WorkOnCurrentTask(character *Character) {
	task := character.CurrentTask
	if task == nil {
		return
	}
	switch task.Type {
	case Move:
		sim.MoveForTask(character)
	case Eat:
		sim.Eat(character)
	case Drink:
		sim.Drink(character)
	case Sleep:
		sim.Sleep(character)
	case PickUp:
		sim.PickUp(character)
	case PlantSeed:
		sim.PlantSeed(character)
	}
	if task.Progress >= 100 {
		sim.CompleteTask(character)
	}
}

// Create next task for a given objective
// Add it to the character's tasks if it's not nil
func (sim *Sim) CreateNextTask(character *Character, objective *Objective) (task *Task) {
	switch objective.Type {
	case EatObjective:
		task = sim.GetNextEatingTask(character, objective)
	case DrinkObjective:
		task = sim.GetNextDrinkingTask(character, objective)
	case SleepObjective:
		task = sim.GetNextSleepingTask(character, objective)
	case MakeFoodObjective:
		task = sim.GetNextMakingFoodTask(character, objective)
	}
	return task
}

func (sim *Sim) CompleteTask(character *Character) {
	if character.CurrentTask == nil {
		return
	}
	fmt.Printf("Completing task:  %v %v %v\n", character.Name, character.CurrentTask.Type, character.CurrentTask.Objective.Type)
	sim.CheckIfObjectiveIsAchieved(character, character.CurrentTask.Objective)
	character.CurrentTask = nil
}

func (sim *Sim) CancelTask(character *Character) {
	if character.CurrentTask == nil {
		return
	}
	fmt.Printf("Cancelling task:  %v %v %v\n", character.Name, character.CurrentTask.Type, character.CurrentTask.Objective)
	character.CurrentTask = nil
}
