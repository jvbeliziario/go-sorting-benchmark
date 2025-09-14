package algorithms

func QuickSort(arr []int) ([]int, SortingStats) {
	if len(arr) <= 1 {
		return arr, SortingStats{}
	}
	
	var stats SortingStats
	quickSortHelper(arr, 0, len(arr)-1, &stats)
	return arr, stats
}

func quickSortHelper(arr []int, start, end int, stats *SortingStats) {
	if start >= end {
		return
	}
	
	// Partition
	pivotIndex := end
	i := start - 1
	
	for j := start; j < end; j++ {
		stats.Comparisons++
		if arr[j] < arr[pivotIndex] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			stats.Swaps++
		}
	}
	
	arr[i+1], arr[pivotIndex] = arr[pivotIndex], arr[i+1]
	stats.Swaps++
	finalPivotPos := i + 1
	
	// Recursive calls
	quickSortHelper(arr, start, finalPivotPos-1, stats)
	quickSortHelper(arr, finalPivotPos+1, end, stats)
}