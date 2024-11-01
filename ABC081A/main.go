package main

import (
	"fmt"
)

func main() {
	var n int
	var s[3] int

   	fmt.Scanf("%1d%1d%1d",&s[0],&s[1],&s[2])
	for i := 0; i < 3; i++ {
		if s[i] == 1 {
			n++	
		}
	}
	fmt.Println(n)
}


