package util

import (
	"log"

	validate "github.com/coreos/ignition/config/validate"
)

// Validate verifies the structure of the Ignition config
func Validate() {
	// TODO: Implement
	log.Println("Validate()")
	report := validate.ValidateConfig(nil, nil)
	log.Println(report)
}

// Merge combines together two or more Ignition config files
func Merge() {
	// TODO: Implement
	log.Println("Merge()")
}
