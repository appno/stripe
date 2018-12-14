package validator

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// DocumentValidator : global document validator
var DocumentValidator = DefaultValidator()

// Validator : Validate JSON data
type Validator struct {
	Schema *gojsonschema.JSONLoader
}

// MakeValidator : Validator factory
func MakeValidator(schema string) *Validator {
	loader := gojsonschema.NewStringLoader(schema)
	return &Validator{Schema: &loader}
}

// DefaultValidator : Create Validator with default Schema
func DefaultValidator() *Validator {
	return MakeValidator(Schema)
}

// Validate : validate json string
func (v *Validator) Validate(data interface{}) ([]string, error) {
	loader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(*v.Schema, loader)
	if err != nil {
		return nil, err
	}
	requirements := []string{}

	addressMissing := false
	addressFieldMissing := false
	for _, error := range result.Errors() {
		field := error.Field()

		context := error.Context().String()

		if len(context) > 7 {
			prefix := context[7:len(context)]
			if prefix != field {
				addressFieldMissing = true
				field = fmt.Sprintf("%s.%s", prefix, field)
			}
		}

		if field == "address" {
			addressMissing = true
		} else {
			requirements = append(requirements, field)
		}
	}

	if addressMissing && !addressFieldMissing {
		fields := []string{"address.city", "address.country", "address.postal_code", "address.street"}
		requirements = append(requirements, fields...)
	}
	return requirements, nil
}
