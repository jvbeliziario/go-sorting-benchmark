package algorithms

func MergeSort(arr *[]int) {
    *arr = mergeSortHelper(*arr)
}

func mergeSortHelper(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    mid := len(arr) / 2
    left, right := arr[:mid], arr[mid:]
    
    ch := make(chan []int, 2)
    
    go func() { ch <- mergeSortHelper(left) }()
    go func() { ch <- mergeSortHelper(right) }()
    
    sortedOne := <-ch
    sortedTwo := <-ch
    
    return merge(sortedOne, sortedTwo)
}

func merge(arrOne, arrTwo []int) []int{
	var mergedArray []int
	
	for range (len(arrOne)+len(arrTwo)){
		if arrOne[0] > arrTwo[0]{
			mergedArray = append(mergedArray, arrOne[0])
			arrOne = arrOne[1:] 
		} else {
			mergedArray = append(mergedArray, arrTwo[0])
			arrTwo = arrTwo[1:]
		}
	}
	
	
	return mergedArray
}
