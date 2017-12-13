package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ashcrow/image-helpgen/types"
	"github.com/ashcrow/image-helpgen/utils"
)

// GuideCommand executes the logic that is exposed via the cli at
// image-helpgen guide [args]
func GuideCommand(template, basename string) error {
	tr, err := types.NewTemplateRenderer(template)
	if err != nil {
		return err
	}
	tr.Context.ImageName = ReadString(&tr, "Image Name")
	tr.Context.ImageAuthor = ReadString(&tr, "Image Author")
	tr.Context.ImageShortDescription = ReadString(&tr, "Short Description")
	tr.Context.ImageLongDescription = ReadText(&tr, "Long Description")
	tr.Context.ImageUsage = ReadText(&tr, "Image Usage")
	tr.Context.ImageSeeAlso = ReadString(&tr, "See Also")
	ReadEnvironmentVariables(&tr)
	ReadPorts(&tr)
	ReadVolumes(&tr)

	tr.Context.ImageDocDate = utils.GenerateDocDate()
	tr.Write(basename)
	return nil
}

// ReadString reads a single string and returns the result
func ReadString(t *types.TemplateRenderer, prompt string) string {
	fmt.Printf(prompt + ": ")
	result, _ := t.Reader.ReadString('\n')
	return strings.TrimSuffix(result, "\n")
}

// ReadText reads a block of text and returns the result
func ReadText(t *types.TemplateRenderer, prompt string) string {
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
func ReadEnvironmentVariables(t *types.TemplateRenderer) {
	fmt.Println("Enter Environment Variable information. Enter empty name to finish.")
	for {
		name := ReadString(t, "Name")
		if name == "" {
			break
		}
		defaultValue := ReadString(t, "Default Value")
		description := ReadString(t, "Description")
		t.Context.ImageEnvironmentVariables = append(
			t.Context.ImageEnvironmentVariables,
			types.EnvironmentVariable{
				Name:        name,
				Default:     defaultValue,
				Description: description})
	}
}

// ReadPorts reads and populates a list of Ports
func ReadPorts(t *types.TemplateRenderer) {
	fmt.Println("Enter port information. Enter empty host port to finish.")
	for {
		hp := ReadString(t, "Host Port")
		if hp == "" {
			break
		}
		cp := ReadString(t, "Container Port")
		description := ReadString(t, "Description")
		containerPort, _ := strconv.Atoi(cp)
		hostPort, _ := strconv.Atoi(hp)
		t.Context.ImagePorts = append(
			t.Context.ImagePorts,
			types.Port{
				Container:   containerPort,
				Host:        hostPort,
				Description: description})
	}
}

// ReadVolumes reads and populates a list of Volumes
func ReadVolumes(t *types.TemplateRenderer) {
	fmt.Println("Enter volume information. Enter empty host volume to finish.")
	for {
		hv := ReadString(t, "Host Volume")
		if hv == "" {
			break
		}
		cv := ReadString(t, "Container Volume")
		description := ReadString(t, "Description")
		t.Context.ImageVolumes = append(
			t.Context.ImageVolumes,
			types.Volume{Container: cv, Host: hv, Description: description})
	}
}
