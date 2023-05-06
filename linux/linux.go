package linux

import (
	"bytes"
	"os/exec"
)

func ListContainers() (containers string, error error) {
	cmd := exec.Command("docker", "ps")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}
