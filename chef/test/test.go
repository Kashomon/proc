package main

import (
	"fmt"
)

func main() {
	for {
		var i int64
		fmt.Scanf("%d", &i)
		if i == 42 {
			break
		}
		fmt.Printf("%d\n", i)
	}
}
