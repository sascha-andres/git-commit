package subcommands

import "path"

const gitCommitMessageHookName = "commit-msg"

func createCommitHookFilePath(gitFolderPath string) string {
	commitHookFilePath := path.Join(gitFolderPath, "hooks", gitCommitMessageHookName)

	return commitHookFilePath
}
