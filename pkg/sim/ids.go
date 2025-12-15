package sim

import "sync/atomic"

var nextID uint32

func getNextID() int {
	return int(atomic.AddUint32(&nextID, 1))
}
