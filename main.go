package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// Read input string from command line arguments
	input := os.Args[len(os.Args)-1]

	fmt.Println("Input string:", input)

	// Check if input string is empty
	if input == "" {
		fmt.Println("Error: No input string provided")
		fmt.Println("---------------------------------")
		fmt.Println("Usage: go-create-branch \"branchname\"")
		fmt.Println("---------------------------------")
		return
	}

	input = SanitizeBranchName(input)
	fmt.Println(input)
	// Check if Git branch already exists
	out, err := exec.Command("git", "branch", "--list", input).Output()

	if string(out) != "" {
		fmt.Println("Error: Git branch already exists")
		return
	}

	if err != nil {
		fmt.Println("Error checking Git branch:", err)
		os.Exit(1)
	}

	// Create a Git branch with the resulting string
	cmd := exec.Command("git", "checkout", "-b", input)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating Git branch:", err)
		return
	}

	// Print the resulting string
	fmt.Printf("Created Git branch: %s\n", input)
}

func SanitizeBranchName(input string) string {
	// Remove all special characters using regular expression
	reg := regexp.MustCompile("[^a-zA-Z0-9 -]+")
	input = reg.ReplaceAllString(input, "")

	// Replace spaces with dashes
	input = strings.ReplaceAll(input, " ", "-")

	// Convert to lowercase
	input = strings.ToLower(input)

	return input
}
