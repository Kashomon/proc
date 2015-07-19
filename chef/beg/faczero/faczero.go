package main

import (
	"fmt"
)

func numzeros(n int) int {
	zeroes := 0
	for v := 5; v <= n; v *= 5 {
		zeroes += n / v
	}
	return zeroes
}

func main() {
	var numvals int
	fmt.Scanf("%d", &numvals)
	if numvals == 0 {
		panic("no numvals")
	}
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		fmt.Printf("%d\n", numzeros(n))
	}
}
