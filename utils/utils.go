package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/cpuguy83/go-md2man/md2man"
)

// ExitOnErr prints the error and exits
func ExitOnErr(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s.\nExiting...\n", e)
		os.Exit(1)
	}
}

// WriteManFromMd writes rendered man file based off the markdown file.
func WriteManFromMd(basename string) {
	// buffer to hold the rendered result
	mdData, err := ioutil.ReadFile(basename + ".md")
	ExitOnErr(err)
	// Write out the man file
	man := md2man.Render(mdData)
	err = ioutil.WriteFile(basename+".1", man, 0644)
	ExitOnErr(err)
}

// StripEmail removes an email address from a string.
func StripEmail(line string) string {
	idx := strings.Index(line, "@")
	if idx != -1 {
		splitLine := strings.Split(line, "@")
		line = splitLine[0][0:strings.LastIndex(splitLine[0], " ")]
	}
	return line
}

// StripQuotes removes surrounding  quotes from a string.
func StripQuotes(line string) string {
	return strings.TrimRight(strings.TrimLeft(line, "\""), "\"")
}

// GenerateDocDate creates a document date for use in man/md files
func GenerateDocDate() string {
	now := time.Now()
	return fmt.Sprintf("%s %d", now.Month(), now.Year())
}
