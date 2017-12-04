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
	"flag"
	"fmt"
	"os"

	"github.com/ashcrow/image-helpgen/cmd"
	"github.com/ashcrow/image-helpgen/utils"
)

const defaultTemplate = "/etc/image-helpgen/template.tpl"

func printHelp() {
	fmt.Printf("Usage: %s <command> [args]\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  guide: Asks for input and builds markdown and man output")
	fmt.Println("  dockerfile: Parses a Dockerfile and generates a markdown template")
	fmt.Println("  man: Generate man page off of a previouslty filled out markdown template")
}

// main function for CLI
func main() {
	var template string
	var basename string
	var dockerfilePath string

	// Setup subcommand parsers
	guideCmd := flag.NewFlagSet("guide", flag.ExitOnError)
	guideCmd.StringVar(
		&template, "template", defaultTemplate, "Template to use when rendering")
	guideCmd.StringVar(
		&basename, "basename", "help", "Base name to use for file writing")

	dockerfileCmd := flag.NewFlagSet("dockerfile", flag.ExitOnError)
	dockerfileCmd.StringVar(
		&template, "template", defaultTemplate, "Template to use when rendering")
	dockerfileCmd.StringVar(
		&basename, "basename", "help", "Base name to use for file writing")
	dockerfileCmd.StringVar(
		&dockerfilePath, "dockerfile", "Dockerfile",
		"Full path to the Dockerfile to read")

	manCmd := flag.NewFlagSet("man", flag.ExitOnError)
	manCmd.StringVar(
		&basename, "basename", "help", "Base name to use for file writing")

	// If we have no subcommand then print help and exit
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(1)
	}

	// Otherwise pass off to the subcommand defaulting to help if
	// the command is not valid
	switch os.Args[1] {
	case "guide":
		guideCmd.Parse(os.Args[2:])
		cmd.GuideCommand(template, basename)
	case "dockerfile":
		dockerfileCmd.Parse(os.Args[2:])
		cmd.DockerfileCommand(dockerfilePath, template, basename)
	case "man":
		manCmd.Parse(os.Args[2:])
		utils.WriteManFromMd(basename)
	default:
		printHelp()
		fmt.Printf("Error: %s is not a valid command\n", os.Args[1])
		os.Exit(1)
	}
}
