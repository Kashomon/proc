// https://www.codechef.com/problems/TLG/
package main

import "fmt"

func main() {
	var init int
	fmt.Scanf("%d", &init)
	if init == 0 {
		panic("bad init value")
	}

	p1score := 0
	p2score := 0
	maxlead := 0
	maxpl := 1
	for {
		var p1s, p2s int
		fmt.Scanf("%d %d", &p1s, &p2s)
		if p1s == 0 && p2s == 0 {
			break
		}
		p1score += p1s
		p2score += p2s
		delta := p1score - p2score

		leadpl := 1
		if delta < 0 {
			leadpl = 2
			delta *= -1
		}
		if delta > maxlead {
			maxlead = delta
			maxpl = leadpl
		}
	}
	fmt.Printf("%d %d", maxpl, maxlead)
}
