package generator

// import (
// 	"strings"

// 	"github.com/santhosh-tekuri/jsonschema"
// )

// var schemaText = `
// {
// 	"type": "object",
// 	"items": {
// 		"type": "string"
// 	}
// }
// `

// // ValidateBuildSpec return an error if the generator cannot be builded with this spec
// func ValidateBuildSpec(b *BuildSpec) error {

// 	compiler := jsonschema.NewCompiler()
// 	if err := compiler.AddResource("schema.json", strings.NewReader(schemaText)); err != nil {
// 		return err
// 	}
// 	schema, err := compiler.Compile("schema.json")
// 	if err != nil {
// 		return err
// 	}
// 	if err := schema.ValidateInterface(b); err != nil {
// 		return err
// 	}

// 	return nil
// }
