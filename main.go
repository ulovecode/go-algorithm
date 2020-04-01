package main

func main() {
	n := 10
	// 1010  1001
	// 1000
	a := n & (n - 1)
	// 1010  0110
	// 0010
	b := n & (^n + 1)
	println(^n + 1)
	println(a)
	println(b)
}
