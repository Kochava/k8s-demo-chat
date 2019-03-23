package broadcast

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

type InputValidator interface {
	Validate(p []byte) (bool, error)
}

var (
	_ InputValidator = &JSONSchemaValidator{}
)

func NewJSONSchemaValidator(schemaPath string) (*JSONSchemaValidator, error) {
	var (
		err          error
		schemaLoader = gojsonschema.NewReferenceLoader(schemaPath)
		validator    = &JSONSchemaValidator{}
	)

	if validator.Schema, err = gojsonschema.NewSchema(schemaLoader); err != nil {
		return nil, err
	}

	return validator, nil
}

type JSONSchemaValidator struct {
	Schema *gojsonschema.Schema
}

func (validator *JSONSchemaValidator) Validate(input []byte) (bool, error) {
	var (
		err              error
		validationResult *gojsonschema.Result

		document = gojsonschema.NewBytesLoader(input)
	)

	if validationResult, err = validator.Schema.Validate(document); err != nil {
		return false, err
	}

	if len(validationResult.Errors()) > 0 {
		return false, errors.New("error validating input")
	}

	return validationResult.Valid(), nil
}
