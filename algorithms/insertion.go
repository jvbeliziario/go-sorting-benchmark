package algorithms

func InsertionSort(arr []int) ([]int, SortingStats) {
	var stats SortingStats
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			stats.Comparisons++
			if arr[j-1] > arr[j] {
				stats.Swaps++
				arr[j-1], arr[j] = arr[j], arr[j-1]
				j--
			} else {
				break
			}
		}
	}
	return arr, stats
}