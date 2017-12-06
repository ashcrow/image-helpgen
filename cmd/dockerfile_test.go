package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDockerfileCommand(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("", "")
	expectedMdFile := tmpFile.Name() + ".md"
	defer os.Remove(tmpFile.Name())
	defer os.Remove(expectedMdFile)
	DockerfileCommand("../example/Dockerfile", "../template.tpl", tmpFile.Name())
}
