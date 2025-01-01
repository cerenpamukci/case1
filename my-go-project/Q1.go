package main

import (
	"fmt"     // Package used to print data to the terminal
	"sort"    // Package used for sorting slices
	"strings" // Package used for string manipulation
)

// Step: Define a function to count the occurrences of the letter 'a' in a word

func countA(word string) int {
	// Step: Use the strings.Count function to count 'a' characters in the input string

	return strings.Count(word, "a")
}

// Define a function to sort the words
func sortWordsByA(words []string) []string {

	// Use sort.Slice to sort the slice of words based on custom criteria
	sort.Slice(words, func(i, j int) bool {

		// Count the number of 'a' characters in the two words being compared
		countA1 := countA(words[i]) // İlk kelimedeki "a" sayısı
		countA2 := countA(words[j]) // İkinci kelimedeki "a" sayısı

		// If the counts of 'a' are different, sort by the number of 'a's in descending order
		if countA1 != countA2 {
			return countA1 > countA2
		}

		// If the counts of 'a' are equal, sort by word length in descending order
		return len(words[i]) > len(words[j])
	})

	// Return the sorted list
	return words
}

func Q1() {
	// Define the input data
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// Call the sorting function
	sortedWords := sortWordsByA(words)

	// Print the sorted list
	fmt.Println(sortedWords)
}
