package build

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kesslerdev/gitgen/pkg/build"
	"github.com/kesslerdev/gitgen/pkg/generator"
	"github.com/rs/zerolog/log"
)

func BuildGenerator(root string, g *generator.Generator) error {
	sublogger := log.With().
		Str("build", g.Metadata.Name).
		Logger()
	if g.Spec.Build.Input.Strategy == "" {
		return errors.New("Cannot build generator with no strategy")
	}
	input, err := build.NewInputStrategy(root, g.Spec.Build.Input)
	if err != nil {
		return err
	}
	sublogger.Info().Msgf("Using input strategy %s", g.Spec.Build.Input.Strategy)

	output, err := build.NewOutputStrategy(root, *g)
	if err != nil {
		return err
	}
	sublogger.Info().Msgf("Using output strategy %s", g.Spec.Build.Output.Strategy)
	files, err := input.GetFiles()
	if err != nil {
		return err
	}
	for _, f := range files {
		fp, err := filepath.Rel(root, f)
		if err != nil {
			return err
		}
		sublogger.Debug().Msgf("Working on file %s(%s)", f, fp)
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		of, err := output.BuildFile(fp, content)
		if err != nil {
			return err
		}

		if err = os.MkdirAll(filepath.Dir(of.GetPath()), 0744); err != nil {
			return err
		}

		if err := ioutil.WriteFile(of.GetPath(), of.GetContent(), 0644); err != nil {
			return err
		}

		sublogger.Debug().Msgf("File %s\n%s", of.GetPath(), of.GetContent())
		sublogger.Info().Msgf("Generated %s", of.GetPath())
	}

	return nil
}
