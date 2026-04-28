package ldapplication

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"
)

// descriptor models the complete sdk.yaml schema (not just the
// ld-application subset) so we can decode with KnownFields(true) and
// catch typos like `Entrpoint:` or `get-startted-file:` rather than
// silently dropping them. Fields we don't read yet are still listed so
// the strict decode doesn't reject them.
type descriptor struct {
	ID              string         `yaml:"id"`
	SDKMetaID       string         `yaml:"sdk-meta-id"`
	DisplayName     string         `yaml:"display-name"`
	Type            string         `yaml:"type"`
	Languages       []descLanguage `yaml:"languages"`
	PackageManagers []string       `yaml:"package-managers"`
	Regions         []string       `yaml:"regions"`
	HelloWorldRepo  string         `yaml:"hello-world-repo"`
	LDApplication   struct {
		GetStartedFile string `yaml:"get-started-file"`
	} `yaml:"ld-application"`
	Docs struct {
		ReferencePage string `yaml:"reference-page"`
	} `yaml:"docs"`
}

type descLanguage struct {
	ID         string   `yaml:"id"`
	Extensions []string `yaml:"extensions"`
}

func loadDescriptor(path string) (*descriptor, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var d descriptor
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	if err := dec.Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
