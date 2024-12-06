package day3

import (
	"os"
	"testing"
)

func TestTestInput(t *testing.T) {
	var input = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	result := solutionOne(input)
	expected := 161

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestRealInput(t *testing.T) {
	data, _ := os.ReadFile("input_test.txt")

	result := solutionOne(string(data))
	expected := 156388521

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestTestInput2(t *testing.T) {
	var input = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	result := solutionTwo(input)
	expected := 48

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
