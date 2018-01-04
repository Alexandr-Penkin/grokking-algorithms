package search

//Search - Typical range search in int slice
func Search(slice []int, value int) (index int, exist bool) {

	for i, v := range slice {
		if v == value {
			index = i
			exist = true
			break
		}
	}

	return
}
