package main // Tells Go this file should compile as an executable program, not a library.

import (
	"bufio"   // buffered I/O operations.
	"fmt"     // formatted I/O with functions similar to C's printf and scanf.
	"os"      // interact with operating system functionality, including standard input.
	"strconv" // converting strings to other types (like integers).
	"strings" // manipulating strings.
)

// Function to check voting eligibility
func isEligibleToVote(age int) bool { // Function name, parameters, and return type
	if age >= 18 {
		return true
	}
	return false
}

func main() { // The main function is the entry point of the program
	reader := bufio.NewReader(os.Stdin) // Creates a new buffered reader for reading from the standard input (os.Stdin).

	fmt.Println("Enter your age: ") // Prompt the user to enter their age

	// Read the input from the user
	input, _ := reader.ReadString('\n') // Read the input until a newline character is encountered
	input = strings.TrimSpace(input)    // Remove leading and trailing whitespace

	// Convert the input string to an integer
	age, err := strconv.Atoi(input) // Convert the input string to an integer, and return error (if any). The error is stored in err.
	if err != nil {
		fmt.Println("Please enter a valid number")
		return // Exit the program if the input is not a number
	}

	// Call the isEligibleToVote function to check eligibility
	if isEligibleToVote(age) {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
