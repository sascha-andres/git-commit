package versioned

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

// GetFromFile loads a version out of a file
func (y *YAML) GetFromFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	return y.GetFromReader(f)
}

// GetFromReader reads everything from a reader and returns the version
func (y *YAML) GetFromReader(reader io.Reader) (string, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return y.GetVersion(content)
}

// GerVersion reads everything from a byte array and returns the version
func (y *YAML) GetVersion(content []byte) (string, error) {
	var cfg configuration
	err := yaml.Unmarshal(content, &cfg)
	if err != nil {
		return "", err
	}
	return cfg.Version, nil
}
