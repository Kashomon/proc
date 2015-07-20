package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func relax(prev, next []int) []int {
	out := make([]int, len(next))
	for i, val := range next {
		var leftsum, rightsum int
		if i > 0 {
			leftsum = prev[i-1] + val
		}
		if i < len(prev) {
			rightsum = prev[i] + val
		}
		if leftsum > rightsum {
			out[i] = leftsum
		} else {
			out[i] = rightsum
		}
	}
	return out
}

func printmax(s []int) {
	max := 0
	for _, val := range s {
		if val > max {
			max = val
		}
	}
	fmt.Printf("%v\n", max)
}

func main() {
	var init int64
	fmt.Scanf("%d", &init)
	if init == 0 {
		panic("init val unset")
	}

	r := bufio.NewReader(os.Stdin)
	lines := 0
	lineNum := 0
	var prevsums []int
	for l, err := r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		l = l[:len(l)-1]
		if lines == 0 {
			lines, _ = strconv.Atoi(l)
		} else {
			nums := strings.Split(l, " ")
			curline := make([]int, len(nums))
			for i, val := range nums {
				n, _ := strconv.Atoi(val)
				curline[i] = n
			}
			if len(prevsums) == 0 {
				prevsums = curline
			} else {
				prevsums = relax(prevsums, curline)
			}
			lineNum++
		}
		if lineNum == lines {
			printmax(prevsums)
			lineNum = 0
			lines = 0
			prevsums = *new([]int)
		}
	}
}
