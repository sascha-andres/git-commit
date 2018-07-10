package hook

import "fmt"

func (cfg *Configuration) validateBodySeparation(commitMessage []string) bool {
	result := true
	if cfg.BodyRequired {
		if len(commitMessage) == 1 {
			fmt.Println("error: body required but not provided")
			result = false
		} else {
			if len(commitMessage) == 2 {
				if cfg.SeparateBody {
					fmt.Println("error: body should be separated")
					result = false
				}
			}
			if cfg.SeparateBody {
				if commitMessage[1] != "" {
					fmt.Println("error: body should be separated")
					result = false
				}
			}
		}
	}
	if cfg.SeparateBody && len(commitMessage) > 1 {
		if len(commitMessage) == 2 {
			if cfg.SeparateBody {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
		if cfg.SeparateBody {
			if commitMessage[1] != "" {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
	}
	return result
}
