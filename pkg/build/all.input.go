package build

import (
	"os"
	fp "path/filepath"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

type allInputStrategy struct {
	root string
}

func (g *allInputStrategy) GetFiles() ([]string, error) {
	f := []string{}
	err := fp.Walk(g.root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false && path != "." {
				if _, err := os.Stat(path); err == nil {
					f = append(f, path)
				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return f, nil
}

// newAllInputStrategy create new Git Gather Strategy
func newAllInputStrategy(root string, ges g.BuildInputSpec) InputStrategy {
	return &allInputStrategy{
		root: root,
	}
}

func init() {
	AddInputStrategy("all", newAllInputStrategy)
}
