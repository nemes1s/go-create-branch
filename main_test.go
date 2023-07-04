package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	input1    = "Test branch name"
	input2    = "Test112()*&^%#@! branch name 2"
	input3    = "Test-dashed-branch-name"
	input4    = "Test with -> and spaces"
	expected1 = "test-branch-name"
	expected2 = "test112-branch-name-2"
	expected3 = "test-dashed-branch-name"
	expected4 = "test-with-and-spaces"
)

func Setup() {

}

func TestMain(m *testing.M) {
	// Set up
	// Remove test branches if they exist

	current, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	exec.Command("git", "checkout", "main").Run()
	exec.Command("git", "branch", "-D", expected1).Run()
	exec.Command("git", "branch", "-D", expected2).Run()

	// Run the tests
	code := m.Run()

	exec.Command("git", "checkout", strings.TrimSpace(string(current))).Run()

	os.Exit(code)
}

func TestMainFunc(t *testing.T) {
	os.Args = append(os.Args, input1)

	main()

	out, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	if strings.TrimSpace(string(out)) != string(expected1) {
		t.Errorf("main() failed, expected %s, got %s", expected1, out)
	}
}

func TestSanitizeBranchName(t *testing.T) {
	result1 := SanitizeBranchName(input1)
	result2 := SanitizeBranchName(input2)
	result3 := SanitizeBranchName(input3)
	result4 := SanitizeBranchName(input4)
	if result1 != expected1 {
		t.Errorf("SanitizeBranchName(%s) != %s, Actually got: %s", input1, expected1, result1)
	}

	if result2 != expected2 {
		t.Errorf("SanitizeBranchName(\"%s\") != %s, Actually got: %s", input2, expected2, result2)
	}

	if result3 != expected3 {
		t.Errorf("SanitizeBranchName(\"%s\") != %s, Actually got: %s", input3, expected3, result3)
	}

	if result4 != expected4 {
		t.Errorf("SanitizeBranchName(\"%s\") != %s, Actually got: %s", input4, expected4, result4)
	}
}
