package build

import (
	"os"
	fp "path/filepath"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

type staticInputStrategy struct {
	root  string
	files []string
}

func (g *staticInputStrategy) GetFiles() ([]string, error) {
	f := []string{}
	for _, file := range g.files {
		path, err := fp.Abs(fp.Join(g.root, file))
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(path); err == nil {
			f = append(f, path)
		}

	}

	return f, nil
}

// NewStaticInputStrategy create new Static Gather Strategy
func newStaticInputStrategy(root string, ges g.BuildInputSpec) InputStrategy {
	return &staticInputStrategy{
		root:  root,
		files: ges.Files,
	}
}

func init() {
	AddInputStrategy("static", newStaticInputStrategy)
}
