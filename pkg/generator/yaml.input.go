package generator

type liteBuildInputSpec struct {
	Strategy string
}

// InputOptionUnmarshalSpec is used to create new OutputStrategy
type InputOptionUnmarshalSpec func(unmarshal func(interface{}) error) interface{}

var inputUnmarshalSpecs = map[string]InputOptionUnmarshalSpec{}

// AddInputOptionUnmarshalSpec add new Output Strategy to the system
func AddInputOptionUnmarshalSpec(name string, f InputOptionUnmarshalSpec) {
	inputUnmarshalSpecs[name] = f
}

func (v *BuildInputSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s liteBuildInputSpec
	if err := unmarshal(&s); err != nil {
		panic(err)
	}
	v.Strategy = s.Strategy
	unmarshalStrategy := inputUnmarshalSpecs[s.Strategy]
	if unmarshalStrategy != nil {
		v.Options = unmarshalStrategy(unmarshal)
	}

	return nil
}
