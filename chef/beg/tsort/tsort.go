package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var init int
	fmt.Scanf("%d", &init)
	if init == 0 {
		panic("bad init value")
	}

	var ints [1000000]int
	r := bufio.NewReader(os.Stdin)
	for l, err := r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		n, _ := strconv.Atoi(l[:len(l)-1])
		ints[n]++
	}

	outintz := make([]string, init, init)
	idx := 0
	for i, v := range ints {
		if v > 0 {
			for j := 0; j < v; j++ {
				outintz[idx] = strconv.Itoa(i)
				idx++
			}
		}
	}
	strOut := strings.Join(outintz, "\n")
	fmt.Printf("%s", strOut)
}
