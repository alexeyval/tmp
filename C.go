package main

import (
	"bufio"
	"os"
)

func main() {
	//file, err := os.Open("/home/alexey/Documents/Go/route256" +
	//	"/Соревнование/now2/problem-b-tests/tests/02")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//reader := bufio.NewReader(file)
	reader := bufio.NewReader(os.Stdin)

	//var t int // 1 - 1000
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		_ = text
	}
}