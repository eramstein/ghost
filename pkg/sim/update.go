package sim

import "fmt"

const SIM_STEP = 1.0 // seconds per simulation tick

// Long running sim update based on SIM_STEP
func (s *Sim) LogicUpdate() {
	fmt.Println("TICK !")
	for i := range s.Characters {
		if len(s.Characters[i].Path) == 0 {
			s.Characters[i].Path = s.FindPath(
				TilePosition{
					X: int(s.Characters[i].WorldPosition.X / TILE_SIZE),
					Y: int(s.Characters[i].WorldPosition.Y / TILE_SIZE),
				},
				TilePosition{
					X: 1,
					Y: 1,
				}, 0)
		}
	}
}

// Things needed to be done every frame (movement...)
func (s *Sim) FrameUpdate(deltaTime float32) {
	for i := range s.Characters {
		s.Characters[i].Move(deltaTime)
	}
}
