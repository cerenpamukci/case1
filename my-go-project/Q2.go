package main

import "fmt"

//	Define a recursive function that takes one integer as a parameter
//
// This function prints numbers in descending order until 1 and then back up
func printPattern(n int) {

	//  Define the base case
	// If n is 1, print it and return (this stops the recursion)
	if n == 1 {
		fmt.Print(n, " ") // Print the number 1
		return
	}

	// Recursive step for descending order
	// Print the current number before recursive call
	fmt.Print(n, " ")

	// Call the function with n-1 to go deeper into the recursion
	printPattern(n - 1)

	// Step 4: Recursive step for ascending order
	// Print the current number after the recursion
	fmt.Print(n, " ")
}

func Q2() {
	// Call the recursive function with an example input
	// Input value
	input := 5

	// Print the pattern
	fmt.Println("Output for input", input, ":")
	printPattern(input)
	fmt.Println() // New line for clean output
}
