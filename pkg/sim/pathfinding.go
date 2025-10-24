package sim

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a tile in the pathfinding graph
type Node struct {
	X, Y     int
	F, G, H  float64 // costs for A* algorithm
	Parent   *Node
	MoveCost MoveCost
}

// PriorityQueue is the open set of nodes to explore
// implements heap.Interface for A* pathfinding
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F // Lower F value has higher priority
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
	item := old[0]     // Get the root element (lowest F value)
	old[0] = old[n-1]  // Move the last element to the root
	*pq = old[0 : n-1] // Remove the last element
	heap.Fix(pq, 0)    // Restore heap property after the swap
	return item
}

// FindPath finds the optimal path between two tiles in a region using A* algorithm
// if vicinity is set, we stop when reaching this distance of the target tile
func (s *Sim) FindPath(start TilePosition, end TilePosition, vicinity int) []TilePosition {
	// Initialize open and closed sets
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	closedSet := make(map[string]bool)

	// Create start and end nodes
	startNode := &Node{
		X:        start.X,
		Y:        start.Y,
		MoveCost: s.Tiles[start.Y*REGION_SIZE+start.X].MoveCost,
	}
	endNode := &Node{
		X:        end.X,
		Y:        end.Y,
		MoveCost: s.Tiles[end.Y*REGION_SIZE+end.X].MoveCost,
	}

	// Initialize start node
	startNode.G = 0
	startNode.H = heuristic(startNode, endNode)
	startNode.F = startNode.G + startNode.H

	// Add start node to open set
	heap.Push(openSet, startNode)

	// Main A* loop
	for openSet.Len() > 0 {
		// Get node with lowest F cost
		current := heap.Pop(openSet).(*Node)

		// Check if we reached the goal
		if vicinity > 0 {
			if current.X <= endNode.X+vicinity && current.X >= endNode.X-vicinity && current.Y <= endNode.Y+vicinity && current.Y >= endNode.Y-vicinity {
				return reconstructPath(current)
			}
		} else {
			if current.X == endNode.X && current.Y == endNode.Y {
				return reconstructPath(current)
			}
		}

		// Add current node to closed set
		closedSet[getNodeKey(current)] = true

		// Check neighbors
		directions := [][2]int{
			{0, 1},   // up
			{1, 1},   // up-right
			{1, 0},   // right
			{1, -1},  // down-right
			{0, -1},  // down
			{-1, -1}, // down-left
			{-1, 0},  // left
			{-1, 1},  // up-left
		}

		for _, dir := range directions {
			newX, newY := current.X+dir[0], current.Y+dir[1]

			// Check bounds
			if newX < 0 || newX >= REGION_SIZE || newY < 0 || newY >= REGION_SIZE {
				continue
			}

			// Check if tile is passable
			moveCost := s.Tiles[newY*REGION_SIZE+newX].MoveCost
			if moveCost == ImpassableCost {
				continue
			}

			// Skip if neighbor is in closed set
			neighborKey := fmt.Sprintf("%d,%d", newX, newY)
			if closedSet[neighborKey] {
				continue
			}

			// Calculate movement cost based on direction (diagonal vs orthogonal)
			dx := newX - current.X
			dy := newY - current.Y
			var movementMultiplier float64
			if dx != 0 && dy != 0 {
				// Diagonal movement - cost is sqrt(2) ≈ 1.414
				movementMultiplier = math.Sqrt2
			} else {
				// Orthogonal movement - cost is 1
				movementMultiplier = 1.0
			}

			// Calculate new G cost
			newG := current.G + float64(moveCost)*movementMultiplier

			// Check if node already exists in open set
			existingNode := findNodeInOpenSet(openSet, newX, newY)
			if existingNode == nil {
				// New node, add to open set
				neighbor := &Node{
					X:        newX,
					Y:        newY,
					MoveCost: moveCost,
					Parent:   current,
					G:        newG,
					H:        0, // Will be calculated below
				}
				neighbor.H = heuristic(neighbor, endNode)
				neighbor.F = neighbor.G + neighbor.H
				heap.Push(openSet, neighbor)
			} else if newG < existingNode.G {
				// Found a better path to existing node, update it
				existingNode.Parent = current
				existingNode.G = newG
				existingNode.H = heuristic(existingNode, endNode)
				existingNode.F = existingNode.G + existingNode.H
				// Note: We don't need to re-heapify since we're only updating values
			}
		}
	}

	// No path found
	return nil
}

// heuristic calculates the diagonal distance between two nodes
func heuristic(a, b *Node) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	// Using diagonal distance formula: max(dx, dy) + (sqrt2 - 1) * min(dx, dy)
	return math.Max(dx, dy) + 0.414*math.Min(dx, dy)
}

// findNodeInOpenSet finds a node in the open set by coordinates
func findNodeInOpenSet(openSet *PriorityQueue, x, y int) *Node {
	for _, node := range *openSet {
		if node.X == x && node.Y == y {
			return node
		}
	}
	return nil
}

// getNodeKey returns a unique key for a node
func getNodeKey(node *Node) string {
	return fmt.Sprintf("%d,%d", node.X, node.Y)
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
