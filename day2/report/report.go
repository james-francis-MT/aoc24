package report

import (
	"slices"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}

func BuildReport(line string) Report {
	levelsAsString := strings.Fields(line)

	var levelsAsInt []int
	for _, v := range levelsAsString {
		converted, _ := strconv.Atoi(v)
		levelsAsInt = append(levelsAsInt, converted)
	}

	return Report{Levels: levelsAsInt}
}

func (r Report) BuildSubsets() []Report {
	var subsets []Report

	for i := range r.Levels {
		newReport := Report{Levels: slices.Concat(r.Levels[:i], r.Levels[i+1:])}
		subsets = append(subsets, newReport)
	}

	return subsets
}

func (r Report) IsValid() bool {
	if r.isAscending() || r.isDescending() {
		return r.areDiffsSmall()
	}
	return false
}

func (r Report) isAscending() bool {
	return slices.IsSorted(r.Levels)
}

func (r Report) isDescending() bool {
	return slices.IsSortedFunc(r.Levels, func(a int, b int) int { return b - a })
}

func (r Report) areDiffsSmall() bool {
	for i := 0; i < len(r.Levels)-1; i++ {
		diff := r.Levels[i+1] - r.Levels[i]

		if diff < 0 {
			diff *= -1
		}

		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}
