package build

import (
	"testing"

	g "github.com/kesslerdev/gitgen/pkg/generator"
)

func TestStaticInputStrategy_success(t *testing.T) {
	strategy := newStaticInputStrategy("../../test", g.BuildInputSpec{
		Strategy: "static",
		Options:  []string{".gitignore", "unknown_file"},
	})
	files, err := strategy.GetFiles()
	if err != nil {
		t.Error("StaticInputStrategy error in test directory", err)
	}

	if len(files) != 1 {
		t.Error("StaticInputStrategy only once file has to be found", err)
	}
}

func TestStaticInputStrategyUnmarshal_success(t *testing.T) {
	gen := &g.Generator{}

	err := g.FromYamlFile("../../test/valid.static.gitgen.yaml", gen)

	if err != nil {
		t.Error("StaticInputStrategy not fail on valid YAML", err)
	}

	if len(gen.Spec.Build.Input.Options.([]string)) != 1 {
		t.Error("StaticInputStrategy Should have only one element", err)
	}
}

func TestStaticInputStrategyUnmarshal_fail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	err := g.FromYamlFile("../../test/invalid.static.gitgen.yaml", &g.Generator{})

	if err == nil {
		t.Error("StaticInputStrategy should fail on invalid YAML", err)
	}
}
