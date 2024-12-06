package main

import (
	"sort"
	"strconv"
	"strings"
)

func process(input string) ([]int, []int){
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var left, right []int
	for _, v := range lines {
		coords := strings.Fields(v)
		l, _ := strconv.Atoi(coords[0])
		r, _ := strconv.Atoi(coords[1])
		left = append(left, l)
		right = append(right, r)
	}

    sort.Ints(right)
    sort.Ints(left)

    return left, right
}

func solutionOne(input string) int {
    left, right := process(input)

    total := 0

    for i := range(left){
        diff := right[i] - left[i]
        if diff < 0 {
            diff *= -1
        }
        total += diff
    }

	return total
}

func solutionTwo(input string) int {
    left, right := process(input)

    var total int
    
    for _, l := range(left){
        count := 0
        for _, r := range(right){
            if l == r {
                count += 1
            }
        }

        total += l * count
    }

    return total
}
