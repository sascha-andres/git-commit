# git-commit

This project is a configurable commit-msg hook. You can either configure using a global configuration file of a repository configuration file. If both exist, repository file will be merged into global configuration file.

The global configuration file is located in the HOME directory of the user and named `.commit-hook.yaml` while the repository specific file is in the root of the repository and named `.commit-hook.yaml`.

Both files have the same structure.

## Configuration file structure

    ---
    
    subject-line-length: 50              # Length of the subject line at a maximum
    body-required: false                 # Is a body required for a commit
    separate-body: true                  # Should the body be separated by a blank line
    body-line-length: 72                 # Maximum length of body lines
    enforce-body-line-length: true       # Should too long body lines be treated as an error or as a warning
    
    ignore:                              # a list of regular expressions to ignore lines (no check)
      - ^#.*                             # ignore comments
    
    subject:                             # Match the subject against those expressions
      - ^[a-z]+(\([a-z]+\))?:.*[^\.]$    # have feat(web): bla match, no . at the end allowed
    
    occurs:                              # Match somewhere, check for existence
      - TICKET-[0-9]+                    # have a TICKET-1 as a match

## Installation

Put the binary into your path.

## Use

### install hook

The application has a helper to install the hook:

    git-commit-hook install

You need to run this in the root of your project ( the folder containing the `.git` folder ). If there is an existing hook you have to force it ( `-f` ).

### uninstall hook

Similar to the installation process, there is an uninstall helper:

    git-commit-hook uninstall

Essentially this just removes the hook, so it would remove any other hook also. So be careful.

## History

|Version|Description|
|---|---|
|0.2.0|configuration file is now named `.commit-hook.yaml`|
||fix: project configuration now correctly located in repository root|
||fix: run subject line length|
|0.1.0|Initial version|