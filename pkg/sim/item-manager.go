package sim

import "fmt"

// ItemManager manages items with memory reuse
type ItemManager struct {
	items     map[ItemType][]Item
	freeSlots map[ItemType][]int // tracks free slots for each item type
	capacity  map[ItemType]int   // capacity for each item type
}

// NewItemManager creates a new item manager with specified capacities
func NewItemManager() *ItemManager {
	im := &ItemManager{
		items:     make(map[ItemType][]Item),
		freeSlots: make(map[ItemType][]int),
		capacity:  make(map[ItemType]int),
	}

	// Initialize with default capacities
	im.SetCapacity(ItemTypeFood, 100)
	im.SetCapacity(ItemTypeTool, 50)
	im.SetCapacity(ItemTypeWeapon, 50)

	return im
}

// SetCapacity sets the capacity for a specific item type
func (im *ItemManager) SetCapacity(itemType ItemType, capacity int) {
	im.capacity[itemType] = capacity
	im.items[itemType] = make([]Item, capacity)
	im.freeSlots[itemType] = make([]int, 0, capacity)

	// Initialize all slots as free
	for i := 0; i < capacity; i++ {
		im.freeSlots[itemType] = append(im.freeSlots[itemType], i)
	}
}

// AddItem adds an item and returns the index where it was placed
func (im *ItemManager) AddItem(itemType ItemType, item Item, location ItemLocation) (int, error) {
	freeSlots, exists := im.freeSlots[itemType]
	if !exists {
		return -1, fmt.Errorf("item type %d not initialized", itemType)
	}

	if len(freeSlots) == 0 {
		return -1, fmt.Errorf("no free slots available for item type %d", itemType)
	}

	// Get the first free slot
	index := freeSlots[0]

	// Remove the slot from free list
	im.freeSlots[itemType] = freeSlots[1:]

	// Place the item in the slice
	item.Type = itemType
	item.Location = location
	im.items[itemType][index] = item

	return index, nil
}

// RemoveItem removes an item at the given index and frees the slot
func (im *ItemManager) RemoveItem(itemType ItemType, index int) error {
	items, exists := im.items[itemType]
	if !exists {
		return fmt.Errorf("item type %d not initialized", itemType)
	}

	if index < 0 || index >= len(items) {
		return fmt.Errorf("index %d out of range for item type %d", index, itemType)
	}

	// Clear the item
	im.items[itemType][index] = Item{}

	// Add the slot back to free list
	im.freeSlots[itemType] = append(im.freeSlots[itemType], index)

	return nil
}

// GetItem returns the item at the given index
func (im *ItemManager) GetItem(itemType ItemType, index int) (Item, error) {
	items, exists := im.items[itemType]
	if !exists {
		return Item{}, fmt.Errorf("item type %d not initialized", itemType)
	}

	if index < 0 || index >= len(items) {
		return Item{}, fmt.Errorf("index %d out of range for item type %d", index, itemType)
	}

	return items[index], nil
}

// GetItems returns all items of a specific type (excluding empty slots)
func (im *ItemManager) GetItems(itemType ItemType) []Item {
	items, exists := im.items[itemType]
	if !exists {
		return nil
	}

	// Filter out empty items
	var result []Item
	for _, item := range items {
		if item.Type != 0 || itemType == 0 { // Include if not empty or if it's the zero value type
			result = append(result, item)
		}
	}

	return result
}

// GetFreeSlotCount returns the number of free slots for an item type
func (im *ItemManager) GetFreeSlotCount(itemType ItemType) int {
	return len(im.freeSlots[itemType])
}

// GetTotalCapacity returns the total capacity for an item type
func (im *ItemManager) GetTotalCapacity(itemType ItemType) int {
	return im.capacity[itemType]
}

// GetUsedSlotCount returns the number of used slots for an item type
func (im *ItemManager) GetUsedSlotCount(itemType ItemType) int {
	return im.capacity[itemType] - len(im.freeSlots[itemType])
}

// Item management convenience methods for Sim
func (s *Sim) AddItem(itemType ItemType, item Item, location ItemLocation) {
	index, _ := s.ItemManager.AddItem(itemType, item, location)
	if location.LocationType == LocTile {
		tile := s.GetTileAt(location.TilePosition)
		tile.AddItem(ItemRef{Type: itemType, Index: index}, location.TilePosition)
	}
}
func (s *Sim) RemoveItem(itemType ItemType, index int) error {
	return s.ItemManager.RemoveItem(itemType, index)
}
func (s *Sim) GetItem(itemType ItemType, index int) (Item, error) {
	return s.ItemManager.GetItem(itemType, index)
}
func (s *Sim) GetItems(itemType ItemType) []Item {
	return s.ItemManager.GetItems(itemType)
}
func (s *Sim) GetItemCount(itemType ItemType) int {
	return s.ItemManager.GetUsedSlotCount(itemType)
}
func (s *Sim) GetFreeItemSlots(itemType ItemType) int {
	return s.ItemManager.GetFreeSlotCount(itemType)
}
