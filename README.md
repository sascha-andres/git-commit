# git-commit

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/git-commit)](https://goreportcard.com/report/github.com/sascha-andres/git-commit) [![Maintainability](https://api.codeclimate.com/v1/badges/89c3d725bc80fe513869/maintainability)](https://codeclimate.com/github/sascha-andres/git-commit/maintainability) [![codebeat badge](https://codebeat.co/badges/e26b8c40-1ada-41b7-82c6-4dd7e96edde0)](https://codebeat.co/projects/github-com-sascha-andres-git-commit-master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsascha-andres%2Fgit-commit.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsascha-andres%2Fgit-commit?ref=badge_shield)

This project is a configurable commit-msg hook. You can either configure using a global configuration file of a repository configuration file. If both exist, repository file will be merged into global configuration file.

The global configuration file is located in the HOME directory of the user and named `.commit-hook.yaml` while the repository specific file is in the root of the repository and named `.commit-hook.yaml`.

Both files have the same structure.

## Configuration file structure

    ---
    
    version: 2                           # version of config file
    
    subject-line-length: 50              # Length of the subject line at a maximum
    body-required: false                 # Is a body required for a commit
    separate-body: true                  # Should the body be separated by a blank line
    body-line-length: 72                 # Maximum length of body lines
    enforce-body-line-length: true       # Should too long body lines be treated as an error or as a warning
    lint_gitlab_ci: true                 # If true this calls out to lint .gitlab-ci.yml files
    gitlab_ci_file: relativePath         # Set a different CI file location
    
    ignore:                              # a list of regular expressions to ignore lines (no check)
      - ^#.*                             # ignore comments
    
    subject:                                         # Match the subject against those expressions
      - name: ensure tagging of subject              # Name of rule
        expression: ^[a-z]+(\([a-z]+\))?:.*[^\.]$    # have feat(web): bla match, no . at the end allowed
        severity: error                              # on error fail commit, else print with severity attached
    
    occurs:                              # Match somewhere, check for existence
      - TICKET-[0-9]+                    # have a TICKET-1 as a match

    external-tools:                      # external tools to run before accepting commit
      - severity: error                  # on error fail commit, else print with severity attached
        name: linter                     # name of rule
        command:                         # command to execute
          - make
          - lint


## Installation

Put the binary into your path. If you want to use GitLab CI linting you need to use: https://gitlab.com/orobardet/gitlab-ci-linter

## Use

### install hook

The application has a helper to install the hook:

    git-hook-commit install

You need to run this in the root of your project ( the folder containing the `.git` folder ). If there is an existing hook you have to force it ( `-f` ).

### uninstall hook

Similar to the installation process, there is an uninstall helper:

    git-hook-commit uninstall

Essentially this just removes the hook, so it would remove any other hook also. So be careful.

## History

|Version|Description|
|---|---|
|1.0.0|update to use go modules|
|0.9.0|add support for GitLab CI linter|
||fix a bug where body lines where multiplied by the number of ignore expressions|
|0.8.0|add support for Co-authored-by|
|0.7.0|rename binary|
|0.6.2|code quality improvements|
|0.6.1|code quality improvements|
|0.6.0|remove v1|
||run external tools|
||fix: version check|
|0.5.0|version 2 of configuration|
||Naming of rules|
||Severity for rules|
|0.4.0|add version to config file|
|0.3.1|integration with code quality tools|
|0.3.0|print version information|
|0.2.0|configuration file is now named `.commit-hook.yaml`|
||fix: project configuration now correctly located in repository root|
||fix: run subject line length|
|0.1.0|Initial version|


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsascha-andres%2Fgit-commit.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsascha-andres%2Fgit-commit?ref=badge_large)
