package sim

import "fmt"

func ObjectiveFailed(character *Character, objective *Objective) {
	objective.Stuck = true
	fmt.Printf("Objective failed: %v %v\n", character.Name, objective.Type)
}
