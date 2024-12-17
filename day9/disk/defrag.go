package disk

// This will involve the defragmentation process. To accomplish this,
// a Queue will be utilized for the "from" and the "to" sections, where
// they will be queued up counting from the highest to the lowest (from)
// and adding to the lowest to the highest (to). If the ID of both
// queues are the same, stop the loop.

func Defragment(filesOnDisk []File) []File {
	// defragments a disk
	var workingFrom, workingTo Queue
	for i := 0; i < len(filesOnDisk); i++ {
		workingTo.Enqueue(filesOnDisk[i])
	}
	for j := len(filesOnDisk) - 1; j >= 0; j-- {
		workingFrom.Enqueue(filesOnDisk[j])
	}
	return defrag_action(workingFrom, workingTo)
}

func defrag_action(from, to Queue) []File {
	// This will assume the from and to queues are in the proper order.
	var retval []File

	workingFrom, err := from.Dequeue()
	if err != nil {
		panic("empty from queue to start from")
	}
	workingTo, err := to.Dequeue()
	if err != nil {
		panic("empty to queue to start from")
	}

	for {
		if workingTo.Id == workingFrom.Id {
			break
			// we intersected, break
		}

		if workingTo.FreeSpace() == 0 {
			// no more free space, move onto
			// the next one
			retval = append(retval, workingTo)
			workingTo, err = to.Dequeue()
			if err != nil {
				// no more free space, so break
				break
			}
			continue
		}

		if workingFrom.IsEmpty() {
			// empty contiguous block, let's move on
			workingFrom, err = from.Dequeue()
			if err != nil {
				// no more from points in the queue
				break
			}
			continue
		}

		// otherwise we'll move stuff
		err = workingTo.Recv(&workingFrom)
		if err != nil {
			panic(err)
		}
	}

	// is there anything left?
	if !workingFrom.IsEmpty() {
		retval = append(retval, workingFrom)
	}
	return retval
}

// This is for part two. only going to move the file
// if and ONLY if it can fit the whole file. Otherwise
// it stays in the current position.
func DefragmentPartTwo(filesOnDisk *[]File) {
	// This works the same as the above, just sends it to another function
	// var workingFrom []File
	// var workingTo []File
	// for i := 0; i < len(*filesOnDisk); i++ {
	// 	workingTo = append(workingTo, (*filesOnDisk)[i])
	// }
	for i := len((*filesOnDisk)) - 1; i >= 0; i-- {
		// in reverse this time
		// workingFrom = append(workingFrom, (*filesOnDisk)[i])
		searchForFile(&(*filesOnDisk)[i], filesOnDisk)
	}
}

func searchForFile(thisFile *File, fileList *[]File) bool {
	// will search in order of the entire file list
	// looking for a file this will fit. If it can,
	// it will move the file contents to the file
	// list's contiguous free space. returns true
	// if so, false if nothing changed.
	for i := 0; i < len((*fileList)); i++ {
		if thisFile.Id == (*fileList)[i].Id {
			// we looped until we hit ourselves
			// so ditch it
			break
		}
		err := (*fileList)[i].RecvOriginFile(thisFile)
		if err != nil {
			continue
		}
		// otherwise it worked so
		return true
	}
	return false
}
