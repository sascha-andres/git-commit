package v2

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func runLinting(cfg Configuration, name string, res bool) {
	It(fmt.Sprintf("should return %v for \"%v\"", res, name), func() {
		Expect(cfg.validateGitLabCI()).To(Equal(res))
	})
}

var _ = Describe("GitLab CI linter", func() {
	runLinting(Configuration{LintGitLabCI: false}, "linter disabled", true)
	runLinting(Configuration{LintGitLabCI: true}, "linter enabled || not yet created", true)
	//	runLinting(Configuration{LintGitLabCI: true, GitLabCIFile: "test.yml"}, "linter enabled || other config file", false)
	// runLinting(Configuration{LintGitLabCI: true, GitLabCIFile: "testdata/empty.yml"}, "linter enabled || empty file", true)
})

func TestGitLabCILinter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitLab CI linter")
}
