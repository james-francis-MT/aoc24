package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func solutionOne(input string) int {
	reg, _ := regexp.Compile(`mul\(\d*,\d*\)`)
	digitsRE, _ := regexp.Compile(`\((\d*),(\d*)\)`)
	muls := reg.FindAllString(input, -1)

	total := 0
	for _, mul := range muls {
		match := digitsRE.FindAllStringSubmatch(mul, -1)
		d1, _ := strconv.Atoi(match[0][1])
		d2, _ := strconv.Atoi(match[0][2])
		total += d1 * d2
	}

	return total
}

func solutionTwo(input string) int {
	input = strings.TrimSpace(input)
	reg, _ := regexp.Compile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)
	digitsRE, _ := regexp.Compile(`\((\d*),(\d*)\)`)
	muls := reg.FindAllString(input, -1)

	total := 0
	compute := true
	for _, mul := range muls {
		if mul == "don't()" {
			compute = false
		}
		if mul == "do()" {
			compute = true
            continue
		}
		if compute {
			match := digitsRE.FindAllStringSubmatch(mul, -1)
			d1, _ := strconv.Atoi(match[0][1])
			d2, _ := strconv.Atoi(match[0][2])
			total += d1 * d2
		}
	}

	return total
}
