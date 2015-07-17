package main

import (
	"fmt"
	"github.com/Kashomon/gcomp"
	"strings"
)

func main() {
	for {
		var l string
		fmt.Scanln("%s", &l)
		fmt.Printf("Zed:%v\n", l)
		splat := strings.Split(l, " ")

		iv := gcomp.Stoi(splat[0])
		ft := gcomp.Stof(splat[1])
		fmt.Printf("%v, %v\n", iv, ft)
	}
}
