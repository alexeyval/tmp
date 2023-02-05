package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//start := time.Now()
	file, err := os.Open("/home/alexey/Documents/Go/route256" +
		"/Соревнование/now2/problem-b-tests/tests/01")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	//reader := bufio.NewReader(os.Stdin)

	var t int // 1 - 1000
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		const N = 10
		var a [N]int // values 1 - 4

		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &a[j])
		}
		fmt.Println(a)
	}

}
