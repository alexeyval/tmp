package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(
		"/home/alexey/Documents" +
			"/Go/route256/Соревнование/now2/problem-i-tests/tests/01")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	//reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var t int // 1 - 1000
	scanner.Scan()
	t, _ = strconv.Atoi(scanner.Text())

	for count := 0; count < t; count++ {
		var k int
		scanner.Scan()
		k, _ = strconv.Atoi(scanner.Text())

		table := ""
		for i := 0; i < k; i++ {
			scanner.Scan()
			table += strings.ReplaceAll(scanner.Text(), " ", "")
		}
		fmt.Println(table)
		fmt.Println()
	}
}
