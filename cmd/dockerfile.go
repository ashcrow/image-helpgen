package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ashcrow/image-helpgen/types"
	"github.com/ashcrow/image-helpgen/utils"
	"github.com/projectatomic/docker/builder/dockerfile/parser"
)

// DockerfileCommand executes the logic that is exposed via the cli at
// image-helpgen dockerfile [args]
func DockerfileCommand(dockerfilePath, template, basename string) error {
	file, err := os.Open(dockerfilePath)
	defer file.Close()

	if err != nil {
		return err
	}
	d := parser.Directive{
		EscapeSeen:           false,
		LookingForDirectives: true,
	}
	parser.SetEscapeToken(parser.DefaultEscapeToken, &d)
	node, err := parser.Parse(file, &d)
	if err != nil {
		return err
	}

	n := node
	tpl, err := types.NewTemplateRenderer(template)

	if err != nil {
		return err
	}

	// parse and set the ImageLongDescription if it exists in the Dockerfile
	if err = parseLongDescription(&tpl, dockerfilePath); err != nil {
		if err.Error() != "EOF" {
			return err
		}
	}

	// parse and set documentation which requires comments. This includes EXPOSE, VOLUME, and ENV.
	err = parseCommentDocumentation(&tpl, dockerfilePath)
	if err != nil {
		return err
	}

	// Parse what we can more efficiently using the official parser
	for {
		for _, child := range n.Children {
			switch child.Value {
			case "label":
				child = parseLabel(child, &tpl)
			case "entrypoint":
				child = parseEntrypoint(child, &tpl)
			case "cmd":
				child = parseCmd(child, &tpl)
			}
		}
		// Move to the next node if one exists. Else break out of the loop.
		if n.Next != nil {
			n = n.Next
		} else {
			break
		}
	}
	// Write out the markdown file
	tpl.WriteMarkdown(basename)
	return nil
}

// parseUsage parses the usage string pulling out specific expectations
// and adding them to the template context.
func parseUsage(tpl *types.TemplateRenderer) {
	tpl.Context.ImageExpectedDaemon = false
	if strings.Index(tpl.Context.ImageUsage, "-d ") > 0 {
		tpl.Context.ImageExpectedDaemon = true
	}

	args := strings.Split(tpl.Context.ImageUsage, " ")
	tpl.Context.ImageExpectedCaps = []string{}
	capNext := false
	for _, item := range args {
		if capNext == true {
			tpl.Context.ImageExpectedCaps = append(tpl.Context.ImageExpectedCaps, item)
			capNext = false
		} else if strings.HasPrefix(item, "--cap-add=") {
			tpl.Context.ImageExpectedCaps = append(tpl.Context.ImageExpectedCaps, item[10:])
		} else if item == "--cap-add" {
			capNext = true
		}
	}
}

func parseLabel(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	for {
		if child.Next != nil {
			switch child.Next.Value {
			case "maintainer":
				tpl.Context.ImageAuthor = utils.StripQuotes(utils.StripEmail(child.Next.Next.Value))
			case "summary":
				tpl.Context.ImageShortDescription = utils.StripQuotes(child.Next.Next.Value)
			case "name":
				tpl.Context.ImageName = utils.StripQuotes(child.Next.Next.Value)
			case "usage":
				tpl.Context.ImageUsage = utils.StripQuotes(child.Next.Next.Value)
				parseUsage(tpl)
			case "url":
				tpl.Context.ImageSeeAlso = append(tpl.Context.ImageSeeAlso, utils.StripQuotes(child.Next.Next.Value))
			}
			child = child.Next.Next
		} else {
			// break out of the inner loop once we've read in all
			break
		}
	}
	return child
}

