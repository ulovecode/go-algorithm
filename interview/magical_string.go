package interview

import (
	"container/list"
)

type sNum struct {
	value int
	num   int
}

func magicalString(n int) int {
	res := 0
	l := list.New()
	l.PushBack(&sNum{
		value: 1,
		num:   1,
	})
	next, count := 1, 1
	preSum := 0
	for count <= n || l.Front().Value.(*sNum).value == 1 {
		element := l.Front()
		l.Remove(element)
		sVal := element.Value.(*sNum)
		if sVal.value == 1 {
			if count == n {
				res += 1
				break
			}
			if count > n {
				res += n - (count - preSum)
				break
			}
			if sVal.num+count > n {
				res += n - count
			}
			res += sVal.num
		}
		next = getNextC(next)
		l.PushBack(&sNum{
			value: next,
			num:   sVal.value,
		})
		preSum = sVal.num * sVal.value
		count += preSum
	}
	return res
}

func getNextC(c int) int {
	if c == 1 {
		return 2
	}
	return 1
}
