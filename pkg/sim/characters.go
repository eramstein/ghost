package sim

const CHARACTER_SPEED = 100

func InitCharacters() []Character {
	return []Character{
		{
			Name: "Henry",
			WorldPosition: WorldPosition{
				X: 10*TILE_SIZE + TILE_SIZE/2,
				Y: 10*TILE_SIZE + TILE_SIZE/2,
			},
		},
		{
			Name: "Emma",
			WorldPosition: WorldPosition{
				X: 11 * TILE_SIZE,
				Y: 13 * TILE_SIZE,
			},
		},
		{
			Name: "Lise",
			WorldPosition: WorldPosition{
				X: 11 * TILE_SIZE,
				Y: 14 * TILE_SIZE,
			},
		},
		{
			Name: "Ousmane",
			WorldPosition: WorldPosition{
				X: 12 * TILE_SIZE,
				Y: 10 * TILE_SIZE,
			},
		},
		{
			Name: "Molly",
			WorldPosition: WorldPosition{
				X: 12 * TILE_SIZE,
				Y: 12 * TILE_SIZE,
			},
		},
		{
			Name: "Robert",
			WorldPosition: WorldPosition{
				X: 20 * TILE_SIZE,
				Y: 14 * TILE_SIZE,
			},
		},
		{
			Name: "Didier",
			WorldPosition: WorldPosition{
				X: 20 * TILE_SIZE,
				Y: 10 * TILE_SIZE,
			},
		},
		{
			Name: "Morgane",
			WorldPosition: WorldPosition{
				X: 20 * TILE_SIZE,
				Y: 12 * TILE_SIZE,
			},
		},
	}
}

func (c *Character) Move(deltaTime float32) {
	if len(c.Path) > 0 {
		nextTile := c.Path[0]
		nextTileWorldPosition := WorldPosition{
			X: float32(nextTile.X)*TILE_SIZE + TILE_SIZE/2,
			Y: float32(nextTile.Y)*TILE_SIZE + TILE_SIZE/2,
		}
		var direction WorldPosition
		if nextTileWorldPosition.X-c.WorldPosition.X > 0 {
			direction.X = 1
		} else {
			direction.X = -1
		}
		if nextTileWorldPosition.Y-c.WorldPosition.Y > 0 {
			direction.Y = 1
		} else {
			direction.Y = -1
		}

		c.WorldPosition.X += direction.X * CHARACTER_SPEED * deltaTime
		c.WorldPosition.Y += direction.Y * CHARACTER_SPEED * deltaTime

		dx := c.WorldPosition.X - nextTileWorldPosition.X
		dy := c.WorldPosition.Y - nextTileWorldPosition.Y
		distance := dx*dx + dy*dy

		if distance < (0.2*TILE_SIZE)*(0.2*TILE_SIZE) {
			if ((direction.X == 1 && c.WorldPosition.X >= nextTileWorldPosition.X) ||
				(direction.X == -1 && c.WorldPosition.X <= nextTileWorldPosition.X)) &&
				((direction.Y == 1 && c.WorldPosition.Y >= nextTileWorldPosition.Y) ||
					(direction.Y == -1 && c.WorldPosition.Y <= nextTileWorldPosition.Y)) {
				c.Path = c.Path[1:]
			}
		}
	}
}
