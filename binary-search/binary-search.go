package biarySearch

//Search - binary search in int slice
func Search(slice []int, value int) (index int, exist bool) {
	low := 0
	hight := len(slice) - 1

	for low <= hight {
		middle := (low + hight) / 2

		if slice[middle] == value {
			index = middle
			exist = true
			break
		} else if slice[middle] > value {
			hight = middle - 1
		} else {
			low = middle + 1
		}
	}

	return
}
