package linux

import (
	"bytes"
	"os/exec"
)

func ListContainers() (containers string, err error) {
	cmd := exec.Command("docker", "ps")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}

func GetAllContainerIDs() (message string, err error) {
	cmd := exec.Command("docker", "ps", "-a", "-q")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}
