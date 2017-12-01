package cmd

import (
	"fmt"
	"time"

	"github.com/ashcrow/image-helpgen/types"
)

// GuideCommand guides the user to fill out sections of the template
func GuideCommand(tr *types.TemplateRenderer) {
	tr.Context.ImageName = tr.ReadString("Image Name")
	tr.Context.ImageAuthor = tr.ReadString("Image Author")
	tr.Context.ImageShortDescription = tr.ReadString("Short Description")
	tr.Context.ImageLongDescription = tr.ReadText("Long Description")
	tr.Context.ImageUsage = tr.ReadText("Image Usage")
	tr.Context.ImageSeeAlso = tr.ReadString("See Also")
	tr.ReadEnvironmentVariables()
	tr.ReadPorts()
	tr.ReadVolumes()

	// Set the doc date to now
	now := time.Now()
	tr.Context.ImageDocDate = fmt.Sprintf("%s %d", now.Month(), now.Year())
}
