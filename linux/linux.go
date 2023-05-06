package linux

import (
	"bytes"
	"os/exec"
	"strings"
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

func GetAllContainerIDs() (ids []string, err error) {
	cmd := exec.Command("docker", "ps", "-a", "-q")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	var sb strings.Builder
	sb.WriteString(out.String())
	container_ids := strings.Split(sb.String(), "\n")
	return container_ids, nil
}

func StopAllContainers(ids []string) (message string, err error) {
	var commands = append([]string{"stop"}, ids...)
	commands = commands[0:(len(commands) - 1)]
	cmd := exec.Command("docker", commands...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}
