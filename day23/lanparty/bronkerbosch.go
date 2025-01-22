package lanparty

// After some diligent research, I discovered the Bron Kerbosch algorithm that I will
// employ for this ridiculously large dataset.

func intersect(a, b []string) []string {
	set := make(map[string]bool)
	for _, item := range b {
		set[item] = true
	}
	intersection := make([]string, 0)
	for _, item := range a {
		if set[item] {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

// helper function to remove element from slice
func remove(slice []string, element string) []string {
	result := make([]string, 0)
	for _, item := range slice {
		if item != element {
			result = append(result, item)
		}
	}
	return result
}

func BronKerbosch(graph map[string]map[string]bool, R, P, X []string, maxClique *[]string) {
	if len(P) == 0 && len(X) == 0 {
		// R is a maximal clique
		if len(R) > len(*maxClique) {
			*maxClique = append([]string{}, R...)
		}
		return
	}

	// iterate over a copy of P to avoid modification during recursion
	Pcopy := append([]string{}, P...)
	for _, v := range Pcopy {
		// Neighbors of v
		neighbors := make([]string, 0)
		for neighbor := range graph[v] {
			neighbors = append(neighbors, neighbor)
		}

		BronKerbosch(
			graph,
			append(R, v),
			intersect(P, neighbors),
			intersect(X, neighbors),
			maxClique,
		)

		// move v from p to x
		P = remove(P, v)
		X = append(X, v)
	}
}
