package main

import (
	"reflect"
	"testing"
)

// Test for Q1: Sorting Words Based on 'a' Count
func TestSortWordsByA(t *testing.T) {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	expected := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}

	result := sortWordsByA(input) // Replace with your implementation

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but got %v", expected, result)
	}

	t.Log("TestSortWordsByA passed successfully.")
}

func sortWordsByA(input []string) any {
	panic("unimplemented")
}

// Test for Q2: Recursive Function
func TestRecursiveFunction(t *testing.T) {
	input := 9
	expected := []int{2, 4, 9}

	var result []int
	testHelperRecursiveFunction(input, &result) // Replace with your recursive function implementation

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but got %v", expected, result)
	}

	t.Log("TestRecursiveFunction passed successfully.")
}

func testHelperRecursiveFunction(n int, result *[]int) {
	// Example logic for testing recursive function
	if n == 2 {
		*result = append(*result, 2)
		return
	} else if n == 4 {
		*result = append(*result, 4)
		return
	}
	*result = append(*result, 9)
}

// Test for Q3: Finding the Most Repeated Data
func TestMostRepeatedData(t *testing.T) {
	input := []string{"apple", "pie", "apple", "red", "red", "red"}
	expected := "red"

	result := findMostRepeated(input)

	if result != expected {
		t.Fatalf("Expected %s, but got %s", expected, result)
	}

	t.Log("TestMostRepeatedData passed successfully.")
}

func findMostRepeated(input []string) string {
	// Example implementation for finding the most repeated data
	counts := make(map[string]int)
	for _, word := range input {
		counts[word]++
	}

	maxCount := 0
	mostRepeated := ""
	for word, count := range counts {
		if count > maxCount {
			maxCount = count
			mostRepeated = word
		}
	}
	return mostRepeated
}
