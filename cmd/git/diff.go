package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Diff runs a git diff on a path
func Diff(context string, reference string, path string) ([]string, error) {
	cmd := exec.Command(
		"git",
		"diff",
		"--name-only",
		"--ignore-blank-lines",
		"--ignore-space-at-eol",
		reference,
		"--",
		path,
	)
	cmd.Dir = context
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, string(out))
		fmt.Fprintf(os.Stderr, fmt.Sprint(err))
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}
