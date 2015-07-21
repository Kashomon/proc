package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	isAmbig  = "ambiguous"
	notAmbig = "not ambiguous"
)

var invArr [10000001]int
var initArr [10000001]int

func testcase(size int, r *bufio.Reader) {
	pos := 1
	buf := make([]byte, size, size)
	buflen := 0

	for b, err := r.ReadByte(); err == nil; b, err = r.ReadByte() {
		if b == byte(' ') || b == byte('\n') {
			num, err := strconv.Atoi(string(buf[:buflen]))
			if err != nil {
				panic(err)
			}
			initArr[pos] = num
			invArr[num] = pos

			pos++
			buflen = 0
			if b == byte('\n') {
				break
			}
		} else {
			buf[buflen] = b
			buflen++
		}
	}
	ambig := true
	for i, v := range initArr[1 : size+1] {
		if i == v || initArr[i] == invArr[i] {
			// still ambiguous
		} else {
			ambig = false
			break
		}
	}
	if ambig {
		fmt.Printf("%v\n", isAmbig)
	} else {
		fmt.Printf("%v\n", notAmbig)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for s, err := r.ReadString('\n'); err == nil; s, err = r.ReadString('\n') {
		s = s[:len(s)-1]
		size, _ := strconv.Atoi(s)
		if size == 0 {
			break
		}
		testcase(size, r)
	}
}
