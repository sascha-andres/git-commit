package versioned

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// GetFromFile loads a version out of a file
func (j *JSON) GetFromFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	return j.GetFromReader(f)
}

// GetFromReader reads everything from a reader and returns the version
func (j *JSON) GetFromReader(reader io.Reader) (string, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return j.GetVersion(content)
}

// GerVersion reads everything from a byte array and returns the version
func (j *JSON) GetVersion(content []byte) (string, error) {
	var cfg configuration
	err := json.Unmarshal(content, &cfg)
	if err != nil {
		return "", err
	}
	return cfg.Version, nil
}
