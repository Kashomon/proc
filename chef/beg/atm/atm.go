package main

import (
	"fmt"
)

func scanner() (float64, error) {
	var amt int64
	fmt.Scanf("%d", &amt)
	var bal float64
	fmt.Scanf("%f", &bal)

	if amt == 0 || bal == 0 {
		return 0.0, fmt.Errorf("Err: incorrect vals for amt %v, bal %v", amt, bal)
	}

	if amt%5 != 0 {
		return bal, nil
	}
	c := bal - float64(amt) - 0.5
	if c < 0 {
		return bal, nil
	}
	return c, nil
}

func main() {
	for {
		v, err := scanner()
		if err != nil {
			break
		}
		fmt.Printf("%.2f\n", v)
	}
}
