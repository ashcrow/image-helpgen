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
during initialization be passing `-e VAR=VALUE` to the Docker run command.

|     Variable name        | Default |      Description                                           |
| :----------------------- | ------- | ---------------------------------------------------------- |
{{ range $_, $_data := .ImageEnvironmentVariables}}| `{{ $_data.Name }}` | `{{ $_data.Default }}`   | {{ $_data.Description}} |
{{ end }}

# SECURITY IMPLICATIONS

## Ports
|     Port Container | Port Host  |       Description             |
| :----------------- | -----------|-------------------------------|
{{ range $_, $_data := .ImagePorts }}| {{ $_data.Container }} | {{ $_data.Host }} | {{ $_data.Description }} |
{{ end }}


## Volumes
|     Volume Container | Volume Host  |       Description             |
| :----------------- | -----------|-------------------------------|
{{ range $_, $_data := .ImageVolumes }}| {{ $_data.Container }} | {{ $_data.Host }} | {{ $_data.Description }} |
{{ end }}


# SEE ALSO
{{ .ImageSeeAlso }}
