package sim

const PLAYER_SPEED = 200.0

func InitPlayer() Player {
	return Player{
		WorldPosition: WorldPosition{
			X: REGION_SIZE / 2 * TILE_SIZE,
			Y: REGION_SIZE / 2 * TILE_SIZE,
		},
	}
}

func (p *Player) MoveLeft(deltaTime float32) {
	newX := p.WorldPosition.X - PLAYER_SPEED*deltaTime
	if newX >= 0 {
		p.WorldPosition.X = newX
	}
}

func (p *Player) MoveRight(deltaTime float32) {
	newX := p.WorldPosition.X + PLAYER_SPEED*deltaTime
	if newX <= REGION_SIZE*TILE_SIZE {
		p.WorldPosition.X = newX
	}
}

func (p *Player) MoveUp(deltaTime float32) {
	newY := p.WorldPosition.Y - PLAYER_SPEED*deltaTime
	if newY >= 0 {
		p.WorldPosition.Y = newY
	}
}

func (p *Player) MoveDown(deltaTime float32) {
	newY := p.WorldPosition.Y + PLAYER_SPEED*deltaTime
	if newY <= REGION_SIZE*TILE_SIZE {
		p.WorldPosition.Y = newY
	}
}
