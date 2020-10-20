package build

import (
	"os/exec"
	"strings"
	fp "path/filepath"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

type gitInputStrategy struct {
	root string
}

func (g *gitInputStrategy) GetFiles() ([]string, error) {
	f := []string{}

	cmd := exec.Command("git", "ls-files")
	cmd.Dir = g.root
	out, err := cmd.Output()
	if  err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		if line != "" {
			path, err := fp.Abs(fp.Join(g.root, line))
			if err != nil {
				return nil, err
			}

			f = append(f, path)
		}
	}

	return f, nil
}

// newGitInputStrategy create new Git Gather Strategy
func newGitInputStrategy(root string, ges g.BuildInputSpec) InputStrategy {
	return &gitInputStrategy{
		root: root,
	}
}

func init() {
	AddInputStrategy("git", newGitInputStrategy)
}
