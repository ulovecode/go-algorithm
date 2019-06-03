package sorts

func BubbleSort(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr)[i] > (arr)[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
