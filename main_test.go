package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	input1    = "Test branch name"
	expected1 = "test-branch-name"
)

var testMap = map[string]string{
	"Test branch name":             "test-branch-name",
	"Test112()*&^%#@! branch name": "test112-branch-name",
	"Test-dashed-branch-name":      "test-dashed-branch-name",
	"Test with -> and spaces":      "test-with-and-spaces",
	`
		Hello
		branch
		world
		!
	`: "hello-branch-world",
}

func Setup() {
	for _, expected := range testMap {
		exec.Command("git", "branch", "-D", expected).Run()
	}
}

func Teardown() {
	for _, expected := range testMap {
		exec.Command("git", "branch", "-D", expected).Run()
	}
}

func TestMain(m *testing.M) {
	// Set up
	// Remove test branches if they exist

	current, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	exec.Command("git", "checkout", "main").Run()

	Setup()

	// Run the tests
	code := m.Run()

	exec.Command("git", "checkout", strings.TrimSpace(string(current))).Run()

	Teardown()

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
	for input, expected := range testMap {
		result := SanitizeBranchName(input)
		if result != expected {
			t.Errorf("SanitizeBranchName(%s) != %s, Actually got: %s", input, expected, result)
		}
	}
}
