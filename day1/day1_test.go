package main

import (
	"os"
	"testing"
)

func TestTestInput(t *testing.T) {
	var input = `3   4
4   3
2   5
1   3
3   9
3   3`
    result := solutionOne(input)
    expected := 11

    if result != expected {
        t.Errorf("result %d, expected %d", result, expected)
    }
}

func TestRealInput(t *testing.T) {
    data, _ := os.ReadFile("input_test.txt")

    result := solutionOne(string(data))
    expected := 3569916

    if result != expected {
        t.Errorf("result %d, expected %d", result, expected)
    }
}

func TestTestInput2(t *testing.T) {
	var input = `3   4
4   3
2   5
1   3
3   9
3   3`
    result := solutionTwo(input)
    expected := 31

    if result != expected {
        t.Errorf("result %d, expected %d", result, expected)
    }
}

func TestRealInput2(t *testing.T) {
    data, _ := os.ReadFile("input_test.txt")

    result := solutionTwo(string(data))
    expected := 26407426

    if result != expected {
        t.Errorf("result %d, expected %d", result, expected)
    }
}
