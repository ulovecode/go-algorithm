package main

import (
	"math"
)

func getMinKByBFPRT(arr []int, k int) int {
	des := make([]int, len(arr))
	copy(des, arr)
	return selectK(des, 0, len(arr)-1, k-1)
}

func selectK(des []int, start int, end int, i int) int {
	if start == end {
		return des[start]
	}
	midNumber := getMinNumberByArray(des, start, end)
	rangeMid := getRangMidPosition(des, start, end, midNumber)
	println("pivotRange = ", rangeMid[0], " ", rangeMid[1])
	if i >= rangeMid[0] && i <= rangeMid[1] {
		return des[i]
	} else if i >= rangeMid[0] {
		return selectK(des, rangeMid[1]+1, end, i)
	} else {
		return selectK(des, start, rangeMid[0]-1, i)
	}
}

func getRangMidPosition(des []int, start int, end int, midNumber int) [2]int {
	small := start
	big := end
	cur := start
	for cur != big+1 {
		if des[cur] < midNumber {
			des[cur], des[small] = des[small], des[cur]
			cur++
			small++
		} else if des[cur] > midNumber {
			des[cur], des[big] = des[big], des[cur]
			big--
		} else {
			cur++
		}
	}
	return [2]int{small, big}

}

func getMinNumberByArray(des []int, begin, end int) int {
	var midArray []int
	num := end - begin + 1
	if num%5 == 0 {
		midArray = make([]int, num/5)
	} else {
		midArray = make([]int, num/5+1)
	}
	for i := 0; i < len(midArray); i++ {
		s := begin + i*5
		e := s + 4
		midArray[i] = getMid(des, s, int(math.Min(float64(e), float64(end))))
	}

	return selectK(midArray, 0, len(midArray)-1, len(midArray)/2)
}

func getMid(des []int, s int, e int) int {
	insertSort(des, s, e)
	sum := s + e
	mid := sum/2 + sum%2
	return des[mid]
}

func insertSort(des []int, s int, e int) {
	for i := s + 1; i != e+1; i++ {
		for j := i; j != s; j-- {
			if des[j-1] > des[j] {
				des[j-1], des[j] = des[j], des[j-1]
			} else {
				break
			}
		}
	}
}
