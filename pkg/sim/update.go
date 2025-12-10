package sim

import "fmt"

const SIM_STEP = 1.0 // seconds per simulation tick

// Long running sim update based on SIM_STEP
func (s *Sim) LogicUpdate() {
	fmt.Println("TICK !")
	s.Time++
	s.UpdateCharacters()
}

// Things needed to be done every frame (movement...)
func (s *Sim) FrameUpdate(deltaTime float32) {
	for i := range s.Characters {
		s.Move(&s.Characters[i], deltaTime)
	}
}
