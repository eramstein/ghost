package sim

import (
	"bytes"
	"encoding/gob"
)

// StructureManager manages structures with stable IDs and efficient lookups.
// Internally it keeps a slice of structures, an occupancy bitmap, and a free-list
// of reusable indices, similar to PlantManager.
type StructureManager struct {
	structures []Structure
	usedSlots  []bool
	freeSlots  []int16
}

const defaultStructureCapacity = 256

// NewStructureManager creates a new structure manager with a default capacity.
func NewStructureManager() *StructureManager {
	sm := &StructureManager{}
	sm.SetCapacity(defaultStructureCapacity)
	return sm
}

// SetCapacity preallocates storage for structures.
func (sm *StructureManager) SetCapacity(capacity int) {
	sm.structures = make([]Structure, capacity)
	sm.usedSlots = make([]bool, capacity)
	sm.freeSlots = make([]int16, capacity)

	for i := 0; i < capacity; i++ {
		// LIFO free list
		sm.freeSlots[i] = int16(capacity - 1 - i)
	}
}

// growCapacity doubles the current storage to accommodate more structures.
func (sm *StructureManager) growCapacity() {
	current := len(sm.structures)
	newSlots := current
	if newSlots == 0 {
		newSlots = defaultStructureCapacity
	}

	sm.structures = append(sm.structures, make([]Structure, newSlots)...)
	sm.usedSlots = append(sm.usedSlots, make([]bool, newSlots)...)

	for i := 0; i < newSlots; i++ {
		sm.freeSlots = append(sm.freeSlots, int16(current+newSlots-1-i))
	}
}

// AddStructure adds a structure and returns its global ID.
func (sm *StructureManager) AddStructure(structure Structure) int16 {
	if len(sm.freeSlots) == 0 {
		sm.growCapacity()
	}

	// Pop last free slot for O(1) allocation
	last := len(sm.freeSlots) - 1
	id := sm.freeSlots[last]
	sm.freeSlots = sm.freeSlots[:last]

	structure.ID = id
	sm.structures[id] = structure
	sm.usedSlots[id] = true

	return id
}

// RemoveStructure removes a structure at the given ID and frees the slot.
func (sm *StructureManager) RemoveStructure(id int16) {
	if id < 0 || id >= int16(len(sm.structures)) {
		return
	}
	if !sm.usedSlots[id] {
		return
	}

	sm.structures[id] = Structure{}
	sm.usedSlots[id] = false
	sm.freeSlots = append(sm.freeSlots, id)
}

// GetStructure returns the structure at the given ID and a bool indicating presence.
func (sm *StructureManager) GetStructure(id int16) (Structure, bool) {
	if id < 0 || id >= int16(len(sm.structures)) {
		return Structure{}, false
	}
	if !sm.usedSlots[id] {
		return Structure{}, false
	}
	return sm.structures[id], true
}

// GetStructurePtr returns a pointer to the structure at the given ID and a bool indicating presence.
func (sm *StructureManager) GetStructurePtr(id int16) (*Structure, bool) {
	if id < 0 || id >= int16(len(sm.structures)) {
		return nil, false
	}
	if !sm.usedSlots[id] {
		return nil, false
	}
	return &sm.structures[id], true
}

// GetStructuresByOwnerAndType returns a slice of structure pointers matching the given owner ID and structure type.
func (sm *StructureManager) GetStructuresByOwnerAndType(ownerID int8, structureType StructureType) []*Structure {
	result := make([]*Structure, 0)
	for id := range sm.structures {
		if sm.usedSlots[id] {
			s := &sm.structures[id]
			if s.Owner == ownerID && s.StructureType == structureType {
				result = append(result, s)
			}
		}
	}
	return result
}

// ForEach calls fn for each existing structure.
func (sm *StructureManager) ForEach(fn func(id int, s *Structure)) {
	for id := range sm.structures {
		if sm.usedSlots[id] {
			fn(id, &sm.structures[id])
		}
	}
}

// All returns a slice copy of all existing structures.
// This is mainly for read-only operations like rendering.
func (sm *StructureManager) All() []Structure {
	result := make([]Structure, 0, len(sm.structures))
	for i, s := range sm.structures {
		if sm.usedSlots[i] {
			result = append(result, s)
		}
	}
	return result
}

// Count returns the number of structures currently in use.
func (sm *StructureManager) Count() int {
	return len(sm.structures) - len(sm.freeSlots)
}

// GobEncode encodes the StructureManager for gob serialization.
func (sm *StructureManager) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(sm.structures); err != nil {
		return nil, err
	}
	if err := enc.Encode(sm.usedSlots); err != nil {
		return nil, err
	}
	if err := enc.Encode(sm.freeSlots); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GobDecode decodes the StructureManager from gob serialization.
func (sm *StructureManager) GobDecode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&sm.structures); err != nil {
		return err
	}
	if err := dec.Decode(&sm.usedSlots); err != nil {
		return err
	}
	if err := dec.Decode(&sm.freeSlots); err != nil {
		return err
	}

	return nil
}

// Convenience helpers for Sim

// GetStructures returns a snapshot slice of all structures.
func (sim *Sim) GetStructures() []Structure {
	if sim.StructureManager == nil {
		return nil
	}
	return sim.StructureManager.All()
}

// AddStructure adds a structure to the StructureManager and registers its ID on the corresponding tile.
func (sim *Sim) AddStructure(structure Structure) int16 {
	if sim.StructureManager == nil {
		sim.StructureManager = NewStructureManager()
	}

	id := sim.StructureManager.AddStructure(structure)

	// Register on tile
	tile := sim.GetTileAt(structure.Position)
	tile.Structure = id

	return id
}

// RemoveStructure removes a structure from the StructureManager and unregisters its ID from the corresponding tile.
func (sim *Sim) RemoveStructure(id int16) {
	if sim.StructureManager == nil {
		return
	}

	// Capture position before removal
	s, ok := sim.StructureManager.GetStructure(id)
	if ok {
		tile := sim.GetTileAt(s.Position)
		tile.Structure = -1
	}

	sim.StructureManager.RemoveStructure(id)
}

// GetStructureByID returns a structure pointer for a given ID, or nil if not found.
func (sim *Sim) GetStructurePtrByID(id int16) *Structure {
	if sim.StructureManager == nil {
		return nil
	}
	s, ok := sim.StructureManager.GetStructurePtr(id)
	if !ok {
		return nil
	}
	return s
}

// GetStructureByID returns a structure data for a given ID, or nil if not found.
func (sim *Sim) GetStructureByID(id int16) Structure {
	if sim.StructureManager == nil {
		return Structure{}
	}
	s, ok := sim.StructureManager.GetStructure(id)
	if !ok {
		return Structure{}
	}
	return s
}
