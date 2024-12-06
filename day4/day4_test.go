package day4

import (
	"os"
	"regexp"
	"testing"
)

func TestTestInput(t *testing.T) {
	var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := solutionOne(input)
	expected := 18

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestBlah(t *testing.T) {
	xmas := regexp.MustCompile(`XMAS`)
	samx := regexp.MustCompile(`SAMX`)
    test := len(xmas.FindAllString("XMASAMX", -1))
    test += len(samx.FindAllString("XMASAMX", -1))

	if test != 2 {
		t.Errorf("result %d, expected %d", test, 2)
	}
}

func TestRealInput(t *testing.T) {
	data, _ := os.ReadFile("input_test.txt")

	result := solutionOne(string(data))
	expected := 2646

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestTestInput2(t *testing.T) {
	var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := solutionTwo(input)
	expected := 9

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestRealInput2(t *testing.T) {
	data, _ := os.ReadFile("input_test.txt")

	result := solutionTwo(string(data))
	expected := 75920122

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
