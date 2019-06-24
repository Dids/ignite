package util

import (
	"log"
	validate "github.com/coreos/ignition/config/validate"
)

// Validate verifies the structure of the Ignition config
func Validate() {
	log.Println("Validate()")
	report := validate.ValidateWithContext(nil, nil)
	log.Println(report)
}
