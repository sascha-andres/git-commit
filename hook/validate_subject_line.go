package hook

import "fmt"

func (cfg *Configuration) validateSubjectLine(subjectLine string) bool {
	result := true
	if "" == subjectLine {
		fmt.Println("error: no subject line provided")
		result = false
	}
	for _, r := range cfg.subjectCompiled {
		if !r.Match([]byte(subjectLine)) {
			fmt.Println(fmt.Sprintf("error: [%s] does not match [%s]", subjectLine, r.String()))
			result = false
		}
	}
	return result
}
