package types

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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
	ImageSeeAlso              string
}

// TemplateRenderer provides a structure for working with a template and then
// rendering the results.
type TemplateRenderer struct {
	Reader   *bufio.Reader
	Context  TplContext
	Template *template.Template
}

// NewTemplateRenderer creates a new TemplateRenderer instance and returns it.
func NewTemplateRenderer(tf string) TemplateRenderer {
	tr := TemplateRenderer{}
	var err error
	tr.Template, err = template.ParseFiles(tf)
	utils.PanicOnErr(err)

	tr.Reader = bufio.NewReader(os.Stdin)
	tr.Context = TplContext{
		ImageDocDate: utils.GenerateDocDate(),
	}
	return tr
}

// ReadString reads a single string and returns the result
func (t *TemplateRenderer) ReadString(prompt string) string {
	fmt.Printf(prompt + ": ")
	result, _ := t.Reader.ReadString('\n')
	return strings.TrimSuffix(result, "\n")
}

// ReadText reads a block of text and returns the result
func (t *TemplateRenderer) ReadText(prompt string) string {
	fmt.Printf(prompt + " (Enter . alone on a line to end):\n")
	result := ""
	for {
		data, _ := t.Reader.ReadString('\n')
		if data == ".\n" {
			break
		}
		result = result + data
	}
	return strings.TrimSuffix(result, ".\n")
}

// ReadEnvironmentVariables populates and returns a list of EnvironmentVariables
func (t *TemplateRenderer) ReadEnvironmentVariables() {
	fmt.Println("Enter Environment Variable information. Enter empty name to finish.")
	for {
		name := t.ReadString("Name")
		if name == "" {
			break
		}
		defaultValue := t.ReadString("Default Value")
		description := t.ReadString("Description")
		t.Context.ImageEnvironmentVariables = append(
			t.Context.ImageEnvironmentVariables,
			EnvironmentVariable{
				Name:        name,
				Default:     defaultValue,
				Description: description})
	}
}

// ReadPorts reads and populates a list of Ports
func (t *TemplateRenderer) ReadPorts() {
	fmt.Println("Enter port information. Enter empty host port to finish.")
	for {
		hp := t.ReadString("Host Port")
		if hp == "" {
			break
		}
		cp := t.ReadString("Container Port")
		description := t.ReadString("Description")
		containerPort, _ := strconv.Atoi(cp)
		hostPort, _ := strconv.Atoi(hp)
		t.Context.ImagePorts = append(
			t.Context.ImagePorts,
			Port{
				Container:   containerPort,
				Host:        hostPort,
				Description: description})
	}
}

// ReadVolumes reads and populates a list of Volumes
func (t *TemplateRenderer) ReadVolumes() {
	fmt.Println("Enter volume information. Enter empty host volume to finish.")
	for {
		hv := t.ReadString("Host Volume")
		if hv == "" {
			break
		}
		cv := t.ReadString("Container Volume")
		description := t.ReadString("Description")
		t.Context.ImageVolumes = append(
			t.Context.ImageVolumes,
			Volume{Container: cv, Host: hv, Description: description})
	}
}

// WriteMarkdown writes a markdown version of the output.
func (t *TemplateRenderer) WriteMarkdown(basename string) {
	data := []byte{}
	out := bytes.NewBuffer(data)
	fileName := basename + ".md"
	// Render the template
	t.Template.Execute(out, t.Context)
	// Write out the markdown
	err := ioutil.WriteFile(fileName, out.Bytes(), 0644)
	utils.PanicOnErr(err)
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
