package linux

import (
	"bytes"
	"os/exec"
	"strings"
)

type Container struct {
	Path   string
	Name   string
	Tag    string
	Volume string
}

type Option func(f *Container)

func WithVolume(volume string) Option {
	return func(f *Container) {
		f.Volume = volume
	}
}

func NewContainer(path string, name string, tag string, opts ...Option) *Container {
	container := &Container{
		Path: path,
		Name: name,
		Tag:  tag,
	}

	for _, opt := range opts {
		opt(container)
	}
	return container
}

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

func StopAllContainers(ids []string) (msg string, err error) {
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

func StopSingleContainer(id string) (msg string, err error) {
	cmd := exec.Command("docker", "stop", id)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}

func BuildDockerFile(build *Container) (msg string, err error) {
	var name_tag string = build.Name + ":" + build.Tag
	path, err := GetPath(build.Path)
	if err != nil {
		return
	}
	cmd := exec.Command("docker", "build", "-t", name_tag, path)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}

func RunContainer(params map[string]string) (msg string, err error) {
	var name_tag string = params["name"] + ":" + params["tag"]
	commands := []string{}

	switch {
	case params["out_port"] != "":
		var port_param string = params["out_port"] + ":" + params["in_port"]
		commands = append(commands, "run", "-p", port_param)
		fallthrough
	default:
		commands = append(commands, name_tag)
	}

	cmd := exec.Command("docker", commands...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return
	}
	return out.String(), nil
}
