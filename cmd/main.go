/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/ashcrow/go-md2man/md2man"
)

const defaultTemplate = "/etc/image-helpgen/template.tpl"

type port struct {
	Container   int
	Host        int
	Description string
}

type volume struct {
	Container   string
	Host        string
	Description string
}

type environmentVariable struct {
	Name        string
	Description string
}

type tplContext struct {
	ImageName                 string
	ImageAuthor               string
	ImageDocDate              string
	ImageShortDescription     string
	ImageLongDescription      string
	ImageUsage                string
	ImageEnvironmentVariables []environmentVariable
	ImageVolumes              []volume
	ImagePorts                []port
	ImageSeeAlso              string
}

// Shortcut to painic when an error is returned
func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

// TemplateRenderer provides a structure for working with a template and then
// rendering the results.
type TemplateRenderer struct {
	reader   *bufio.Reader
	context  tplContext
	template *template.Template
}

// NewTemplateRenderer creates a new TemplateRenderer instance and returns it.
func NewTemplateRenderer(tf string) TemplateRenderer {
	tr := TemplateRenderer{}
	var err error
	tr.template, err = template.ParseFiles(tf)
	panicOnErr(err)

	tr.reader = bufio.NewReader(os.Stdin)
	tr.context = tplContext{}
	return tr
}

func (t *TemplateRenderer) readString(prompt string) string {
	fmt.Printf(prompt + ": ")
	result, _ := t.reader.ReadString('\n')
	return strings.TrimSuffix(result, "\n")
}

func (t *TemplateRenderer) readText(prompt string) string {
	fmt.Printf(prompt + " (Enter . alone on a line to end):\n")
	result := ""
	for {
		data, _ := t.reader.ReadString('\n')
		if data == ".\n" {
			break
		}
		result = result + data
	}
	return strings.TrimSuffix(result, ".\n")
}

func (t *TemplateRenderer) readEnvironmentVariables() {
	fmt.Println("Enter Environment Variable information. Enter empty name to finish.")
	for {
		name := t.readString("Name")
		if name == "" {
			break
		}
		description := t.readString("Description")
		t.context.ImageEnvironmentVariables = append(
			t.context.ImageEnvironmentVariables,
			environmentVariable{Name: name, Description: description})
	}
}

func (t *TemplateRenderer) readPorts() {
	fmt.Println("Enter port information. Enter empty host port to finish.")
	for {
		hp := t.readString("Host Port")
		if hp == "" {
			break
		}
		cp := t.readString("Container Port")
		description := t.readString("Description")
		containerPort, _ := strconv.Atoi(cp)
		hostPort, _ := strconv.Atoi(hp)
		t.context.ImagePorts = append(
			t.context.ImagePorts,
			port{Container: containerPort, Host: hostPort, Description: description})
	}
}

func (t *TemplateRenderer) readVolumes() {
	fmt.Println("Enter volume information. Enter empty host volume to finish.")
	for {
		hv := t.readString("Host Volume")
		if hv == "" {
			break
		}
		cv := t.readString("Container Volume")
		description := t.readString("Description")
		t.context.ImageVolumes = append(
			t.context.ImageVolumes,
			volume{Container: cv, Host: hv, Description: description})
	}
}

// Write writes rendered templates to md and man formats.
func (t *TemplateRenderer) Write(basename string) {
	// buffer to hold the rendered result
	data := []byte{}
	out := bytes.NewBuffer(data)

	// Render the template
	t.template.Execute(out, t.context)
	// Write out the markdown
	err := ioutil.WriteFile(basename+".md", out.Bytes(), 0644)
	panicOnErr(err)
	// Write out the man file
	man := md2man.Render(out.Bytes())
	err = ioutil.WriteFile(basename+".1", man, 0644)
	panicOnErr(err)
}

// main function for CLI
func main() {
	fmt.Println("container-help:  Copyright helpgen(C) 2017 Steve Milner")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY.")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions. See COPYING for details.")

	var basename = flag.String(
		"basename", "help", "Base name to use for file writing")
	var template = flag.String(
		"template", defaultTemplate, "Template to use when rendering")
	flag.Parse()

	tr := NewTemplateRenderer(*template)
	tr.context.ImageName = tr.readString("Image Name")
	tr.context.ImageAuthor = tr.readString("Image Author")
	tr.context.ImageShortDescription = tr.readString("Short Description")
	tr.context.ImageLongDescription = tr.readText("Long Description")
	tr.context.ImageUsage = tr.readText("Image Usage")
	tr.context.ImageSeeAlso = tr.readString("See Also")
	tr.readEnvironmentVariables()
	tr.readPorts()
	tr.readVolumes()

	// Set the doc date to now
	now := time.Now()
	tr.context.ImageDocDate = fmt.Sprintf("%s %d", now.Month(), now.Year())
	// Render the template
	tr.Write(*basename)
}
