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
		rp := filepath.Join(output.OutPath(), of.GetPath())
		op, err := filepath.Abs(rp)
		if err != nil {
			return nil
		}
		if err = os.MkdirAll(filepath.Dir(op), 0744); err != nil {
			return err
		}

		if err := ioutil.WriteFile(op, of.GetContent(), 0644); err != nil {
			return err
		}

		sublogger.Debug().Msgf("File %s\n%s", rp, of.GetContent())
		sublogger.Info().Msgf("Generated %s", rp)
	}

	for _, c := range g.Spec.Build.Output.Copy {
		path, err := filepath.Abs(filepath.Join(root, c))
		if err != nil {
			return nil
		}

		if content, err := ioutil.ReadFile(path); err == nil {
			rp := filepath.Join(output.OutPath(), c)
			outputPath, err := filepath.Abs(rp)
			if err != nil {
				return nil
			}

			if err := ioutil.WriteFile(outputPath, content, 0644); err != nil {
				return err
			}
			sublogger.Debug().Msgf("File %s\n%s", rp, content)
			sublogger.Info().Msgf("Copy %s", rp)
		}
	}

	return nil
}
