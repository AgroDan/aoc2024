package lanparty

import "fmt"

// For every single computer, check to see if all the nodes it is connected to
// are all connected to each other.

func AreNodesConnected(graph map[string]map[string]bool, node string) bool {
	// Given a node, check to see if all the nodes it is connected to are also
	// connected to each other.
	connected := graph[node]
	for k := range connected {
		for lk := range graph[k] {
			if !graph[lk][node] {
				return false
			}
		}
	}
	return true
}

func CountNetworkSize(graph map[string]map[string]bool, node string) int {
	// consider this data:
	// co:
	// 	- ka
	// 	- ta
	//	- de
	//	- tc
	members := 0
	fmt.Printf("\nChecking %s\n", node)
	for k := range graph[node] {
		// each member will be checked
		// this is now looping ka,ta,de,tc
		for j := range graph[node] {
			if k == node || j == node || k == j {
				continue
			}
			if graph[j][k] {
				fmt.Printf("%s contains %s\n", j, k)
				members++
			}
		}

	}
	return members
}
