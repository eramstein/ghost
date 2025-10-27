package sim

type ItemType int

const (
	ItemTypeFood ItemType = iota
	ItemTypeTool
	ItemTypeWeapon
)

func (sim *Sim) InitItems() {
	sim.AddItem(ItemTypeFood, Item{Type: ItemTypeFood})
}
