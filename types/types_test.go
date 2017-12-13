package types

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewTemplateRenderer(t *testing.T) {
	_, err := NewTemplateRenderer("../template.tpl")
	if err != nil {
		t.Error("Error returned when not expected")
	}
	_, err = NewTemplateRenderer("")
	if err == nil {
		t.Error("Error not returned but expected")
	}
}

func writeMarkdown() (string, string) {
	tr, _ := NewTemplateRenderer("../template.tpl")
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
	tr, _ := NewTemplateRenderer("../template.tpl")
	tr.WriteMarkdown(tmpPath)
}

func TestWrite(t *testing.T) {
	tr, _ := NewTemplateRenderer("../template.tpl")
	tmpFile, _ := ioutil.TempFile("", "")
	expectedMdPath := tmpFile.Name() + ".md"
	expectedManPath := tmpFile.Name() + ".1"
	defer os.Remove(tmpFile.Name())
	defer os.Remove(expectedMdPath)
	defer os.Remove(expectedManPath)
	tr.Write(tmpFile.Name())
}
