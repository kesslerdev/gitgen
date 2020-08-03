package generator

type liteBuildOutputSpec struct {
	Strategy string
	Copy     []string
}

// OutputOptionUnmarshalSpec is used to create new OutputStrategy
type OutputOptionUnmarshalSpec func(unmarshal func(interface{}) error) interface{}

var outputUnmarshalSpecs = map[string]OutputOptionUnmarshalSpec{}

// AddOutputOptionUnmarshalSpec add new Output Strategy to the system
func AddOutputOptionUnmarshalSpec(name string, f OutputOptionUnmarshalSpec) {
	outputUnmarshalSpecs[name] = f
}

func (v *BuildOutputSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s liteBuildOutputSpec
	if err := unmarshal(&s); err != nil {
		panic(err)
	}
	v.Strategy = s.Strategy
	v.Copy = s.Copy
	v.Options = outputUnmarshalSpecs[s.Strategy](unmarshal)

	return nil
}
