package algorithms

func HeapSort(arr []int) ([]int, SortingStats) {
	var stats SortingStats
	n := len(arr)
	
	if n <= 1 {
		return arr, stats
	}
	
	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, &stats)
	}
	
	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]
		stats.Swaps++
		
		// Call heapify on the reduced heap
		heapify(arr, i, 0, &stats)
	}
	
	return arr, stats
}

func heapify(arr []int, n, i int, stats *SortingStats) {
	largest := i  // Initialize largest as root
	left := 2*i + 1   // left child
	right := 2*i + 2  // right child
	
	// If left child is larger than root
	if left < n {
		stats.Comparisons++
		if arr[left] > arr[largest] {
			largest = left
		}
	}
	
	// If right child is larger than largest so far
	if right < n {
		stats.Comparisons++
		if arr[right] > arr[largest] {
			largest = right
		}
	}
	
	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		stats.Swaps++
		
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest, stats)
	}
}