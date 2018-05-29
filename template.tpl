% {{ .ImageName }}(2) Container Image Pages
% {{ .ImageAuthor }}
% {{ .ImageDocDate }}

# NAME
{{ .ImageName }} - {{ .ImageShortDescription}}

# DESCRIPTION
{{ .ImageLongDescription }}


# USAGE
{{ .ImageUsage }}

# ENVIRONMENT VARIABLES

The image recognizes the following environment variables that you can set
during initialization by passing `-e VAR=VALUE` to the `docker run` command.

|     Variable name        | Default |      Description                                           |
| :----------------------- | ------- | ---------------------------------------------------------- |
{{ range $_, $_data := .ImageEnvironmentVariables}}| `{{ $_data.Name }}` | `{{ $_data.Default }}`   | {{ $_data.Description}} |
{{ end }}

# SECURITY IMPLICATIONS
The following sections describe potential security issues related to how the container image was designed to run.

## Ports

Exposed TCP (default) or UDP ports that the container listens on at runtime include the following:

|     Port Container | Port Host  |       Description             |
| :----------------- | -----------|-------------------------------|
{{ range $_, $_data := .ImagePorts }}| {{ $_data.Container }} | {{ $_data.Host }} | {{ $_data.Description }} |
{{ end }}


## Volumes

Directories that are mounted from the host system to a mount point inside the container include the following:

|     Volume Container | Volume Host  |       Description             |
| :----------------- | -----------|-------------------------------|
{{ range $_, $_data := .ImageVolumes }}| {{ $_data.Container }} | {{ $_data.Host }} | {{ $_data.Description }} |
{{ end }}

{{ if .ImageExpectedDaemon }}## Daemon
This image is expected to be run as a daemon{{ end }}
{{ if .ImageExpectedCaps }}

## Expected Capabilities

This container needs to open one or more Linux capabilities (see `man capabilities 7`) to the host computer. The following capababilities (added with the \-\-cap\-add option) are expected:

{{ range $_, $_cap := .ImageExpectedCaps }}
- {{ $_cap }}{{ end }}
{{ end}}

# SEE ALSO
{{ .ImageSeeAlso }}
