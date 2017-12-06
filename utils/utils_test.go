package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPanicOnErr(t *testing.T) {
	// Should not panic
	PanicOnErr(nil)
	// recover for upcoming panic test
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic did not occur when it should have")
		}
	}()
	// PANIC!!!
	PanicOnErr(errors.New("error"))
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
			t.Error("%s != %s", tc.Input, tc.Expected)
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
			t.Error("%s != %s", tc.Input, tc.Expected)
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
		t.Errorf("%i != %i", now.Year(), year)
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

// func WriteManFromMd(basename string) {
// 	// buffer to hold the rendered result
// 	mdData, err := ioutil.ReadFile(basename + ".md")
// 	PanicOnErr(err)
// 	// Write out the man file
// 	man := md2man.Render(mdData)
// 	err = ioutil.WriteFile(basename+".1", man, 0644)
// 	PanicOnErr(err)
// }
//
