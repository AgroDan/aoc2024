package lanparty

import "strings"

// func ParseLanparty(lines []string) map[string]*Computer {
// 	computers := make(map[string]*Computer)

// 	// now parse the lines
// 	for _, line := range lines {
// 		pcs := strings.Split(strings.TrimSpace(line), "-")
// 		// length should always be two
// 		if _, exists := computers[pcs[0]]; !exists {
// 			computers[pcs[0]] = NewComputer(pcs[0])
// 		}

// 		if _, exists := computers[pcs[1]]; !exists {
// 			computers[pcs[1]] = NewComputer(pcs[1])
// 		}

// 		// now to link them
// 		computers[pcs[0]].Connect(pcs[1])
// 		computers[pcs[1]].Connect(pcs[0])
// 	}
// 	return computers
// }

// func FindConnectedNetworks(computers map[string]*Computer) [][3]string {
// 	// I'm sure this will be obsolete after the part two question is unveiled,
// 	// but this will check the connected computers looking if the neighbors are
// 	// connected to the origin computer.

// 	retval := make([][3]string, 0)
// 	for _, v := range computers {
// 		for lk := range v.Links {
// 			// final layer...
// 			for _, lk2 := range computers[lk].Links {
// 				thisNetwork := [3]string{v.Name, lk, lk2}
// 				if lk2 != v.Name && computers[lk].IsConnected(lk2) {
// 					thisNetwork := [3]string{v.Name, lk, lk2}
// 				}
// 			computers[lk].IsConnected(v.Name)
// 			link := []string{v.Name, computers[lk].Name, }
// 		}
// 	}
// }

func ParseConnections(pairs []string) map[string]map[string]bool {
	// This should specify a connection between two computers by specifying computer[a][b], if that
	// exists then a connection exists between the two.
	outGraph := make(map[string]map[string]bool)
	for _, pair := range pairs {
		pcs := strings.Split(strings.TrimSpace(pair), "-")
		a, b := pcs[0], pcs[1]
		if outGraph[a] == nil {
			outGraph[a] = make(map[string]bool)
		}
		if outGraph[b] == nil {
			outGraph[b] = make(map[string]bool)
		}
		outGraph[a][b] = true
		outGraph[b][a] = true
	}
	return outGraph
}

func CountTripleNetworkedPartOne(computers map[string]map[string]bool) int {
	netCount := 0
	for a := range computers {
		for b := range computers[a] {
			if b > a { // to avoid dupes
				for c := range computers[b] {
					if c > b && computers[a][c] {
						if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
							netCount++
						}
					}
				}
			}
		}
	}
	return netCount
}
