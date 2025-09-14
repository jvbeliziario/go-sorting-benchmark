package benchmark

import (
	"fmt"
	"sync"
	"testing"

	"github.com/guiwoch/go-sorting-benchmark/algorithms"
)

// Standardized test sizes for all algorithms
//var standardSizes = []int{100, 500, 1000, 2500, 5000, 10000, 25000, 50000, 70000, 100000, }

var standardSizes = []int{
    100, 500, 1000, 2500, 5000, 10000, 25000, 50000, 100000, 250000, 500000,
    1000000, 2500000, 5000000, 10000000, 25000000, // Stop here for safety
}

// Sorting algorithms
var testConfigs = map[string]func([]int) ([]int, algorithms.SortingStats){
	"QuickSort":     algorithms.QuickSort,
	"MergeSort":     algorithms.MergeSort,
	"InsertionSort": algorithms.InsertionSort,
	"SelectionSort": algorithms.SelectionSort,
	"HeapSort":      algorithms.HeapSort,
}

// Dataset generators
var datasetGenerators = map[string]func(int) []int{
	"Random":        generateRandomArray,
	"Sorted":        generateSortedArray,
	"ReverseSorted": generateReverseSortedArray,
	"Duplicate":     generateDuplicateArray,
}

// Accumulated stats
var (
	statsMap = make(map[string]algorithms.SortingStats)
	statsMux sync.Mutex
)

func BenchmarkSortingAlgorithms(b *testing.B) {
	for algName, sortFunc := range testConfigs {
		for datasetName, generator := range datasetGenerators {
			for _, size := range standardSizes {
				b.Run(fmt.Sprintf("%s-%s-%d", algName, datasetName, size), func(b *testing.B) {
					data := generator(size)
					var totalComparisons, totalSwaps int64

					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						testData := copyArray(data)
						_, stats := sortFunc(testData)
						totalComparisons += int64(stats.Comparisons)
						totalSwaps += int64(stats.Swaps)
					}

					// Calculate average per iteration and add to map
					avgComparisons := int(totalComparisons) / b.N
					avgSwaps := int(totalSwaps) / b.N

					key := fmt.Sprintf("%s-%s-%d", algName, datasetName, size)
					statsMux.Lock()
					statsMap[key] = algorithms.SortingStats{
						Comparisons: avgComparisons,
						Swaps:       avgSwaps,
					}
					statsMux.Unlock()
				})
			}
		}
	}

	// Print stats after all benchmarks complete
	printStats()
}

// Helper function to print stats
func printStats() {
	if len(statsMap) == 0 {
		fmt.Println("No stats collected.")
		return
	}

	fmt.Println("\nSorting Algorithm Statistics:")
	fmt.Println("========================================")

	for _, alg := range []string{"QuickSort", "MergeSort", "InsertionSort", "SelectionSort"} {
		fmt.Printf("== %s ==\n", alg)
		for _, dataset := range []string{"Random", "Sorted", "ReverseSorted", "Duplicate"} {
			fmt.Printf("  -- %s --\n", dataset)
			for _, size := range standardSizes {
				key := fmt.Sprintf("%s-%s-%d", alg, dataset, size)
				if stats, exists := statsMap[key]; exists {
					fmt.Printf("    %d: Comparisons: %d, Swaps: %d\n",
						size, stats.Comparisons, stats.Swaps)
				}
			}
		}
		fmt.Println()
	}
}

// Print all accumulated stats
func TestPrintStats(t *testing.T) {
	if len(statsMap) == 0 {
		t.Log("No stats collected. Run benchmarks first with: go test -bench=.")
		return
	}

	fmt.Println("\nSorting Algorithm Statistics (Flat List):")
	fmt.Println("========================================")
	for key, stats := range statsMap {
		fmt.Printf("%s: Comparisons: %d, Swaps: %d\n", key, stats.Comparisons, stats.Swaps)
	}
}
