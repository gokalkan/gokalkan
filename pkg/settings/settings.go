package settings

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Settings domain with application settings variables
type Settings struct {
	DigitalSignaturePath string `yaml:"digital_signature_path"`
	DigitalSignaturePass string `yaml:"digital_signature_pass"`
	Port                 string `yaml:"port"`
}

// ParseYAML parses given yaml file and returns Settings instance
func ParseYAML(fileName string) (Settings, error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Settings{}, err
	}
	settings := Settings{}
	if err := yaml.Unmarshal([]byte(body), &settings); err != nil {
		return settings, err
	}
	return settings, nil
}
