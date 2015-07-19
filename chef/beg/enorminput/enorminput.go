package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nums int64
	fmt.Scanf("%d", &nums)
	if nums == 0 {
		panic("No max nums")
	}

	var divisor int64
	fmt.Scanf("%d", &divisor)
	if divisor == 0 {
		panic("Invalid divisor")
	}

	r := bufio.NewReader(os.Stdin)
	total := 0

	for s, err := r.ReadString('\n'); err == nil; s, err = r.ReadString('\n') {
		num, _ := strconv.Atoi(s[:len(s)-1])
		if int64(num)%divisor == 0 {
			total++
		}
	}
	fmt.Printf("%d\n", total)
}
