package sorts

func HeapSort(arr []int) {
	for i := range arr {
		heapInsert(arr, i)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, i int, size int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	for left < size {
		if arr[i] < arr[left] {
			largest = left
		}
		if right < size && arr[largest] < arr[right] {
			largest = right
		}
		if largest != i {
			arr[largest], arr[i] = arr[i], arr[largest]
		} else {
			break
		}
		i = largest
		left = i*2 + 1
		right = i*2 + 2
	}
}

func heapInsert(arr []int, i int) {
	parent := 0
	for i != 0 {
		parent = (i - 1) / 2
		if arr[parent] < arr[i] {
			arr[i], arr[parent] = arr[parent], arr[i]
		}
		i = parent
	}
}
