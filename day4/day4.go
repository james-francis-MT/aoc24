package day4

import (
	"regexp"
	"strings"
)

func solutionOne(input string) int {
	horizontalLines := strings.Split(input, "\n")
	var verticalLines [200]string
	var rightDiagonalLines [500]string
	var leftDiagonalLines [500]string

	xmas := regexp.MustCompile(`XMAS`)
	samx := regexp.MustCompile(`SAMX`)

	total := 0
	width := len(horizontalLines[0])
	for i, v := range horizontalLines {
		total += len(xmas.FindAllString(v, -1))
		total += len(samx.FindAllString(v, -1))
		for j, c := range v {
			verticalLines[j] += string(c)
			rightDiagonalLines[width+i-j] += string(c)
			leftDiagonalLines[i+j] += string(c)
		}
	}

	for _, v := range verticalLines {
		total += len(xmas.FindAllString(v, -1))
		total += len(samx.FindAllString(v, -1))
	}
	for _, v := range rightDiagonalLines {
		total += len(xmas.FindAllString(v, -1))
		total += len(samx.FindAllString(v, -1))
	}
	for _, v := range leftDiagonalLines {
		total += len(xmas.FindAllString(v, -1))
		total += len(samx.FindAllString(v, -1))
	}

	return total
}

func solutionTwo(input string) int {
	hl := strings.Split(strings.TrimSpace(input), "\n")

	total := 0
    width := len(hl[0])
    height := len(hl)
	for i, v := range hl {
		for j, c := range v {
			if c == 'A' {
				if i-1 < 0 || j-1 < 0 || i+1 >= height || j+1 >= width {
					continue
				}
				upperLeft := hl[i-1][j-1]
				lowerRight := hl[i+1][j+1]
				upperRight := hl[i-1][j+1]
				lowerLeft := hl[i+1][j-1]
				if isMOrS(upperLeft) && isMOrS(lowerRight) && upperLeft != lowerRight {
					if isMOrS(upperRight) && isMOrS(lowerLeft) && upperRight != lowerLeft {
						total += 1
					}
				}
			}
		}
	}

	return total
}

func isMOrS(c byte) bool {
	return c == 'M' || c == 'S'
}
