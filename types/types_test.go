package types

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewTemplateRenderer(t *testing.T) {
	NewTemplateRenderer("../template.tpl")
	// recover for upcoming expected panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic did not occur when it should have")
		}
	}()
	NewTemplateRenderer("")
}

func writeMarkdown() (string, string) {
	tr := NewTemplateRenderer("../template.tpl")
	tmpFile, _ := ioutil.TempFile("", "")
	expectedMdPath := tmpFile.Name() + ".md"
	tr.WriteMarkdown(tmpFile.Name())
	return tmpFile.Name(), expectedMdPath
}

func TestWriteMarkdown(t *testing.T) {
	tmpPath, mdPath := writeMarkdown()
	defer os.Remove(tmpPath)
	defer os.Remove(mdPath)
}

func TestWriteMan(t *testing.T) {
	tmpPath, mdPath := writeMarkdown()
	defer os.Remove(tmpPath)
	defer os.Remove(mdPath)
	defer os.Remove(tmpPath + ".1")
	tr := NewTemplateRenderer("../template.tpl")
	tr.WriteMarkdown(tmpPath)
}

func TestWrite(t *testing.T) {
	tr := NewTemplateRenderer("../template.tpl")
	tmpFile, _ := ioutil.TempFile("", "")
	expectedMdPath := tmpFile.Name() + ".md"
	expectedManPath := tmpFile.Name() + ".1"
	defer os.Remove(tmpFile.Name())
	defer os.Remove(expectedMdPath)
	defer os.Remove(expectedManPath)
	tr.Write(tmpFile.Name())
}
