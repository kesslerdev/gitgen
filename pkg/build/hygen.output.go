package build

import (
	"fmt"
	"os"
	"strings"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

type hygenOutputSpec struct {
	Options hygenOutputOptions
}

type hygenOutputOptions struct {
	ParentDir bool `yaml:"parentDir"`
}
type hygenOutputStrategy struct {
	root      string
	generator g.Generator
}

func hygenReplacer(i *ReplacerInfos) string {
	switch i.Case {
	case "upper":
		return fmt.Sprintf("<%%= %s.toUpperCase() %%>", i.Var)
	case "lower":
		return fmt.Sprintf("<%%= %s.toLowerCase() %%>", i.Var)
	case "camel":
		return fmt.Sprintf("<%%= h.inflection.camelize(%s, true) %%>", i.Var)
	case "camel_up":
		return fmt.Sprintf("<%%= h.inflection.camelize(%s, false) %%>", i.Var)
	case "underscore":
		return fmt.Sprintf("<%%= h.inflection.underscore(%s); %%>", i.Var)
	case "constant":
		return fmt.Sprintf("<%%= h.changeCase.constant(%s); %%>", i.Var)
	case "kebab":
		return fmt.Sprintf("<%%= h.inflection.dasherize(%s); %%>", i.Var)
	case "kebab_uppercase":
		return fmt.Sprintf("<%%= h.inflection.dasherize(%s).toUpperCase(); %%>", i.Var)
	default:
		return fmt.Sprintf("<%%= %s %%>", i.Var)
	}
}

func (g *hygenOutputStrategy) OutPath() string {
	return fmt.Sprintf("_templates/%s/new", g.generator.Metadata.Name)
}

func (g *hygenOutputStrategy) BuildFile(output string, content []byte) (OutputFile, error) {
	orig := output

	for _, r := range g.generator.Spec.Build.Replacers {
		content = ApplyReplacer(&r, content, hygenReplacer)
	}

	for _, r := range g.generator.Spec.Build.Replacers {
		output = string(ApplyReplacer(&r, []byte(output), hygenReplacer))
	}

	if g.generator.Spec.Build.Output.Options.(hygenOutputOptions).ParentDir {
		output = fmt.Sprint(hygenReplacer(&ReplacerInfos{
			Case: "exact",
			Var:  "name",
		}), string(os.PathSeparator), output)
	}

	return &hygenOutputFile{
		originalPath: orig,
		path:         output,
		content:      content,
	}, nil
}

type hygenOutputFile struct {
	originalPath string
	path         string
	content      []byte
}

func (g *hygenOutputFile) GetPath() string {
	return fmt.Sprintf("%s.ejs.t", strings.Replace(g.originalPath, string(os.PathSeparator), "_", -1))
}

const hygenFileFormat = `---
to: %s
---
%s`

func (g *hygenOutputFile) GetContent() []byte {
	return []byte(fmt.Sprintf(hygenFileFormat, g.path, g.content))
}

func newHygenOutputStrategy(root string, g g.Generator) OutputStrategy {
	return &hygenOutputStrategy{
		root:      root,
		generator: g,
	}
}

func init() {
	AddOutputStrategy("hygen", newHygenOutputStrategy)
	g.AddOutputOptionUnmarshalSpec("hygen", func(unmarshal func(interface{}) error) interface{} {
		h := &hygenOutputSpec{}
		if err := unmarshal(&h); err != nil {
			panic(err)
		}

		return h.Options
	})
}
