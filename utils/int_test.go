package utils

import (
	"reflect"
	"testing"
)

func TestIntMin(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{0, 0, 0},
		{4, 0, 0},
		{0, 4, 0},
		{4, 2, 2},
		{2, 4, 2},
		{-4, 2, -4},
		{2, -4, -4},
	}
	for _, testCase := range testCases {
		actual := IntMin(testCase.x, testCase.y)
		if actual != testCase.expected {
			t.Errorf(
				"%v, %v => %v, want %v",
				testCase.x, testCase.y, actual, testCase.expected)
		}
	}
}

func TestIntMax(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{0, 0, 0},
		{4, 0, 4},
		{0, 4, 4},
		{4, 2, 4},
		{2, 4, 4},
		{-4, 2, 2},
		{2, -4, 2},
	}
	for _, testCase := range testCases {
		actual := IntMax(testCase.x, testCase.y)
		if actual != testCase.expected {
			t.Errorf(
				"%v, %v => %v, want %v",
				testCase.x, testCase.y, actual, testCase.expected)
		}
	}
}

func TestIntMaxReduce(t *testing.T) {
	testCases := []struct {
		xs       []int
		init     int
		expected int
	}{
		{[]int{}, 0, 0},
		{[]int{}, 1, 1},
		{[]int{1}, 0, 1},
		{[]int{1}, 5, 5},
		{[]int{1, 2, 3, 4}, 0, 4},
		{[]int{1, 2, 3, 4}, 5, 5},
	}
	for _, testCase := range testCases {
		actual := IntMaxReduce(testCase.xs, testCase.init)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v, %v => %v, want %v",
				testCase.xs, testCase.init, actual, testCase.expected)
		}
	}
}

func TestIntSum(t *testing.T) {
	testCases := []struct {
		xs       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, 15},
	}
	for _, testCase := range testCases {
		actual := IntSum(testCase.xs)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v => %v, want %v",
				testCase.xs, actual, testCase.expected)
		}
	}
}
