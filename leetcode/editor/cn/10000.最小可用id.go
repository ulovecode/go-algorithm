package leetcode

func main() {

}

func min_free(arr []int, l, r, n int) int {
	if n == 0 {
		return l
	}
	m := (l + r) / 2
	left, right := 0, 0
	for ; right < n; right++ {
		if arr[right] <= m {
			arr[right], arr[left] = arr[left], arr[right]
			left++
		}
	}
	if left == m-l+1 {
		arr = arr[left:]
		n -= left
		l = m + 1
	} else {
		n = left
		r = m
	}

	return min_free(arr, l, r, n)
}
func min_free1(arr []int, n int) int {
	l := 0
	r := n - 1
	for n != 0 {
		m := (l + r) / 2
		left, right := 0, 0
		for ; right < n; right++ {
			if arr[right] <= m {
				arr[right], arr[left] = arr[left], arr[right]
				left++
			}
		}
		if left == m-l+1 {
			arr = arr[left:]
			n -= left
			l = m + 1
		} else {
			n = left
			r = m
		}

	}
	return l
}
