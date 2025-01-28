package main

import (
	"day25/locksandkeys"
	"flag"
	"fmt"
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
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Insert code here

	keys, locks := locksandkeys.ParseInput(challengeText)

	// fmt.Printf("Keys:\n")
	// for _, k := range keys {
	// 	k.Print()
	// }

	// fmt.Printf("\nLocks:\n")
	// for _, l := range locks {
	// 	l.Print()
	// }
	totalPartOne := 0
	loopCount := 0
	for _, k := range keys {
		for _, l := range locks {
			loopCount++
			if locksandkeys.DoTheyFit(&k, &l) {
				totalPartOne++
			}
		}
	}

	fmt.Printf("Total locks: %d, total keys: %d\n", len(locks), len(keys))
	fmt.Printf("Total loops: %d\n", loopCount)
	fmt.Printf("Total keys that fit: %d\n", totalPartOne)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
