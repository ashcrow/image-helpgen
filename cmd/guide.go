package cmd

import (
	"github.com/ashcrow/image-helpgen/types"
	"github.com/ashcrow/image-helpgen/utils"
)

// GuideCommand executes the logic that is exposed via the cli at
// image-helpgen guide [args]
func GuideCommand(template, basename string) {
	tr := types.NewTemplateRenderer(template)
	tr.Context.ImageName = tr.ReadString("Image Name")
	tr.Context.ImageAuthor = tr.ReadString("Image Author")
	tr.Context.ImageShortDescription = tr.ReadString("Short Description")
	tr.Context.ImageLongDescription = tr.ReadText("Long Description")
	tr.Context.ImageUsage = tr.ReadText("Image Usage")
	tr.Context.ImageSeeAlso = tr.ReadString("See Also")
	tr.ReadEnvironmentVariables()
	tr.ReadPorts()
	tr.ReadVolumes()

	tr.Context.ImageDocDate = utils.GenerateDocDate()
	tr.Write(basename)
}
