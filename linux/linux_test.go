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
