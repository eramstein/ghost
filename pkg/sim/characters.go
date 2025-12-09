package sim

import "gociv/pkg/config"

const CHARACTER_SPEED = 100

func (sim *Sim) InitCharacters() {
	sim.MakeCharacter("Henry", TilePosition{
		X: 10,
		Y: 10,
	})
	sim.MakeCharacter("Emma", TilePosition{
		X: 11,
		Y: 13,
	})
	sim.MakeCharacter("Lise", TilePosition{
		X: 11,
		Y: 14,
	})
	sim.MakeCharacter("Ousmane", TilePosition{
		X: 12,
		Y: 10,
	})
	sim.MakeCharacter("Molly", TilePosition{
		X: 12,
		Y: 12,
	})
	sim.MakeCharacter("Robert", TilePosition{
		X: 20,
		Y: 14,
	})
	sim.MakeCharacter("Didier", TilePosition{
		X: 20,
		Y: 10,
	})
	sim.MakeCharacter("Morgane", TilePosition{
		X: 20,
		Y: 12,
	})
}

func (sim *Sim) MakeCharacter(name string, pos TilePosition) {
	character := Character{
		ID:           len(sim.Characters),
		Name:         name,
		TilePosition: pos,
		WorldPosition: WorldPosition{
			X: float32(pos.X*TILE_SIZE + TILE_SIZE/2),
			Y: float32(pos.Y*TILE_SIZE + TILE_SIZE/2),
		},
	}
	sim.Characters = append(sim.Characters, character)
	sim.GetTileAt(pos).CharacterID = character.ID
}

func (sim *Sim) UpdateCharacters() {
	for i := range sim.Characters {
		character := &sim.Characters[i]
		character.UpdateNeeds()
		sim.UpdateObjectives(character)
	}
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food += config.NeedFoodTick
	character.Needs.Water += config.NeedWaterTick
	character.Needs.Sleep += config.NeedSleepTick
}

func (c *Character) Move(deltaTime float32) {
	if len(c.Path) > 0 {
		nextTile := c.Path[0]
		nextTileWorldPosition := WorldPosition{
			X: float32(nextTile.X),
			Y: float32(nextTile.Y),
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
