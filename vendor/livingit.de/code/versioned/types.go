package versioned

type (
	// configuration just loads a version field
	// from a configuration
	configuration struct {
		Version string `yaml:"version" json:"version"`
	}

	// VersionReader is a bucket for config processing
	VersionReader struct {
		YAML YAML
		JSON JSON
	}

	// YAML is a bucket to process YAML files
	YAML struct { }

	// JSON is a bucket to process JSON files
	JSON struct { }
)

// NewVersionReader creates a new version reader object
func NewVersionReader() *VersionReader {
	return &VersionReader{
		YAML: YAML{},
		JSON: JSON{},
	}
}
