package main

import (
	"fmt"
	"regexp"
	"strings"
	"os"
	"os/exec"
)

func main() {
	// Read input string from command line arguments
	input := strings.Join(os.Args[1:], " ")

	// Remove all special characters using regular expression
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	input = reg.ReplaceAllString(input, "")

	// Replace spaces with dashes
	input = strings.ReplaceAll(input, " ", "-")

	// Convert to lowercase
	input = strings.ToLower(input)

	// Create a Git branch with the resulting string
	cmd := exec.Command("git", "checkout", "-b", input)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error creating Git branch:", err)
		return
	}

	// Print the resulting string
	fmt.Printf("Created Git branch: %s\n", input)
}

