package cmd

import (
	"os"
	"strconv"

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

	tpl.Context.ImageLongDescription = "TODO"

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
			case "url":
				tpl.Context.ImageSeeAlso = utils.StripQuotes(child.Next.Next.Value)
			}
			child = child.Next.Next
		} else {
			// break out of the inner loop once we've read in all
			break
		}
	}
	return child
}
