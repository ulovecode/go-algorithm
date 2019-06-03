package sorts

func InsertionSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	var current int
	for i := 0; i < len(arr)-1; i++ {
		current = arr[i+1]
		preIndex := i
		for ; preIndex >= 0 && current < arr[preIndex]; preIndex-- {
			arr[preIndex+1] = arr[preIndex]
		}
		arr[preIndex+1] = current
	}

}
