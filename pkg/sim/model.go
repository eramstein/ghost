package sim

type Sim struct {
	Time        int // in minutes since the start of the simulation
	UI          UIState
	Player      Player
	Tiles       []Tile
	Characters  []Character
	ItemManager *ItemManager
}

type Item struct {
	Type     ItemType
	Location ItemLocation
}

type ItemLocation struct {
	LocationType ItemLocationType
	TilePosition TilePosition
	CharacterID  int
}

type Character struct {
	ID            int
	Name          string
	WorldPosition WorldPosition
	TilePosition  TilePosition
	Path          []TilePosition
	Needs         Needs
	Task          Task
	Objectives    []Objective
	Ambitions     []Ambition
}

type UIState struct {
	EditMode            bool
	Pause               bool
	EditorTileType      TileType
	SelectedCharacterID int
}

type Tile struct {
	Type     TileType
	Position TilePosition
	MoveCost MoveCost
	Items    []int
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
	Food  int
	Water int
	Sleep int
}

type Task struct {
	ID             uint64
	Type           TaskType
	ObjectiveID    int
	Progress       float32      // by default, 0 to 1, as percent of task already done, but can be used otherwise like for movement
	ProductType    int          // optional, precises the task is producing based on the Task Type, for example for bulding tasks it's the StructureType to build (e.g. Wall)
	ProductVariant int          // optional, further precises the task's product by providing a variant (e.g. Wooden Wall, Stone Wall)
	TargetItemID   int          // optional (-1 for nil), e.g. for eating tasks it's the food item to eat
	TargetTile     TilePosition // optional, e.g. for building it's the tile ot build on
	SourceItemID   int          // optional (-1 for nil), for bulding tasks it's the material item to use
}

type Objective struct {
	Type    ObjectiveType
	Variant int // optional, further precises the objective by providing a variant (e.g. "build a house")
	Stuck   bool
	Plan    []Task // optional, sometimes we pre-plan list of tasks as the objective is defined
}

type Ambition struct {
	Description string
}
