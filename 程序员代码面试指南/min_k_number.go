package main

import (
	"fmt"
	"math"
)

func getMinKByBFPRT(arr []int, k int) int {
	des := make([]int, len(arr))
	copy(des, arr)
	return selectK(des, 0, len(arr)-1, k-1)
}

func selectK(arr []int, start int, end int, i int) int {
	if start == end {
		return arr[start]
	}
	midNumber := getMinNumberByArray(arr, start, end)
	rangeMid := getRangMidPosition(arr, start, end, midNumber)
	println("=========================")
	fmt.Println("arr = ", arr)
	fmt.Println("midNumber = ", midNumber)
	println("pivotRange = ", rangeMid[0], " ", rangeMid[1])
	println("i = ", i)
	if i >= rangeMid[0] && i <= rangeMid[1] {
		return arr[i]
	} else if i >= rangeMid[0] {
		return selectK(arr, rangeMid[1]+1, end, i)
	} else {
		return selectK(arr, start, rangeMid[0]-1, i)
	}
}

func getRangMidPosition(arr []int, start int, end int, midNumber int) [2]int {
	small := start
	big := end
	cur := start
	for cur != big+1 {
		if arr[cur] < midNumber {
			arr[cur], arr[small] = arr[small], arr[cur]
			cur++
			small++
		} else if arr[cur] > midNumber {
			arr[cur], arr[big] = arr[big], arr[cur]
			big--
		} else {
			cur++
		}
	}
	return [2]int{small, big}
}

func getMinNumberByArray(arr []int, begin, end int) int {
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
		midArray[i] = getMid(arr, s, int(math.Min(float64(e), float64(end))))
	}

	return selectK(midArray, 0, len(midArray)-1, len(midArray)/2)
}

func getMid(arr []int, s int, e int) int {
	insertSort(arr, s, e)
	sum := s + e
	mid := sum/2 + sum%2
	return arr[mid]
}

func insertSort(arr []int, s int, e int) {
	for i := s + 1; i != e+1; i++ {
		for j := i; j != s; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}
