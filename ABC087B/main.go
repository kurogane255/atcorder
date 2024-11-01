package main

import (
        "fmt"
)

func main() {
        var a, b, c, x, n int

        fmt.Scanf("%d\n", &a)
        fmt.Scanf("%d\n", &b)
        fmt.Scanf("%d\n", &c)
        fmt.Scanf("%d\n", &x)

        y := x / 50

        for i := 0; i <= c; i++ {
                for j := 0; j <= b; j++ {
                        for k := 0; k <= a; k++ {
                                //                                      fmt.Println(y,i,j,k)
                                if y == i+j*2+k*10 && i+j+k != 0 {
                                        n++
                                }
                        }
                }
        }

        fmt.Println(n)
}
