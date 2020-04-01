package main

import (
	"strconv"
	"strings"
)

func convertTerminalToInput(terminal string, size int) []int {
	result := make([]int, size)
	inputSlice := strings.Split(terminal, " ")
	for i, str := range inputSlice {
		n, _ := strconv.Atoi(str)
		result[i] = n
	}
	return result
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	s1, _ := reader.ReadString('\n')
//	s2, _ := reader.ReadString('\n')
//	s1 = s1[:len(s1)-1]
//	s2 = s2[:len(s2)-1]
//
//	size, _ := strconv.Atoi(s1)
//	array := convertTerminalToInput(s2, size)
//	minValue := getMinValue(array)
//	for _, v := range minValue {
//		fmt.Printf("%d %d\n", v[0],v[1])
//	}
//}

func getMinValue(array []int) [][2]int {
	res := make([][2]int, len(array))
	sk := make([]int, 0, 100)
	for i := 0; i < len(array); i++ {
		for len(sk) != 0 && array[sk[len(sk)-1]] > array[i] {
			left := -1
			if len(sk) > 1 {
				left = sk[len(sk)-2]
			}
			res[sk[len(sk)-1]] = [2]int{left, i}
			sk = sk[:len(sk)-1]
		}
		sk = append(sk, i)
	}
	for len(sk) != 0 {
		right := -1
		left := -1
		if len(sk) > 1 {
			left = sk[len(sk)-2]
		}
		res[sk[len(sk)-1]] = [2]int{left, right}
		sk = sk[:len(sk)-1]
	}
	return res
}
