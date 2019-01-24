package git

import (
	"os/exec"
	"strings"
)

// Diff runs a git diff on a path
func Diff(context string, reference string, path string) ([]string, error) {
	cmd := exec.Command(
		"git", "diff",
		"--name-only",
		"--name-only",
		"--ignore-blank-lines",
		"--ignore-space-at-eol",
		reference, path)
	cmd.Dir = context
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}
