package util

import (
	"log"

	ignConfig "github.com/coreos/ignition/config/v2_2"
)

// Validate verifies the structure of the Ignition config
func Validate(config []byte) error {
	// TODO: Implement
	log.Println("Validate()")

	cfg, report, err := ignConfig.Parse(config)
	if err != nil {
		return err
	}

	log.Println("Cfg:", cfg)
	log.Println("Report:", report)

	return nil
}

// Merge combines together two or more Ignition config files
func Merge() {
	// TODO: Implement
	log.Println("Merge()")
}
