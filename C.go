package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	//file, err := os.Open("/home/alexey/Documents/Go/route256/Соревнование/now2/c.test")
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

	reg, _ := regexp.Compile(`([A-Z]\d\d[A-Z][A-Z])|[A-Z]\d[A-Z][A-Z]`)
	for scanner.Scan() {
		text := scanner.Text()
		size := 0
		findAll := reg.FindAllString(text, -1)
		for _, v := range findAll {
			size += len(v)
		}
		if len(text) != size {
			fmt.Println("-")
		} else {
			for _, v := range findAll {
				fmt.Print(v + " ")
			}
			fmt.Println()
		}
	}
}
