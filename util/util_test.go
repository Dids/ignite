package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestValidate(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get current working directory: %s", err)
	}

	config, err := ioutil.ReadFile(path.Join(path.Dir(cwd), "ignition.test.json"))
	if err != nil {
		t.Errorf("Failed to open Ignition config: %s", err)
	}

	config = bytes.Replace(config, []byte("\r"), []byte{}, -1)

	err = Validate(config)
	if err != nil {
		t.Errorf("Failed to validate Ignition config: %s", err)
	}

	// log.Println(string(config))
}
