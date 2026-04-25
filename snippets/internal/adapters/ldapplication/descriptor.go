package ldapplication

import (
	"os"

	"gopkg.in/yaml.v3"
)

// descriptor is the subset of sdk.yaml that the ld-application adapter
// cares about.
type descriptor struct {
	ID            string `yaml:"id"`
	LDApplication struct {
		GetStartedFile string `yaml:"get-started-file"`
	} `yaml:"ld-application"`
}

func loadDescriptor(path string) (*descriptor, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var d descriptor
	if err := yaml.Unmarshal(raw, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
