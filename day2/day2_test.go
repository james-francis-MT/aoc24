package day2

import (
	"os"
	"testing"
)

func TestTestInput(t *testing.T) {
	var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := solutionOne(input)
	expected := 2

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestRealInput(t *testing.T) {
	data, _ := os.ReadFile("input_test.txt")

	result := solutionOne(string(data))
	expected := 379

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestTestInput2(t *testing.T) {
	var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := solutionTwo(input)
	expected := 4

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestRealInput2(t *testing.T) {
	data, _ := os.ReadFile("input_test.txt")

	result := solutionTwo(string(data))
	expected := 430

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
