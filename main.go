package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	// Version of the application
	Version = "1.0.0"
)

func main() {
	// Read input string from command line arguments
	var inputFlag string
	var input string

	var versionFlag bool
	var helpFlag bool

	flag.BoolVar(&helpFlag, "help", false, "Print help information")
	flag.BoolVar(&versionFlag, "version", false, "Print version information")
	flag.StringVar(&inputFlag, "name", "", "Branch name to create")

	flag.Parse()

	if versionFlag {
		fmt.Println("go-create-branch version:", Version)
		return
	}

	if helpFlag {
		fmt.Println("go-create-branch version:", Version)
		fmt.Println("---------------------------------")
		fmt.Println("Usage: go-create-branch \"branchname\"")
		fmt.Println("---------------------------------")
		flag.PrintDefaults()
		return
	}

	if inputFlag == "" {
		input = strings.Join(flag.Args(), "-")
	} else {
		input = inputFlag
	}

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
	fmt.Println("Sanitized string:", input)
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
	//Replace newlines with spaces
	input = strings.ReplaceAll(input, "\n", " ")

	//Trim leading and trailing spaces
	input = strings.TrimSpace(input)

	// Remove all special characters using regular expression
	reg := regexp.MustCompile("[^a-zA-Z0-9 -]+")
	input = reg.ReplaceAllString(input, "")

	// Replace spaces with dashes
	input = strings.ReplaceAll(input, " ", "-")

	reg = regexp.MustCompile("-{2,}")
	input = reg.ReplaceAllString(input, "-")

	//Trim leading and trailing dashes
	input = strings.Trim(input, "-")

	// Convert to lowercase
	input = strings.ToLower(input)

	return input
}
