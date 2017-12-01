package utils

import (
	"io/ioutil"

	"github.com/ashcrow/go-md2man/md2man"
)

// Shortcut to painic when an error is returned
func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

// WriteMan writes rendered man file based off the markdown file.
func WriteManFromMd(basename string) {
	// buffer to hold the rendered result
	mdData, err := ioutil.ReadFile(basename + ".md")
	PanicOnErr(err)
	// Write out the man file
	man := md2man.Render(mdData)
	err = ioutil.WriteFile(basename+".1", man, 0644)
	PanicOnErr(err)
}
