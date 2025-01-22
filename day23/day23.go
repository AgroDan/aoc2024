package main

import (
	"day23/lanparty"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here

	party := lanparty.ParseConnections(lines)
	partOneTotal := lanparty.CountTripleNetworkedPartOne(party)

	fmt.Printf("Total sets of 3 computers with names starting with 't' (answer to part one): %d\n", partOneTotal)

	// for k := range party {
	// 	largestNet := lanparty.CountNetworkSize(party, k)
	// 	fmt.Printf("Computer %s is in a network of %d computers\n", k, largestNet)
	// 	// fmt.Printf("Is %s in a giant party with their neighbors? %t\n", k, lanparty.AreNodesConnected(party, k))
	// }

	// for k, v := range party {
	// 	fmt.Printf("Computer %s has %d links\n", k, len(v))
	// }

	// biggestLan := lanparty.FindLargestLanParty(party)
	P := []string{}
	for node := range party {
		P = append(P, node)
	}
	biggestLan := []string{}
	lanparty.BronKerbosch(party, []string{}, P, []string{}, &biggestLan)

	fmt.Printf("Largest interconnected network: %v\n", biggestLan)
	fmt.Printf("Size of largest network: %d\n", len(biggestLan))
	sort.Strings(biggestLan)
	fmt.Printf("Answer to part two: %v\n", strings.Join(biggestLan, ","))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
