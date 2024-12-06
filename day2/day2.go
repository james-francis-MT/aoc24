package day2

import (
	"strings"

	"day2/report"
)

func solutionOne(input string) int {
	reports := buildReports(input)

	var total int

	for _, report := range reports {
		if report.IsValid() {
			total += 1
		}
	}

	return total
}

func solutionTwo(input string) int {
	reports := buildReports(input)

	var total int

	for _, report := range reports {
		if reportSubsetsAreValid(report) {
			total += 1
		}
	}

	return total
}

func reportSubsetsAreValid(r report.Report) bool {
	subsets := r.BuildSubsets()
	for _, subset := range subsets {
		if subset.IsValid() {
			return true
		}
	}
	return false
}


func buildReports(input string) []report.Report {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var reports []report.Report

	for _, line := range lines {
		reports = append(reports, report.BuildReport(line))
	}
	return reports
}

