package linux

import "path/filepath"

func GetPath(file string) (path string, err error) {
	absPath, err := filepath.Abs(file)
	if err != nil {
		return
	}
	return filepath.Dir(absPath), nil
}
