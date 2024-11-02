package main

import (
	"fmt"
)

func getDigits(n int) int {
	var digits int
//	fmt.Println("check n",n)
	for n > 0 {
		digit := n % 10
		digits += digit
		n /= 10
	}
//	fmt.Println("check",digits)
	return digits
}
func main() {
	var a,b int
	var n int
	var sum int

	fmt.Scanf("%d %d %d\n",&n,&a,&b)

	for i := 1; i <= n ; i++ {
		digits := getDigits(i)
//		fmt.Println(digits,n,i)
		if a <= digits &&  digits <= b {
			sum += i	   
		}
	}
	fmt.Println(sum)
}
