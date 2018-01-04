package sort

//SelectionSort - selection sort
func SelectionSort(slice []int) []int {

	size := len(slice)

	sortedSlice := make([]int, size)

	for i := 0; i < size; i++ {
		minIndex := findMinIndex(slice)
		sortedSlice[i] = slice[minIndex]
		slice = append(slice[:minIndex], slice[minIndex+1:]...)
	}

	return sortedSlice
}

func findMinIndex(slice []int) int {
	minValue := slice[0]
	minIndex := 0

	for i := 1; i < len(slice); i++ {
		if slice[i] < minValue {
			minValue = slice[i]
			minIndex = i
		}
	}

	return minIndex
}
