package main

func FibonacciSequence(n int) map[int]int {
	fibonacciMap := make(map[int]int)

	a, b := 0, 1

	for i := 0; a <= n; i++ {
		fibonacciMap[i] = a
		a, b = b, a+b
	}

	return fibonacciMap
}
