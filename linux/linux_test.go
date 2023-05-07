package linux

import (
	"testing"
)

func TestListContainers(t *testing.T) {
	_, err := ListContainers()

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAllContainerIDs(t *testing.T) {
	_, err := GetAllContainerIDs()

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStopALLContainers(t *testing.T) {
	ids, err := GetAllContainerIDs()
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = StopAllContainers(ids)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestStopSingleContainer(t *testing.T) {
	_, err := StopSingleContainer("463631feb379")

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestBuildDockerFile(t *testing.T) {
	var path string = "/home/princewillo/code/Personal/Docker/Test/Dockerfile"
	_, err := BuildDockerFile(NewContainer(path, "test", "v1"))

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRunContainer(t *testing.T) {
	params := map[string]string{
		"name":     "test",
		"tag":      "v1",
		"in_port":  "8000",
		"out_port": "8000",
	}
	_, err := RunContainer(params)

	if err != nil {
		t.Errorf(err.Error())
	}
}
