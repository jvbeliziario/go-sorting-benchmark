package algorithms

func SelectionSort(arr []int) ([]int, SortingStats) {
	var stats SortingStats
	for i := range arr {
		intMin := i
		for j := i + 1; j < len(arr); j++ {
			stats.Comparisons++
			if arr[j] < arr[intMin] {
				intMin = j
			}
		}
		if i != intMin {
			stats.Swaps++
			arr[i], arr[intMin] = arr[intMin], arr[i]
		}
	}
	return arr, stats
}

/*
Total comparisons: n(n-1)/2
*/