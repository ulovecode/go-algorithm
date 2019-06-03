package sorts

func SelectSort(arr []int) {
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		pIndex := i
		curV := arr[pIndex]
		for j := i; j < arrLength; j++ {
			if curV > arr[j] {
				pIndex = j
				curV = arr[j]
			}
		}
		temp := arr[i]
		arr[i] = arr[pIndex]
		arr[pIndex] = temp
	}
}
