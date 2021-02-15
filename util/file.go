package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
)

// FileExists checks if a file exists
func FileExists(p string) (bool, error) {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// ReadFile reads a file
func ReadFile(file string, out interface{}) (err error) {
	var content []byte

	content, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	switch ext := strings.ToLower(filepath.Ext(file)); ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, out)
		return

	case ".json":
		err = json.Unmarshal(content, out)
		return

	default:
		err = fmt.Errorf("Invalid file encoding. Expected yaml or json, got %v", ext)
		return
	}
}
