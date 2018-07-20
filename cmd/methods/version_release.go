package methods

import "fmt"

var (
	versionNumber = "development"
	commit        = "unset"
	date          = "unset"
)

// PrintVersion writes out the version number to stdout
func PrintVersion() {
	fmt.Println(fmt.Sprintf("git-hook-commit version %s", versionNumber))
	fmt.Println(fmt.Sprintf("compiled from %s on %s", commit, date))
	fmt.Println()
}
