package dijkstra_algorithm

import (
	"math"
)

func DijkstraAlgorithm(graph map[string]map[string]int, startNode string, endNode string) []string {
	peaks := make(map[string]int)
	peaks[endNode] = math.MaxInt32

	path := make(map[string]string)
	path[endNode] = ""

	for node, peak := range graph[startNode] {
		peaks[node] = peak
		path[node] = startNode
	}

	visited := make(map[string]bool)
	smallNode := findSmallestNode(peaks, visited)

	for smallNode != "" {
		for node, peak := range graph[smallNode] {
			newPeak := peaks[smallNode] + peak
			if newPeak < peaks[node] {
				peaks[node] = newPeak
				path[node] = smallNode
			}
		}

		visited[smallNode] = true
		smallNode = findSmallestNode(peaks, visited)
	}

	distance := []string{endNode, path[endNode]}
	distanceNode := path[endNode]
	for distanceNode != "" {
		distanceNode = path[distanceNode]
		if distanceNode != "" {
			distance = append(distance, distanceNode)
		}
	}

	return reverseDistance(distance)
}

// Find small node
func findSmallestNode(peaks map[string]int, visited map[string]bool) string {
	element := math.MaxInt32
	smallNode := ""

	for node, peak := range peaks {
		if _, isVisited := visited[node]; peak < element && !isVisited {
			smallNode = node
			element = peak
		}
	}

	return smallNode
}

// Reverse distance
func reverseDistance(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseDistance(input[1:]), input[0])
}
