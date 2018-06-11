package cmd

import (
	"bufio"
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

	for {
		for _, child := range n.Children {
			switch child.Value {
			case "env":
				child = parseEnvironmentVariables(child, &tpl)
			case "expose":
				child = parseExpose(child, &tpl)
			case "volume":
				child = parseVolume(child, &tpl)
			case "label":
				child = parseLabel(child, &tpl)
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

func parseEnvironmentVariables(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	for {
		if child.Next != nil {
			// Append to EnvironmentVariable container
			tpl.Context.ImageEnvironmentVariables = append(tpl.Context.ImageEnvironmentVariables, types.EnvironmentVariable{
				Name:        child.Next.Value,
				Default:     child.Next.Next.Value,
				Description: "TODO",
			})
			// Set child to the last used instance
			child = child.Next.Next
		} else {
			// break out of the inner loop once we've read in all environment variables
			break
		}
	}
	return child
}

func parseExpose(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	for {
		if child.Next != nil {
			containerPort, err := strconv.Atoi(child.Next.Value)
			utils.ExitOnErr(err)
			tpl.Context.ImagePorts = append(tpl.Context.ImagePorts, types.Port{
				Container:   containerPort,
				Host:        0,
				Description: "TODO",
			})
			child = child.Next
		} else {
			// break out of the inner loop once we've read in all exposes
			break
		}
	}
	return child
}

func parseVolume(child *parser.Node, tpl *types.TemplateRenderer) *parser.Node {
	for {
		if child.Next != nil {
			tpl.Context.ImageVolumes = append(tpl.Context.ImageVolumes, types.Volume{
				Container:   child.Next.Value,
				Host:        "TODO",
				Description: "TODO",
			})
			child = child.Next
		} else {
			// break out of the inner loop once we've read in all volumes
			break
		}
	}
	return child
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
	cap_next := false
	for _, item := range args {
		if cap_next == true {
			tpl.Context.ImageExpectedCaps = append(tpl.Context.ImageExpectedCaps, item)
			cap_next = false
		} else if strings.HasPrefix(item, "--cap-add=") {
			tpl.Context.ImageExpectedCaps = append(tpl.Context.ImageExpectedCaps, item[10:])
		} else if item == "--cap-add" {
			cap_next = true
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
