package lanparty

// Helper function to check if all nodes in a subset are interconnected
func isClique(graph map[string]map[string]bool, subset []string) bool {
	for i := 0; i < len(subset); i++ {
		for j := i + 1; j < len(subset); j++ {
			if !graph[subset[i]][subset[j]] {
				return false
			}
		}
	}
	return true
}

// Generate all subsets of nodes
func generateSubsets(nodes []string) [][]string {
	subsets := [][]string{}
	n := len(nodes)
	for i := 0; i < (1 << n); i++ {
		subset := []string{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nodes[j])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

// Find the largest lan party in the network
func FindLargestLanParty(graph map[string]map[string]bool) []string {
	nodes := []string{}
	for node := range graph {
		nodes = append(nodes, node)
	}
	// fmt.Printf("Amount of nodes: %d\n", len(nodes))
	subsets := generateSubsets(nodes)
	// fmt.Printf("Amount of subsets: %d\n", len(subsets))
	maxNet := []string{}

	for _, subset := range subsets {
		if len(subset) > len(maxNet) && isClique(graph, subset) {
			maxNet = subset
		}
	}

	return maxNet
}
