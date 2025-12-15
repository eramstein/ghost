package sim

type Sim struct {
	Time        int // in minutes since the start of the simulation
	Calendar    Calendar
	UI          UIState
	Player      Player
	Tiles       []Tile
	Characters  []Character
	Plants      []Plant
	ItemManager *ItemManager
}

type Calendar struct {
	Minute int
	Hour   int
	Day    int
}

type Item struct {
	ID         int
	Type       ItemType
	Variant    int
	Location   ItemLocation
	OwnedBy    int   // character id, -1 if not owned
	Efficiency int   // for food it's nutrition value
	Durability int   // for materials it's how many builds they can support
	Items      []int // Item IDs
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
	CurrentTask   *Task
	Objectives    []Objective
	Ambitions     []Ambition
	Inventory     []int // Object IDs
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
	Objective      *Objective
	Progress       float32       // by default, 0 to 1, as percent of task already done, but can be used otherwise like for movement
	ProductType    int           // optional, precises the task is producing based on the Task Type, for example for bulding tasks it's the StructureType to build (e.g. Wall)
	ProductVariant int           // optional, further precises the task's product by providing a variant (e.g. Wooden Wall, Stone Wall)
	TargetItem     *Item         // optional, e.g. for eating tasks it's the food item to eat
	TargetTile     *TilePosition // optional, e.g. for building it's the tile ot build on
	MaterialSource *Item         // optional, for bulding tasks it's the material item to use
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

// Plants grow and can produce edible or craft materials (fruits, wood, etc.)
type Plant struct {
	ID          int
	Position    TilePosition
	PlantType   PlantType
	Variant     int
	GrowthStage int // 0-100
	GrowthRate  int // How many growth stages per update
	Produces    PlantProduction
}

type PlantProduction struct {
	Type            ItemType
	Variant         int
	ProductionStage int // 0-100
	ProductionRate  int // How many production stages per update
}
