package types

import (
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/ashcrow/image-helpgen/utils"
)

// Port is a representation of a single Port entity
type Port struct {
	Container   int
	Host        int
	Description string
}

// Volume is a representation of a single Volume entity
type Volume struct {
	Container   string
	Host        string
	Description string
}

// EnvironmentVariable is a representation of a single environment variable entity
type EnvironmentVariable struct {
	Name        string
	Default     string
	Description string
}

// TplContext represents the context used by TemplateRenderer to render content
type TplContext struct {
	ImageName                 string
	ImageAuthor               string
	ImageDocDate              string
	ImageShortDescription     string
	ImageLongDescription      string
	ImageUsage                string
	ImageEnvironmentVariables []EnvironmentVariable
	ImageVolumes              []Volume
	ImagePorts                []Port
	ImageExpectedCaps         []string
	ImageExpectedDaemon       bool
	ImageSeeAlso              []string
}

// TemplateRenderer provides a structure for working with a template and then
// rendering the results.
type TemplateRenderer struct {
	Reader   *bufio.Reader
	Context  TplContext
	Template *template.Template
}

// NewTemplateRenderer creates a new TemplateRenderer instance and returns it.
func NewTemplateRenderer(tf string) (TemplateRenderer, error) {
	tr := TemplateRenderer{}
	var err error
	tr.Template, err = template.ParseFiles(tf)
	if err != nil {
		return TemplateRenderer{}, err
	}

	tr.Reader = bufio.NewReader(os.Stdin)
	tr.Context = TplContext{
		ImageDocDate: utils.GenerateDocDate(),
	}
	return tr, nil
}

// WriteMarkdown writes a markdown version of the output.
func (t *TemplateRenderer) WriteMarkdown(basename string) error {
	data := []byte{}
	out := bytes.NewBuffer(data)
	fileName := basename + ".md"
	// Render the template
	t.Template.Execute(out, t.Context)
	// Write out the markdown
	err := ioutil.WriteFile(fileName, out.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

// WriteMan writes rendered man file based off the markdown file.
func (t *TemplateRenderer) WriteMan(basename string) {
	utils.WriteManFromMd(basename)
}

// Write writes rendered templates to md and man formats.
func (t *TemplateRenderer) Write(basename string) {
	t.WriteMarkdown(basename)
	t.WriteMan(basename)
}
