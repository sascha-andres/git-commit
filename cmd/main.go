package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"path/filepath"

	"bufio"

	"github.com/imdario/mergo"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"livingit.de/code/git-commit/hook"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("too few arguments")
		fmt.Println("")
		fmt.Println("install 	- helps to install git-commit-hook")
		fmt.Println("uninstall 	- helps to uninstall git-commit-hook")
		os.Exit(0)
		return
	}

	commitMessageFile := os.Args[1]
	if commitMessageFile == "" {
		fmt.Println(errors.New("no commit message file passed as parameter 1"))
		os.Exit(1)
		return
	}

	if !fileExists(commitMessageFile) {
		fmt.Println(errors.New("passed commit message file not found"))
		os.Exit(1)
		return
	}

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config *hook.Configuration
	globalConfig := fmt.Sprintf("%s/.git-commit.yaml", home)
	if fileExists(globalConfig) {
		data, err := ioutil.ReadFile(globalConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		var cfg hook.Configuration
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		config = &cfg
	}

	projectPath, err := filepath.Abs(filepath.Dir(commitMessageFile))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	fmt.Println(projectPath)

	localConfig := fmt.Sprintf("%s/git-commit.yaml", projectPath)
	if fileExists(localConfig) {
		data, err := ioutil.ReadFile(localConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		var cfg hook.Configuration
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		if nil == config {
			config = &cfg
		} else {
			if err := mergo.Merge(config, cfg); err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
		}
	}

	if nil == config {
		fmt.Println(errors.New("no suitable configuration found"))
	}

	file, err := os.Open(commitMessageFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	scanner := bufio.NewScanner(file)
	commitFileContent := make([]string, 0)
	for scanner.Scan() {
		commitFileContent = append(commitFileContent, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	ok, err := config.Validate(commitFileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	if !ok {
		os.Exit(1)
		return
	}

	os.Exit(0)
}

// fileExists implements a lazy way to check for a file
func fileExists(path string) bool {
	if stat, err := os.Stat(path); err == nil || os.IsExist(err) {
		if stat.IsDir() {
			return false
		}
		return true
	}
	return false
}
