package main

import (
	"fmt"
)

//var time int

func main() {
	var n int
	var tmp int
	a := []int{}

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		a = append(a, tmp)
	}

	fmt.Println(Check(a,0))
}

func Check(a []int,time int) int {
	for _, v := range a {
		if v%2 != 0 {
			return time // 奇数が見つかったら終了
		}
	}
	time++

	for i := range a {
		a[i] /= 2
	}

	return Check(a,time)
}

