package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// WriteNewYaml write a new yaml file
func WriteNewYaml(path string, obj interface{}) error {

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("File at %s exists", path)
	}

	if err := WriteYaml(path, obj); err != nil {
		return fmt.Errorf("Unable to write file at %s: %s", path, err.Error())
	}

	return nil
}

// WriteYaml write a struct to file as yaml
func WriteYaml(path string, obj interface{}) error {

	d, err := yaml.Marshal(&obj)
	if err != nil {
		return errors.New("unable to marshal yaml")
	}

	if err := ioutil.WriteFile(path, d, 0644); err != nil {
		return fmt.Errorf("Unable to write file at %s: %s", path, err.Error())
	}

	return nil
}
