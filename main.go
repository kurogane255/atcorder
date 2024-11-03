package main

import (
	"fmt"
)

func main() {
	var a, b, c, x, n int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	y := x / 50

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			for k := 0; k < a; k++ {
				if y == i+j*2+k*10 {
					n++
				}
			}
		}
	}

	fmt.Println(n)
}
