package sim

import (
	"bytes"
	"encoding/gob"
)

// PlantManager manages plants with stable IDs and efficient lookups.
// Internally it keeps a slice of plants, an occupancy bitmap, and a free-list
// of reusable indices, similar to ItemManager.
type PlantManager struct {
	plants    []Plant
	usedSlots []bool
	freeSlots []int
}

const defaultPlantCapacity = 256

// NewPlantManager creates a new plant manager with a default capacity.
func NewPlantManager() *PlantManager {
	pm := &PlantManager{}
	pm.SetCapacity(defaultPlantCapacity)
	return pm
}

// SetCapacity preallocates storage for plants.
func (pm *PlantManager) SetCapacity(capacity int) {
	pm.plants = make([]Plant, capacity)
	pm.usedSlots = make([]bool, capacity)
	pm.freeSlots = make([]int, capacity)

	for i := 0; i < capacity; i++ {
		// LIFO free list, like ItemManager
		pm.freeSlots[i] = capacity - 1 - i
	}
}

// growCapacity doubles the current storage to accommodate more plants.
func (pm *PlantManager) growCapacity() {
	current := len(pm.plants)
	newSlots := current
	if newSlots == 0 {
		newSlots = defaultPlantCapacity
	}

	pm.plants = append(pm.plants, make([]Plant, newSlots)...)
	pm.usedSlots = append(pm.usedSlots, make([]bool, newSlots)...)

	for i := 0; i < newSlots; i++ {
		pm.freeSlots = append(pm.freeSlots, current+newSlots-1-i)
	}
}

// addPlant adds a plant and returns its global ID.
func (pm *PlantManager) addPlant(plant Plant) int {
	if len(pm.freeSlots) == 0 {
		pm.growCapacity()
	}

	// Pop last free slot for O(1) allocation
	last := len(pm.freeSlots) - 1
	id := pm.freeSlots[last]
	pm.freeSlots = pm.freeSlots[:last]

	plant.ID = id
	pm.plants[id] = plant
	pm.usedSlots[id] = true

	return id
}

// removePlant removes a plant at the given ID and frees the slot.
func (pm *PlantManager) removePlant(id int) {
	if id < 0 || id >= len(pm.plants) {
		return
	}
	if !pm.usedSlots[id] {
		return
	}

	pm.plants[id] = Plant{}
	pm.usedSlots[id] = false
	pm.freeSlots = append(pm.freeSlots, id)
}

// getPlant returns the plant at the given ID and a bool indicating presence.
func (pm *PlantManager) getPlant(id int) (Plant, bool) {
	if id < 0 || id >= len(pm.plants) {
		return Plant{}, false
	}
	if !pm.usedSlots[id] {
		return Plant{}, false
	}
	return pm.plants[id], true
}

// getPlantPtr returns a pointer to the plant at the given ID and a bool indicating presence.
func (pm *PlantManager) getPlantPtr(id int) (*Plant, bool) {
	if id < 0 || id >= len(pm.plants) {
		return nil, false
	}
	if !pm.usedSlots[id] {
		return nil, false
	}
	return &pm.plants[id], true
}

// ForEach calls fn for each existing plant.
func (pm *PlantManager) ForEach(fn func(id int, p *Plant)) {
	for id := range pm.plants {
		if pm.usedSlots[id] {
			fn(id, &pm.plants[id])
		}
	}
}

// All returns a slice copy of all existing plants.
// This is mainly for read‑only operations like rendering.
func (pm *PlantManager) All() []Plant {
	result := make([]Plant, 0, len(pm.plants))
	for i, p := range pm.plants {
		if pm.usedSlots[i] {
			result = append(result, p)
		}
	}
	return result
}

// Count returns the number of plants currently in use.
func (pm *PlantManager) Count() int {
	return len(pm.plants) - len(pm.freeSlots)
}

// GobEncode encodes the PlantManager for gob serialization.
func (pm *PlantManager) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	// Encode the unexported fields
	if err := enc.Encode(pm.plants); err != nil {
		return nil, err
	}
	if err := enc.Encode(pm.usedSlots); err != nil {
		return nil, err
	}
	if err := enc.Encode(pm.freeSlots); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GobDecode decodes the PlantManager from gob serialization.
func (pm *PlantManager) GobDecode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	// Decode the unexported fields
	if err := dec.Decode(&pm.plants); err != nil {
		return err
	}
	if err := dec.Decode(&pm.usedSlots); err != nil {
		return err
	}
	if err := dec.Decode(&pm.freeSlots); err != nil {
		return err
	}

	return nil
}

// Convenience helpers for other packages

// GetPlants returns a snapshot slice of all plants.
func (sim *Sim) GetPlants() []Plant {
	if sim.PlantManager == nil {
		return nil
	}
	return sim.PlantManager.All()
}

// GetPlantByID returns a plant pointer for a given ID, or nil if not found.
func (sim *Sim) GetPlantByID(id int) *Plant {
	if sim.PlantManager == nil {
		return nil
	}
	p, ok := sim.PlantManager.getPlantPtr(id)
	if !ok {
		return nil
	}
	return p
}
