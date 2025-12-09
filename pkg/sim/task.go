package sim

type TaskType int

const (
	NoTaskType TaskType = iota
	Move
	Eat
	Drink
	Sleep
	PickUp
	Build
	Chop
)
