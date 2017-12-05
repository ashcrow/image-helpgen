package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cpuguy83/go-md2man/md2man"
)

// PanicOnErr is a shortcut to painic when an error is returned
func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

// WriteManFromMd writes rendered man file based off the markdown file.
func WriteManFromMd(basename string) {
	// buffer to hold the rendered result
	mdData, err := ioutil.ReadFile(basename + ".md")
	PanicOnErr(err)
	// Write out the man file
	man := md2man.Render(mdData)
	err = ioutil.WriteFile(basename+".1", man, 0644)
	PanicOnErr(err)
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
