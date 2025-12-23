package sim

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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

// addItem adds an item and returns its global ID.
func (im *ItemManager) addItem(item Item, location ItemLocation) (int, error) {
	if len(im.freeSlots) == 0 {
		im.growCapacity()
	}

	// Pop last free slot for O(1) allocation
	last := len(im.freeSlots) - 1
	id := im.freeSlots[last]
	im.freeSlots = im.freeSlots[:last]

	item.ID = id
	item.Location = location
	im.items[id] = item
	im.usedSlots[id] = true

	return id, nil
}

// removeItem removes an item at the given ID and frees the slot.
func (im *ItemManager) removeItem(id int) error {
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

// getItem returns the item at the given ID.
func (im *ItemManager) getItem(id int) Item {
	return im.items[id]
}

// getItemPtr returns a pointer to the item at the given ID, or nil if invalid.
func (im *ItemManager) getItemPtr(id int) *Item {
	if id < 0 || id >= len(im.items) {
		return nil
	}
	if !im.usedSlots[id] {
		return nil
	}
	return &im.items[id]
}

// UpdateItemLocation updates the location of an item.
func (im *ItemManager) UpdateItemLocation(id int, location ItemLocation) error {
	if id < 0 || id >= len(im.items) {
		return fmt.Errorf("item id %d out of range", id)
	}
	if !im.usedSlots[id] {
		return fmt.Errorf("item id %d is not in use", id)
	}
	im.items[id].Location = location
	return nil
}

// getItems returns all items of a specific type (excluding empty slots).
func (im *ItemManager) getItems(itemType ItemType) []Item {
	var result []Item
	for i, item := range im.items {
		if im.usedSlots[i] && item.Type == itemType {
			result = append(result, item)
		}
	}
	return result
}

// getFreeSlotCount returns the number of free slots available.
func (im *ItemManager) getFreeSlotCount() int {
	return len(im.freeSlots)
}

// GetTotalCapacity returns the total capacity across all item types.
func (im *ItemManager) GetTotalCapacity() int {
	return len(im.items)
}

// getUsedSlotCount returns the number of used slots.
func (im *ItemManager) getUsedSlotCount() int {
	return len(im.items) - len(im.freeSlots)
}

// GobEncode encodes the ItemManager for gob serialization.
func (im *ItemManager) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	// Encode the unexported fields
	if err := enc.Encode(im.items); err != nil {
		return nil, err
	}
	if err := enc.Encode(im.usedSlots); err != nil {
		return nil, err
	}
	if err := enc.Encode(im.freeSlots); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GobDecode decodes the ItemManager from gob serialization.
func (im *ItemManager) GobDecode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	// Decode the unexported fields
	if err := dec.Decode(&im.items); err != nil {
		return err
	}
	if err := dec.Decode(&im.usedSlots); err != nil {
		return err
	}
	if err := dec.Decode(&im.freeSlots); err != nil {
		return err
	}

	return nil
}

// Item management convenience methods for Sim
func (s *Sim) AddItemToOwner(item Item, location ItemLocation, owner int) int {
	item.OwnedBy = owner
	return s.AddItem(item, location)
}
func (s *Sim) AddItem(item Item, location ItemLocation) int {
	item.OwnedBy = -1
	index, _ := s.ItemManager.addItem(item, location)
	if location.LocationType == LocTile {
		tile := s.GetTileAt(location.TilePosition)
		tile.AddItem(index)
	}
	return index
}
func (s *Sim) RemoveItem(id int) error {
	item := s.ItemManager.getItem(id)
	fmt.Printf("Removing item %d\n", id)
	if item.Location.LocationType == LocTile {
		tile := s.GetTileAt(item.Location.TilePosition)
		fmt.Printf("Removing item %d from tile %v with items %v\n", id, item.Location.TilePosition, tile.Items)
		tile.RemoveItem(id)
	} else if item.Location.LocationType == LocCharacter {
		// Remove from character inventory
		characterID := item.Location.CharacterID
		if characterID >= 0 && characterID < len(s.Characters) {
			character := &s.Characters[characterID]
			for i, invItemID := range character.Inventory {
				if invItemID == id {
					character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)
					fmt.Printf("Removed item %d from character %d inventory\n", id, characterID)
					break
				}
			}
		}
	}
	return s.ItemManager.removeItem(id)
}
func (s *Sim) DecreaseItemStackCount(id int) error {
	item := s.ItemManager.getItem(id)
	item.StackCount--
	if item.StackCount <= 0 {
		return s.RemoveItem(id)
	}
	return nil
}
func (s *Sim) GetItem(id int) Item {
	return s.ItemManager.getItem(id)
}
func (s *Sim) GetItemPtr(id int) *Item {
	return s.ItemManager.getItemPtr(id)
}
func (s *Sim) GetItems(itemType ItemType) []Item {
	return s.ItemManager.getItems(itemType)
}
func (s *Sim) GetItemCount() int {
	return s.ItemManager.getUsedSlotCount()
}
func (s *Sim) GetFreeItemSlots() int {
	return s.ItemManager.getFreeSlotCount()
}
