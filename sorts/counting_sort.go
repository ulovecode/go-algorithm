package sorts

func CountingSort(arr []int) (result []int) {
	if len(arr) <= 1 {
		result = arr
		return
	}
	var max = -int(^uint32(0) >> 1)
	var min = int(^uint32(0) >> 1)
	result = make([]int, len(arr))
	for _, v := range arr {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	length := max - min
	m := make([]int, length+1)
	for _, v := range arr {
		m[v-min]++
	}
	i := 0
	for k, v := range m {
		for v != 0 {
			result[i] = k + min
			i++
			v--
		}

	}

	return
}
