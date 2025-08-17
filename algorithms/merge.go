package algorithms

func MergeSort(arr []int) ([]int, SortingStats) {
	return mergeSortHelper(arr)
}

func mergeSortHelper(arr []int) ([]int, SortingStats) {
	if len(arr) <= 1 {
		return arr, SortingStats{}
	}

	mid := len(arr) / 2
	left, right := arr[:mid], arr[mid:]

	ch := make(chan struct {
		sorted []int
		stats  SortingStats
	}, 2)

	go func() {
		sorted, stats := mergeSortHelper(left)
		ch <- struct {
			sorted []int
			stats  SortingStats
		}{sorted, stats}
	}()

	go func() {
		sorted, stats := mergeSortHelper(right)
		ch <- struct {
			sorted []int
			stats  SortingStats
		}{sorted, stats}
	}()

	leftResult := <-ch
	rightResult := <-ch

	merged, mergeStats := merge(leftResult.sorted, rightResult.sorted)

	// Combine stats
	totalStats := SortingStats{
		Comparisons: leftResult.stats.Comparisons + rightResult.stats.Comparisons + mergeStats.Comparisons,
		Swaps:       leftResult.stats.Swaps + rightResult.stats.Swaps + mergeStats.Swaps,
	}

	return merged, totalStats
}

func merge(arrOne, arrTwo []int) ([]int, SortingStats) {
	var mergedArray []int
	var stats SortingStats
	i, j := 0, 0

	for i < len(arrOne) && j < len(arrTwo) {
		stats.Comparisons++
		if arrOne[i] <= arrTwo[j] {
			mergedArray = append(mergedArray, arrOne[i])
			i++
		} else {
			mergedArray = append(mergedArray, arrTwo[j])
			j++
		}
		stats.Swaps++ //appends
	}


    mergedArray = append(mergedArray, arrOne[i:]...)
	stats.Swaps += len(arrOne[i:])

	mergedArray = append(mergedArray, arrTwo[j:]...)
	stats.Swaps += len(arrTwo[j:])

	return mergedArray, stats
}
