package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(
		"/home/alexey/Documents" +
			"/Go/route256/Соревнование/now2/problem-g-tests/tests/01")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	//reader := bufio.NewReader(os.Stdin)

	var t int // 1 - 1000
	fmt.Fscan(reader, &t)

	for count := 0; count < t; count++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		for i := 0; i < m; i++ {
			var w int
			fmt.Fscan(reader, &w)
		}
	}
}
