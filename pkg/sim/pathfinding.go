package sim

import (
	"container/heap"
	"fmt"
	"gociv/pkg/config"
	"math"
)

// Node represents a tile in the pathfinding graph
type Node struct {
	X, Y     int16
	F, G, H  float64 // costs for A* algorithm
	Parent   *Node
	MoveCost MoveCost
}

// PriorityQueue is the open set of nodes to explore
// implements heap.Interface for A* pathfinding
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Standard A*: lower F value has higher priority
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1] // heap.Pop swaps root to end, so we remove from end
	*pq = old[0 : n-1]
	return item
}

// FindPath finds the optimal path between two tiles in a region using A* algorithm
// if vicinity is set, we stop when reaching this distance of the target tile
func (s *Sim) FindPath(start TilePosition, end TilePosition, vicinity int) []TilePosition {
	// Initialize open and closed sets
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	closedSet := make(map[string]bool)
	openSetMap := make(map[string]*Node)  // Map for O(1) lookup of nodes in open set
	allNodesMap := make(map[string]*Node) // Map of ALL nodes we've seen (open + closed) for G comparison

	// Create start node
	startNode := &Node{
		X:        start.X,
		Y:        start.Y,
		G:        0,
		MoveCost: s.Tiles[start.Y*config.RegionSize+start.X].MoveCost,
	}
	startNode.H = heuristic(startNode.X, startNode.Y, end.X, end.Y)
	startNode.F = startNode.G + startNode.H

	// Add start node to open set
	startKey := getNodeKey(startNode.X, startNode.Y)
	heap.Push(openSet, startNode)
	openSetMap[startKey] = startNode
	allNodesMap[startKey] = startNode

	// Main A* loop
	for openSet.Len() > 0 {
		// Get node with lowest F cost
		current := heap.Pop(openSet).(*Node)
		delete(openSetMap, getNodeKey(current.X, current.Y))

		// Check if we reached the goal
		if vicinity > 0 {
			if current.X <= end.X+int16(vicinity) && current.X >= end.X-int16(vicinity) &&
				current.Y <= end.Y+int16(vicinity) && current.Y >= end.Y-int16(vicinity) {
				return reconstructPath(current)
			}
		} else {
			if current.X == end.X && current.Y == end.Y {
				return reconstructPath(current)
			}
		}

		// Add current node to closed set
		closedSet[getNodeKey(current.X, current.Y)] = true

		// Check all 8 neighbors
		for _, dir := range EightDirections {
			newX, newY := current.X+int16(dir[0]), current.Y+int16(dir[1])

			// Check bounds
			if newX < 0 || newX >= config.RegionSize || newY < 0 || newY >= config.RegionSize {
				continue
			}

			// Check if tile is passable
			tileIndex := newY*config.RegionSize + newX
			moveCost := s.Tiles[tileIndex].MoveCost
			if moveCost == ImpassableCost {
				continue
			}

			neighborKey := getNodeKey(newX, newY)

			// Calculate movement cost based on direction (diagonal vs orthogonal)
			var movementMultiplier float64
			if dir[0] != 0 && dir[1] != 0 {
				// Diagonal movement - cost is sqrt(2) ≈ 1.414
				movementMultiplier = math.Sqrt2
			} else {
				// Orthogonal movement - cost is 1
				movementMultiplier = 1.0
			}

			// Calculate new G cost
			// If moveCost is 0.0, use movementMultiplier as minimum cost to ensure paths accumulate cost
			tileCost := float64(moveCost)
			if tileCost <= 0.0 {
				tileCost = movementMultiplier
			} else {
				tileCost *= movementMultiplier
			}
			newG := current.G + tileCost

			// Check if we've seen this node before (open or closed)
			seenNode, seenBefore := allNodesMap[neighborKey]

			if seenBefore {
				// We've seen this node - check if new path is better
				// Use small epsilon to handle floating point precision
				if newG < seenNode.G-0.0001 {
					// Found a better path - update the node
					seenNode.Parent = current
					seenNode.G = newG
					seenNode.MoveCost = moveCost
					seenNode.F = seenNode.G + seenNode.H

					if closedSet[neighborKey] {
						// Node was closed - reopen it (heuristic was misleading)
						delete(closedSet, neighborKey)
						heap.Push(openSet, seenNode)
						openSetMap[neighborKey] = seenNode
					} else {
						// Node is in open set - remove, update, re-insert for reliable heap ordering
						// Find and remove from heap
						for i := 0; i < openSet.Len(); i++ {
							if (*openSet)[i] == seenNode {
								// Remove by swapping with last and truncating
								openSet.Swap(i, openSet.Len()-1)
								*openSet = (*openSet)[:openSet.Len()-1]
								// Fix heap from position we modified (if not last element)
								if i < openSet.Len() {
									heap.Fix(openSet, i)
								}
								break
							}
						}
						// Re-insert with updated F value
						heap.Push(openSet, seenNode)
					}
				}
			} else {
				// New node, add to open set
				neighbor := &Node{
					X:        newX,
					Y:        newY,
					MoveCost: moveCost,
					Parent:   current,
					G:        newG,
				}
				neighbor.H = heuristic(neighbor.X, neighbor.Y, end.X, end.Y)
				neighbor.F = neighbor.G + neighbor.H
				heap.Push(openSet, neighbor)
				openSetMap[neighborKey] = neighbor
				allNodesMap[neighborKey] = neighbor
			}
		}
	}

	// No path found
	return nil
}

// heuristic calculates the diagonal distance between two positions
func heuristic(x1, y1, x2, y2 int16) float64 {
	dx := math.Abs(float64(x1 - x2))
	dy := math.Abs(float64(y1 - y2))
	// Using diagonal distance formula: max(dx, dy) + (sqrt2 - 1) * min(dx, dy)
	return math.Max(dx, dy) + 0.414*math.Min(dx, dy)
}

// getNodeKey returns a unique key for a position
func getNodeKey(x, y int16) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// reconstructPath builds the path from end node to start node, excluding the initial position
func reconstructPath(endNode *Node) []TilePosition {
	path := make([]TilePosition, 0)
	current := endNode

	for current != nil && current.Parent != nil {
		path = append([]TilePosition{{X: current.X, Y: current.Y}}, path...)
		current = current.Parent
	}

	return path
}
