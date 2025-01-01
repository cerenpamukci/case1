package main

import "fmt"

// Function to find the most repeated string in an array
func mostRepeatedElement(arr []string) string {
	// Step 1: Create a map to count the occurrences of each string
	// The keys are the strings in the array, and the values are their respective counts.
	counts := make(map[string]int)

	// Step 2: Iterate through the array and populate the map
	// For each string in the array, increment its count in the map.
	for _, str := range arr {
		counts[str]++ // Increment the count for the current string
	}

	// Step 3: Initialize variables to track the most repeated string
	// `maxCount` stores the highest frequency found so far.
	// `mostRepeated` stores the string associated with the highest frequency.
	maxCount := 0
	mostRepeated := ""

	// Step 4: Iterate through the map to find the string with the highest frequency
	// Use a loop to compare the count of each string with the current `maxCount`.
	for key, value := range counts {
		if value > maxCount {
			maxCount = value   // Update the highest frequency
			mostRepeated = key // Update the most repeated string
		}
	}

	// Step 5: Return the most repeated string
	return mostRepeated
}

func Q3() {
	// Step 6: Test with different datasets
	// Example dataset
	arr := []string{"apple", "pie", "apple", "red", "red", "red"}

	// Step 7: Call the function and store the result
	result := mostRepeatedElement(arr)

	// Step 8: Print the result
	fmt.Println("The most repeated element is:", result)
}
