package build

import (
	"fmt"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

// OutputFile is the generated template
type OutputFile interface {
	GetPath() string
	GetContent() []byte
}

// OutputStrategy is used to generate files for specific template runner
type OutputStrategy interface {
	BuildFile(output string, content []byte) (OutputFile, error)
}

// OutputStrategyCreator is used to create new OutputStrategy
type OutputStrategyCreator func(root string, g g.Generator) OutputStrategy

var outputStrategies = map[string]OutputStrategyCreator{}

// AddOutputStrategy add new Output Strategy to the system
func AddOutputStrategy(name string, f OutputStrategyCreator) {
	outputStrategies[name] = f
}

// NewOutputStrategy create new Output strategy
func NewOutputStrategy(root string, g g.Generator) (OutputStrategy, error) {
	output := outputStrategies[g.Spec.Build.Output.Strategy]
	if output == nil {
		return nil, fmt.Errorf("Output Strategy %s unknown", g.Spec.Build.Output.Strategy)
	}

	return output(root, g), nil
}
