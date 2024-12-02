package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./input1.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		ele := strings.Fields(line)
		if len(ele) == 0 {
			continue
		}
		list1[i], _ = strconv.Atoi(ele[0])
		list2[i], _ = strconv.Atoi(ele[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var d int = 0
	for i := range list1 {
		d += diff(list1[i], list2[i])
	}

	fmt.Println("1st answer:", d)

	var lastNo, lastCount, c, l, r, ss int
	for l < 1000 {
		if lastNo == list1[l] {
			ss += lastCount
			l++
			continue
		}
		lastNo = 0
		lastCount = 0
		if list1[l] < list2[r] {
			l++
			continue
		}

		if list1[l] != list2[r] {
			r++
			continue
		}

		c = 0
		for c = 0; ; {
			if list1[l] == list2[r+c] {
				c++
				continue
			}
			break
		}
		lastNo = list1[l]
		lastCount = list1[l] * c
		ss += lastCount
		l++
		r += c
		if l >= len(list1) {
			break
		}
	}
	fmt.Println("2nd answer:", ss)
	fmt.Println("vim-go")
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
