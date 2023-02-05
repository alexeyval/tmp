package main

import (
	"bufio"
	"fmt"
	"os"
)

type Fail bool

func (fail Fail) String() string { return map[Fail]string{true: "NO", false: "YES"}[fail] }

func main() {
	//start := time.Now()
	//file, err := os.Open("/home/alexey/Documents/Go/route256" +
	//	"/Соревнование/now2/problem-b-tests/tests/01")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//reader := bufio.NewReader(file)
	reader := bufio.NewReader(os.Stdin)

	var t int // 1 - 1000
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		const N = 10
		const K = 4
		var counter [K]int
		results := [K]int{4, 3, 2, 1}

		for j := 0; j < N; j++ {
			var a int // 1 - 4
			fmt.Fscan(reader, &a)
			counter[a-1]++
		}
		fail := false
		for j := 0; j < K; j++ {
			fail = fail || (counter[j] != results[j])
		}
		fmt.Println(Fail(fail))
	}

}
