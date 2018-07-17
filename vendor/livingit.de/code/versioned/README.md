# versioned

a little library to load a version attribute from a YAML or JSON file

## Usage

    package main

    import "livingit.de/code/versioned"
    import "fmt"

    const yamlContent := `---

    version: 1.0.0`

    func main() {
      rdr := versioned.NewVersionReader()
      version, err := rdr.YAML.GetVersion([]byte(yamlContent))
      if err != nil {
        panic(err)
      }
      fmt.Printf("Read version is: %s\n", version)
    }

## History

|Version|Description|
|---|---|
|1.0.0|Initial release|
