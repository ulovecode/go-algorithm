package main

func numTeams(rating []int) int {
	ratingLen := len(rating)
	if ratingLen < 3 {
		return 0
	}
	ans := 0
	for i := 0; i < ratingLen; i++ {
		for j := 0; j < ratingLen; j++ {
			for k := 0; k < ratingLen; k++ {
				if i < j && j < k && ((rating[i] < rating[j] && rating[j] < rating[k]) ||
					rating[i] > rating[j] && rating[j] > rating[k]) {
					ans++
				}
			}
		}
	}

	return ans
}

func main() {
	print(numTeams([]int{2, 5, 3, 4, 1}))
}
