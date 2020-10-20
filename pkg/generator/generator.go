package generator

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// KMetadata represent a Kubernetes likes metadata struct
type KMetadata struct {
	Name string
}

// KObject represent a Kubernetes likes object
type KObject struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   KMetadata
	Spec       interface{} `yaml:"-"`
}

// GeneratorFileFormat is the format used for generator file name
const GeneratorFileFormat = "%s.gitgen.yaml"

// GeneratorFileSuffix is the suffix used for generator file name
const GeneratorFileSuffix = "gitgen.yaml"

// Generator represent a generator config object
type Generator struct {
	KObject `yaml:",inline"`
	Spec    Spec
}

// Spec represent a generator specification
type Spec struct {
	Build BuildSpec
}

// BuildSpec represent a build related generator specification
type BuildSpec struct {
	Input     BuildInputSpec
	Replacers []BuildReplacerSpec
	Output    BuildOutputSpec
}

// BuildInputSpec represent a build input related generator specification
type BuildInputSpec struct {
	Strategy string
	Options  interface{} `yaml:",omitempty"`
}

// BuildReplacerSpec represent a build replacer related generator specification
type BuildReplacerSpec struct {
	Find  string
	When  string
	As    string
	Cases bool
}

// BuildOutputSpec represent a build output related generator specification
type BuildOutputSpec struct {
	Strategy string
	Copy     []string
	Options  interface{} `yaml:",omitempty"`
}

// NewGenerator create a generator config object
func NewGenerator(name string) *Generator {
	return &Generator{
		KObject: KObject{
			APIVersion: "gitgen.kesslerdev.io/v1",
			Kind:       "Generator",
			Metadata: KMetadata{
				Name: name,
			},
		},
		Spec: Spec{
			Build: BuildSpec{
				Input: BuildInputSpec{
					Strategy: "git",
				},
				Output: BuildOutputSpec{
					Strategy: "hygen",
				},
			},
		},
	}
}

// FromYamlFile load generator struct form yaml file
func FromYamlFile(file string, g *Generator) error {
	if _, err := os.Stat(file); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, &g); err != nil {
		return err
	}

	return nil
}
