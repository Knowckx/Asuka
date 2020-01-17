package math

import "fmt"

func GenFibonacci(n int) int {
	if n < 2 {
		return n
	}

	l1 := 0
	l2 := 1
	for i := 2; i <= n; i++ {
		l1, l2 = l2, l1+l2
		fmt.Println(i, l1, l2)
	}
	return l2
}
