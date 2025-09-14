package benchmark

import (
	"math/rand"
	"time"
)

// Helper function to generate random test data (fixed per benchmark run)
func generateRandomArray(size int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rng.Intn(1000)
	}
	return arr
}

// Helper function to generate sorted test data
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

// Helper function to generate reverse sorted test data
func generateReverseSortedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = size - i
	}
	return arr
}

// Helper function to generate nearly sorted test data with custom percentage
// percentSorted: 0.0 = completely random, 1.0 = completely sorted
func generateNearlySortedArray(size int, percentSorted float64) []int {
	arr := generateSortedArray(size)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Calculate how many elements to shuffle based on percentage
	elementsToShuffle := int(float64(size) * (1.0 - percentSorted))

	for i := 0; i < elementsToShuffle; i++ {
		a := rng.Intn(size)
		b := rng.Intn(size)
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr
}

// Helper function to generate array with many duplicates
func generateDuplicateArray(size int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	// Only use 10 different values to create many duplicates
	for i := range arr {
		arr[i] = rng.Intn(10)
	}
	return arr
}

// Helper function to copy array (avoids multiple allocations)
func copyArray(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	return copied
}
