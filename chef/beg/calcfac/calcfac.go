// https://www.codechef.com/problems/FCTRL2
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// Map from init val to calculated val
var facmap map[int]*big.Int = make(map[int]*big.Int)

// factorial calculations with memoization
func calcfac(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if val, ok := facmap[n]; ok {
		return val
	}
	c := calcfac(n - 1)
	calc := new(big.Int)
	calc = calc.Mul(c, big.NewInt(int64(n)))
	facmap[n] = calc
	return calc
}

func main() {
	var init int64
	fmt.Scanf("%d", &init)
	if init == 0 {
		panic("init val unset")
	}
	rdr := bufio.NewReader(os.Stdin)
	for s, err := rdr.ReadString('\n'); err == nil; s, err = rdr.ReadString('\n') {
		val, _ := strconv.Atoi(s[:len(s)-1])
		fmt.Printf("%v\n", calcfac(val).String())
	}
}
