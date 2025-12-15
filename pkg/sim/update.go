package sim

import "fmt"

const SIM_STEP = 1.0 // seconds per simulation tick

// Long running sim update based on SIM_STEP
func (s *Sim) LogicUpdate() {
	fmt.Println("TICK !")
	s.UpdateTime()
	s.UpdateCharacters()
	s.UpdatePlants()
}

// Things needed to be done every frame (movement...)
func (s *Sim) FrameUpdate(deltaTime float32) {
	for i := range s.Characters {
		s.Move(&s.Characters[i], deltaTime)
	}
}

func (sim *Sim) UpdateTime() {
	sim.Time++
	sim.Calendar.Minute++
	if sim.Calendar.Minute%60 == 0 {
		sim.Calendar.Hour++
		sim.Calendar.Minute = 0
	}
	if sim.Calendar.Hour >= 24 {
		sim.Calendar.Hour = 0
		sim.Calendar.Day++
	}
}
