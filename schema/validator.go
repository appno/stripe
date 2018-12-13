package schema

import (
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// Validator : Validate JSON data
type Validator struct {
	Schema *gojsonschema.JSONLoader
}

// MakeValidator : Validator factory
func MakeValidator(schema string) *Validator {
	loader := gojsonschema.NewStringLoader(schema)
	return &Validator{Schema: &loader}
}

// StoreData : Unmarshal bytes into a struct
func (v *Validator) StoreData(bytes []byte) (*Compliance, error) {
	var result Document
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	if len(result.ID) == 0 {
		return ComplianceFail(), nil
	}

	return v.IsCompliant(result)
}

// DefaultValidator : Create Validator with default Schema
func DefaultValidator() *Validator {
	return MakeValidator(Schema)
}

// IsCompliant : Check data for compliance
func (v *Validator) IsCompliant(data interface{}) (*Compliance, error) {
	requirements, err := v.Validate(data)
	if err != nil {
		return nil, err
	}

	compliant := true
	if len(requirements) > 0 {
		compliant = false
	}

	return &Compliance{compliant, requirements}, nil
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
