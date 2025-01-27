package logic

import (
	"fmt"
	"strings"
	"utils"
)

func ConvertBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// I simply cannot believe that go does not offer boolean logic on BOOLEAN VARIABLES!
func Xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func And(a, b bool) bool {
	return a && b
}

func Or(a, b bool) bool {
	return a || b
}

func SwapInstructions(inst1, inst2 string) (string, string) {
	// this will be a little esoteric, but this will basically swap the end result
	// of two arbitrary instructions, keeping the operation the same but swapping
	// the destination of each.
	inst1Parts := strings.Fields(inst1)
	inst2Parts := strings.Fields(inst2)

	newinst1 := fmt.Sprintf("%s %s %s -> %s", inst1Parts[0], inst1Parts[1], inst1Parts[2], inst2Parts[4])
	newinst2 := fmt.Sprintf("%s %s %s -> %s", inst2Parts[0], inst2Parts[1], inst2Parts[2], inst1Parts[4])
	return newinst1, newinst2
}

func SwapInstructionsFromQueue(inst1, inst2 string, queue utils.GQueue[string]) utils.GQueue[string] {
	// This will look inside the queue and pull out the two instructions, then swap them.
	// To accomplish this without going back and messing with my utils package, I will simply
	// empty the current queue into a slice, swap the appropriate instructions, then re-add the
	// instructions in a new queue
	outQueue := utils.NewGQueue[string]()

	instList := make([]string, 0)
	for !queue.IsEmpty() {
		inst, _ := queue.Dequeue()
		instList = append(instList, inst)
	}

	swappedOne, swappedTwo := SwapInstructions(inst1, inst2)
	for i := 0; i < len(instList); i++ {
		if instList[i] == inst1 {
			outQueue.Enqueue(swappedOne)
		} else if instList[i] == inst2 {
			outQueue.Enqueue(swappedTwo)
		} else {
			outQueue.Enqueue(instList[i])
		}
	}
	return outQueue
}
