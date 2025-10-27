package sim

type Sim struct {
	UI          UIState
	Player      Player
	Tiles       []Tile
	Characters  []Character
	ItemManager *ItemManager
}

type Item struct {
	Type ItemType
}

type Character struct {
	Name          string
	WorldPosition WorldPosition
	Path          []TilePosition
	Needs         Needs
}

type UIState struct {
	EditMode       bool
	Pause          bool
	EditorTileType TileType
}

type Tile struct {
	Type     TileType
	Position TilePosition
	MoveCost MoveCost
}

type WorldPosition struct {
	X float32
	Y float32
}

type TilePosition struct {
	X int
	Y int
}

type Player struct {
	WorldPosition WorldPosition
}

type Needs struct {
	Hunger int
}
