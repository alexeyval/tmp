package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Person struct {
	t   int
	pos int
}

// ByTime implements sort.Interface for []Person based on
// the Age field.
type ByTime []Person

//type ByPos []Person

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].t < a[j].t }

func (a ByTime) String() string {
	s := ""
	place := 1
	var oldT int
	cash := 0
	tmp := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		switch i {
		case 0:
			tmp[i] = place
			//a[i].pos = place
			s += fmt.Sprint(place) + " "
			oldT = a[i].t
		default:
			if int(math.Abs(float64(a[i].t-oldT))) <= 1 {
				tmp[i] = place
				s += fmt.Sprint(place) + " "
				cash++
			} else {
				place = place + cash + 1
				tmp[i] = place
				s += fmt.Sprint(place) + " "
				cash = 0
			}
			oldT = a[i].t
		}
	}
	for i := range tmp {
		a[i].t = tmp[i]
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].pos < a[j].pos
	})
	for _, v := range a {
		fmt.Print(fmt.Sprint(v.t) + " ")
	}

	return ""
	//return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	//file, err := os.Open("/home/alexey/Documents/Go/route256/Соревнование/now2/problem-d-tests/tests/01")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//reader := bufio.NewReader(file)
	reader := bufio.NewReader(os.Stdin)

	var t int // 1 - 1000
	fmt.Fscan(reader, &t)

	for count := 0; count < t; count++ {
		var n int
		fmt.Fscan(reader, &n) // 2 - 10^5

		p := make([]Person, n)

		for i := 0; i < n; i++ {
			var t2 int
			fmt.Fscan(reader, &t2) // 2 - 10^9
			p[i].t = t2
			p[i].pos = i + 1

		}
		r := ByTime(p)
		sort.Sort(r)

		fmt.Println(r)
	}
}
