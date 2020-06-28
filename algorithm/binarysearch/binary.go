package binarysearch

// BinarySearch makes binary searching of value v in slice a
// from start to stop position and returns index of found
// item or -1 (failure)
func BinarySearch(a []int, start, stop int, v int) int {
	// Check bounding conditions
	if len(a) == 0 {
		return -1
	}

	// Evaluate index of a middle item
	idx := (stop + start) / 2
	if a[idx] == v {
		return idx
	}

	// Check subslices
	if a[idx] > v {
		// Item should be in a left slice
		return BinarySearch(a, start, idx, v)
	} else {
		if a[idx] < v {
			// Item should be in a right slice
			return BinarySearch(a, idx+1, stop, v)
		}
	}

	// No item found
	return -1
}
