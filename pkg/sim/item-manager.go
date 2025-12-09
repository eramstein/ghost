package sim

import "fmt"

const defaultItemCapacity = 200

// ItemManager manages items with memory reuse using a single slice.
type ItemManager struct {
	items     []Item // indexed by global item ID
	usedSlots []bool // mirrors items slice to track occupancy
	freeSlots []int  // stack of available indices
}

// NewItemManager creates a new item manager with a default capacity.
func NewItemManager() *ItemManager {
	im := &ItemManager{}
	im.SetCapacity(defaultItemCapacity)
	return im
}

// SetCapacity preallocates storage for items.
func (im *ItemManager) SetCapacity(capacity int) {
	im.items = make([]Item, capacity)
	im.usedSlots = make([]bool, capacity)
	im.freeSlots = make([]int, capacity)

	// Initialize all slots as free (LIFO for cache-friendly reuse)
	for i := 0; i < capacity; i++ {
		im.freeSlots[i] = capacity - 1 - i
	}
}

// growCapacity doubles the current storage to accommodate more items.
func (im *ItemManager) growCapacity() {
	current := len(im.items)
	newSlots := current
	if newSlots == 0 {
		newSlots = defaultItemCapacity
	}

	im.items = append(im.items, make([]Item, newSlots)...)
	im.usedSlots = append(im.usedSlots, make([]bool, newSlots)...)

	// Add new slots to the free list
	for i := 0; i < newSlots; i++ {
		im.freeSlots = append(im.freeSlots, current+newSlots-1-i)
	}
}

// AddItem adds an item and returns its global ID.
func (im *ItemManager) AddItem(item Item, location ItemLocation) (int, error) {
	if len(im.freeSlots) == 0 {
		im.growCapacity()
	}

	// Pop last free slot for O(1) allocation
	last := len(im.freeSlots) - 1
	id := im.freeSlots[last]
	im.freeSlots = im.freeSlots[:last]

	item.Location = location
	im.items[id] = item
	im.usedSlots[id] = true

	return id, nil
}

// RemoveItem removes an item at the given ID and frees the slot.
func (im *ItemManager) RemoveItem(id int) error {
	if id < 0 || id >= len(im.items) {
		return fmt.Errorf("item id %d out of range", id)
	}
	if !im.usedSlots[id] {
		return fmt.Errorf("item id %d is not in use", id)
	}

	// Clear the item
	im.items[id] = Item{}
	im.usedSlots[id] = false

	// Add the slot back to free list
	im.freeSlots = append(im.freeSlots, id)

	return nil
}

// GetItem returns the item at the given ID.
func (im *ItemManager) GetItem(id int) (Item, error) {
	if id < 0 || id >= len(im.items) {
		return Item{}, fmt.Errorf("item id %d out of range", id)
	}
	if !im.usedSlots[id] {
		return Item{}, fmt.Errorf("item id %d is not in use", id)
	}

	return im.items[id], nil
}

// GetItems returns all items of a specific type (excluding empty slots).
func (im *ItemManager) GetItems(itemType ItemType) []Item {
	var result []Item
	for i, item := range im.items {
		if im.usedSlots[i] && item.Type == itemType {
			result = append(result, item)
		}
	}
	return result
}

// GetFreeSlotCount returns the number of free slots available.
func (im *ItemManager) GetFreeSlotCount() int {
	return len(im.freeSlots)
}

// GetTotalCapacity returns the total capacity across all item types.
func (im *ItemManager) GetTotalCapacity() int {
	return len(im.items)
}

// GetUsedSlotCount returns the number of used slots.
func (im *ItemManager) GetUsedSlotCount() int {
	return len(im.items) - len(im.freeSlots)
}

// Item management convenience methods for Sim
func (s *Sim) AddItem(item Item, location ItemLocation) int {
	index, _ := s.ItemManager.AddItem(item, location)
	if location.LocationType == LocTile {
		tile := s.GetTileAt(location.TilePosition)
		tile.AddItem(ItemRef{ID: index, Type: item.Type}, location.TilePosition)
	}
	return index
}
func (s *Sim) RemoveItem(id int) error {
	return s.ItemManager.RemoveItem(id)
}
func (s *Sim) GetItem(id int) (Item, error) {
	return s.ItemManager.GetItem(id)
}
func (s *Sim) GetItems(itemType ItemType) []Item {
	return s.ItemManager.GetItems(itemType)
}
func (s *Sim) GetItemCount() int {
	return s.ItemManager.GetUsedSlotCount()
}
func (s *Sim) GetFreeItemSlots() int {
	return s.ItemManager.GetFreeSlotCount()
}
