package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, w int
	fmt.Scanln(&n, &w)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var input []int
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	ans := generateWindowMaxArray(input, w)
	printAns(ans)
}

func printAns(ans []int) {
	for _, a := range ans {
		fmt.Print(a, " ")
	}
}

func generateWindowMaxArray(input []int, size int) []int {
	res := []int{}
	q := make([]int, 0)
	for i := range input {
		for len(q) != 0 && input[q[len(q)-1]] < input[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= size-1 {
			res = append(res, input[q[0]])
		}
		if len(q) != 0 && q[0] == i-size+1 {
			q = q[1:]
		}
	}
	return res
}
