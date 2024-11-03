package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// var time int
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 標準入力から1行を読み込み
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	//	fmt.Println(n)

	scanner.Scan()
	input := scanner.Text()

	// スペース区切りで文字列を分割
	parts := strings.Fields(input)

	// 結果を格納するスライス
	numbers := make([]int, 0, len(parts))

	// 文字列をintに変換してスライスに追加
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("変換エラー:", err)
			return
		}
		numbers = append(numbers, num)
	}

	if n < len(numbers) {
		numbers = numbers[:n]
	}

	fmt.Println(GamePlay(numbers))
}

func GamePlay(a []int) int {
	var alice, bob int
	//	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//	fmt.Println(a)
	for v, i := range a {
		if v%2 == 0 {
			alice = alice + i
		} else {
			bob = bob + i
		}
	}

	return alice - bob
}
