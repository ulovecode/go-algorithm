package sorts

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1
	for head < tail {
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			i++
			head++
		}
	}
	arr[head] = mid
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
}