// parseLongDescription parses the start of the Dockerfile and turning the
// starting comment(s) into the LongDescription.
func parseLongDescription(tpl *types.TemplateRenderer, dockerfilePath string) error {
	file, err := os.OpenFile(dockerfilePath, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	// Get a reader that lets us go by lines
	r := bufio.NewReader(file)
	prevLine := ""
	hashFound := false

	for {
		lineBytes, isPrefix, err := r.ReadLine()
		line := string(lineBytes)
		if err != nil {
			return err
		}

		// If this is a prefix for a line, add it to prevLine and continue
		if isPrefix {
			prevLine = prevLine + line
			continue
			// if prevLine is not empty then use it
		} else if prevLine != "" {
			tpl.Context.ImageLongDescription = tpl.Context.ImageLongDescription + prevLine
			prevLine = ""
		}

		// If we get here we need to ensure the line starts with a hash
		if strings.HasPrefix(line, "#") {
			hashFound = true
			line = strings.Trim(line, "#")
			if line == "" {
				// If it's an empty line that started with a hash then we treat
				// it like the start of a new paragraph.
				tpl.Context.ImageLongDescription = tpl.Context.ImageLongDescription + "\n\n"
			} else {
				tpl.Context.ImageLongDescription = tpl.Context.ImageLongDescription + line
			}
		} else if hashFound {
			// If we had found hashes and no longer have them we assume this is
			// the end of the LongDescription block.
			break
		}
		// If it doesn't end with a has we s
	}
	return nil
}

// parseEntrypoint parses the entrypoint defenition from the Dockerfile
func parseEntrypoint(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	if len(tpl.Context.ImageDefaultCommand) > 0 {
		tpl.Context.ImageDefaultCommand = child.Next.Value + " " + tpl.Context.ImageDefaultCommand
	} else {
		tpl.Context.ImageDefaultCommand = child.Next.Value
	}
	return child.Next.Next
}

// parseCmd parses the cmd defenition from the Dockerfile
func parseCmd(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	if len(tpl.Context.ImageDefaultCommand) > 0 {
		tpl.Context.ImageDefaultCommand = tpl.Context.ImageDefaultCommand + " " + child.Next.Value
	} else {
		tpl.Context.ImageDefaultCommand = child.Next.Value
	}

	// If we have another
	if child.Next.Next != nil {
		child = parseCmd(child.Next, tpl)
	}
	return child.Next

}

// parseCommentDocumentation takes care of parsing Volume, Port, and Env sections and adding
// found documentation to the structures.
func parseCommentDocumentation(tpl *types.TemplateRenderer, dockerfilePath string) error {
	file, err := os.OpenFile(dockerfilePath, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return nil
	}

	// Get a reader that lets us go by lines
	r := bufio.NewReader(file)
	prevLine := ""

	for {
		lineBytes, _, err := r.ReadLine()
		line := string(lineBytes)
		if err != nil {
			// If the error is the end of the file, we're done. No true error.
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}

		if strings.HasPrefix(line, "EXPOSE") {
			parseExpose(line, prevLine, tpl)
		} else if strings.HasPrefix(line, "VOLUME") {
			parseVolume(line, prevLine, tpl)
		} else if strings.HasPrefix(line, "ENV") {
			parseEnv(line, prevLine, tpl)
		}
		prevLine = line
	}
}

// parseExpose parses an EXPOSE line and attempts to pull in documentation
func parseExpose(line, prevLine string, tpl *types.TemplateRenderer) {
	expectedContainerPort, err := strconv.Atoi(line[7:])
	if err != nil {
		return
	}
	newPort := types.Port{
		Container:   expectedContainerPort,
		Host:        0,
		Description: "No documentation provided",
	}
	if strings.HasPrefix(prevLine, "#") {
		doc := strings.SplitN(strings.Trim(prevLine, "#"), " ", 2)
		mapping := strings.SplitN(doc[0], "->", 2)
		containerPort, _ := strconv.Atoi(mapping[0])
		hostPort, _ := strconv.Atoi(mapping[1])
		// If the expected port and found doc doesn't match use the expected port
		// and note in the documentation there is a problem
		if containerPort != expectedContainerPort {
			newPort.Description = fmt.Sprintf(
				"Unknown. Documentation error found. %d != %d", expectedContainerPort, containerPort)
		} else {
			// otherwise use the doc as the port
			newPort.Container = containerPort
			newPort.Host = hostPort
			newPort.Description = doc[1]
		}
	}
	tpl.Context.ImagePorts = append(tpl.Context.ImagePorts, newPort)
}

// parseVolume parses a VOLUME line and attempts to pull in documentation
func parseVolume(line, prevLine string, tpl *types.TemplateRenderer) {
	expectedVolume := line[7:]
	newVolume := types.Volume{
		Container:   expectedVolume,
		Host:        "UNKNOWN",
		Description: "No documentation provided",
	}

	if strings.HasPrefix(prevLine, "#") {
		doc := strings.SplitN(strings.Trim(prevLine, "#"), " ", 2)
		mapping := strings.SplitN(doc[0], "->", 2)
		containerVolume := mapping[0]
		hostVolume := mapping[1]

		if expectedVolume != containerVolume {
			newVolume.Description = fmt.Sprintf(
				"UNKNOWN. Documentation error. %s != %s", expectedVolume, containerVolume)
		} else {
			newVolume.Container = containerVolume
			newVolume.Host = hostVolume
			newVolume.Description = doc[1]
		}
	}
	tpl.Context.ImageVolumes = append(tpl.Context.ImageVolumes, newVolume)
}

// parseEnv parses an ENV line and attempts to pull in documentation
func parseEnv(line, prevLine string, tpl *types.TemplateRenderer) {
	mapping := strings.SplitN(line[3:], "=", 2)
	envVar := types.EnvironmentVariable{
		Name:        mapping[0],
		Default:     mapping[1],
		Description: "No documentation provided",
	}
	if strings.HasPrefix(prevLine, "#") {
		envVar.Description = strings.Trim(prevLine, "#")
	}
	tpl.Context.ImageEnvironmentVariables = append(tpl.Context.ImageEnvironmentVariables, envVar)
}
