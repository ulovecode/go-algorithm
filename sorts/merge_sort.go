package sorts

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[mid:])
	right := MergeSort(arr[:mid])
	return Merge(left, right)
}

func Merge(left []int, right []int) (result []int) {
	for j, i := 0, 0; i < len(left) || j < len(right); {
		if i == len(left) {
			result = append(result, right[j])
		}
		if j == len(right) {
			result = append(result, left[i])
		}
		if left[i] >= right[j] {
			result = append(result, right[j])
		}
		if left[i] < right[j] {
			result = append(result, left[i])
		}
	}
	return
}