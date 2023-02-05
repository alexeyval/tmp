package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pages map[int]bool

func (p Pages) String() string {
	keys := make([]int, 0, len(p))
	for id := range p {
		keys = append(keys, id)
	}
	sort.Ints(keys)
	var s []string
	for i := 0; i < len(keys); i++ {
		begin := i + 1
		for begin < len(keys) &&
			keys[begin]-keys[begin-1] == 1 {
			begin++
		}
		if begin == i+1 {
			s = append(s, fmt.Sprint(keys[i]))
		} else {
			s = append(s, fmt.Sprintf("%v-%v", keys[i], keys[begin-1]))
		}
		i = begin - 1
	}

	return strings.Join(s, ",")
}

func main() {
	//file, err := os.Open(
	//	"/home/alexey/Documents" +
	//		"/Go/route256/Соревнование/now2/problem-f-tests/tests/01")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//reader := bufio.NewReader(file)
	reader := bufio.NewReader(os.Stdin)

	var t int // 1 - 100
	fmt.Fscan(reader, &t)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for count := 0; count < t; count++ {
		var k int
		fmt.Fscan(reader, &k)
		var k2 string
		fmt.Fscan(reader, &k2)

		split := strings.Split(k2, ",")
		a := map[int]bool{}
		for i := 1; i <= k; i++ {
			a[i] = true
		}
		//fmt.Println(split)
		for _, v := range split {
			setSplit(v, a)
		}
		fmt.Println(Pages(a))
	}

}

func setSplit(s string, a map[int]bool) {
	switch {
	case strings.Contains(s, "-"):
		split := strings.Split(s, "-")
		b, _ := strconv.Atoi(split[0])
		e, _ := strconv.Atoi(split[1])

		for i := b; i <= e; i++ {
			delete(a, i)
		}
	default:
		n, _ := strconv.Atoi(s)
		delete(a, n)
	}
}
