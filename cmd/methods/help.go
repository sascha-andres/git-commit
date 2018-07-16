package methods

import "fmt"

// Help prints out usage to stdout
func Help() {
	fmt.Println("too few arguments")
	fmt.Println("")
	fmt.Println("install 	- helps to install git-commit-hook")
	fmt.Println("uninstall 	- helps to uninstall git-commit-hook")
}
