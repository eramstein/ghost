package sim

import "gociv/pkg/config"

const PLAYER_SPEED = 200.0

func InitPlayer() Player {
	return Player{
		WorldPosition: WorldPosition{
			X: config.RegionSize / 2 * config.TileSize,
			Y: config.RegionSize / 2 * config.TileSize,
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
	if newX <= config.RegionSize*config.TileSize {
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
	if newY <= config.RegionSize*config.TileSize {
		p.WorldPosition.Y = newY
	}
}
