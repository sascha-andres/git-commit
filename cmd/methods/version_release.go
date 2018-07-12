package methods

import "fmt"

var (
	versionNumber = "development"
	commit        = "unset"
	date          = "unset"
)

func PrintVersion() {
	fmt.Println(fmt.Sprintf("git-commit-hook version %s", versionNumber))
	fmt.Println(fmt.Sprintf("compiled from %s on %s", commit, date))
	fmt.Println()
}
