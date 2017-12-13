package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestExitOnErr verifies that if an error is passed in the execution exits.
// Based off https://stackoverflow.com/questions/26225513/how-to-test-os-exit-scenarios-in-go
func TestExitOnErr(t *testing.T) {
	if os.Getenv("EXIT_TEST") == "1" {
		ExitOnErr(errors.New("error"))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExitOnErr")
	cmd.Env = append(os.Environ(), "EXIT_TEST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// Exit case
	t.Fatalf("process ran with err %v, want exit status 1", err)
	// Non exit case
	ExitOnErr(nil)
}

func TestStripEmail(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{"Some Person <blah@example.org>", "Some Person"},
		{"Some Person blah@example.org", "Some Person"},
		{"Some Person \"blah@example.org\"", "Some Person"},
		{"Some Person", "Some Person"},
	}

	for _, tc := range cases {
		if StripEmail(tc.Input) != tc.Expected {
			t.Errorf("%s != %s", tc.Input, tc.Expected)
		}
	}
}

func TestStripQuotes(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{"test", "test"},
		{"\"test\"", "test"},
		{"\"test", "test"},
		{"test\"", "test"},
	}

	for _, tc := range cases {
		if StripQuotes(tc.Input) != tc.Expected {
			t.Errorf("%s != %s", tc.Input, tc.Expected)
		}
	}
}

func TestGenerateDocDate(t *testing.T) {
	result := GenerateDocDate()
	now := time.Now()
	dateSplit := strings.Split(result, " ")
	if fmt.Sprintf("%s", now.Month()) != dateSplit[0] {
		t.Errorf("Month doesn not match: |%#v| != |%#v|", now.Month(), dateSplit[0])
	}
	year, _ := strconv.Atoi(dateSplit[1])
	if now.Year() != year {
		t.Errorf("%d != %d", now.Year(), year)
	}
}

func TestWriteManFromMd(t *testing.T) {
	// Create a temporary md file with the content from the example
	tmpFile, _ := ioutil.TempFile("", "md")
	expectedManPath := tmpFile.Name() + ".1"
	// Make sure to clean up
	defer os.Remove(tmpFile.Name())
	defer os.Remove(tmpFile.Name() + ".md")
	defer os.Remove(expectedManPath)

	// Copy content from example to temp file
	mdContent, _ := ioutil.ReadFile("example/help.md")
	tmpFile.Write(mdContent)
	tmpFile.Close()
	// Link it to the same name with a .md suffix
	os.Link(tmpFile.Name(), tmpFile.Name()+".md")
	// Run the WriteManFromMd command with the temp file as the basename
	WriteManFromMd(tmpFile.Name())
	_, err := os.Stat(expectedManPath)
	if err != nil {
		t.Errorf("The man page was not generated")
	}
}
