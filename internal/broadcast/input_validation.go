package broadcast

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

// InputValidator determines if data being written is acceptable
type InputValidator interface {
	Valid([]byte) (bool, error)
}

var (
	_ InputValidator = &JSONSchemaValidator{}
)

// NewJSONSchemaValidator constructs a JSONSchemaValidator
func NewJSONSchemaValidator(schemaPath string) (*JSONSchemaValidator, error) {
	var (
		err       error
		validator = &JSONSchemaValidator{}
	)

	if validator.Schema, err = gojsonschema.NewSchema(gojsonschema.NewReferenceLoader(schemaPath)); err != nil {
		return nil, err
	}

	return validator, nil
}

// JSONSchemaValidator allows input validation
type JSONSchemaValidator struct {
	Schema *gojsonschema.Schema
}

// Valid determines if the input implements the specified schema
func (validator *JSONSchemaValidator) Valid(input []byte) (bool, error) {
	var (
		err             error
		results         *gojsonschema.Result
		inputJSONLoader = gojsonschema.NewBytesLoader(input)
	)

	if results, err = validator.Schema.Validate(inputJSONLoader); err != nil {
		return false, err
	}

	if len(results.Errors()) > 0 {
		return false, errors.New("invalid input")
	}

	return results.Valid(), nil
}
