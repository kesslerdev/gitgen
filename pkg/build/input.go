package build

import (
	"fmt"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

// InputStrategy is used to find files for generation
type InputStrategy interface {
	GetFiles() ([]string, error)
}

// InputStrategyCreator is used to create new InputStrategy
type InputStrategyCreator func(root string, ges g.BuildInputSpec) InputStrategy

var inputStrategies = map[string]InputStrategyCreator{}

// AddInputStrategy add new Input Strategy to the system
func AddInputStrategy(name string, f InputStrategyCreator) {
	inputStrategies[name] = f
}

// NewInputStrategy create new Input strategy
func NewInputStrategy(root string, ges g.BuildInputSpec) (InputStrategy, error) {
	input := inputStrategies[ges.Strategy]
	if input == nil {
		return nil, fmt.Errorf("Input Strategy %s unknown", ges.Strategy)
	}

	return input(root, ges), nil
}
