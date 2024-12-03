package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")

	reports := make([][]int, len(lines))
	var safeCount int
	for i, l := range lines {
		ele := strings.Fields(l)
		if len(ele) == 0 {
			break
		}
		reports[i] = make([]int, len(ele))

		for j := 0; j < len(ele); j++ {
			reports[i][j], _ = strconv.Atoi(ele[j])

		}
		if r := isLevelSafe(reports[i], len(reports[i])); r {
			safeCount++
		}
	}
	fmt.Println("1st Answer: ", safeCount)

	safeCount = 0
	var r bool
	for i := range reports {

		newReport := make([]int, len(reports[i])-1)

		var l int
		for j := 0; j < len(reports[i]); j++ {
			//newReport = append(reports[i][:j], reports[i][j+1:]...)
			l = 0
			for k := 0; k < len(reports[i]); k++ {
				if k == j {
					continue
				}
				newReport[l] = reports[i][k]
				l++
			}
			if r = isLevelSafe(newReport, len(newReport)); r {
				safeCount++
				break
			}
		}
	}
	fmt.Println("2st Answer: ", safeCount)
	fmt.Println("vim-go")
}

func isLevelSafe(a []int, l int) bool {
	var isSafe bool = true
	var direction bool = true
	for i := 1; i < l; i++ {
		d, dir := diff(a[i], a[i-1])
		if i == 1 {
			direction = dir
		}
		if isSafe && (d == 0 || d > 3 || direction != dir) {
			isSafe = false
		}
	}
	return isSafe
}

func diff(i, j int) (int, bool) {
	if i > j {
		return i - j, true
	}
	return j - i, false
}
