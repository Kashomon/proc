// https://www.codechef.com/problems/MUFFINS3
// tc: 1000
// ex-size: 100,000,000
package main

import (
	"fmt"
)

func main() {
	var init int
	fmt.Scanf("%d", &init)
	if init == 0 {
		panic("bad init value")
	}
	for {
		var num int
		fmt.Scanf("%d", &num)
		if num == 0 {
			break
		}
		if num == 1 || num == 2 {
			fmt.Printf("%d\n", num)
		} else {
			fmt.Printf("%d\n", num/2+1)
		}
	}
}
