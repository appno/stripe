package document

import (
	"encoding/json"
	"reflect"
	"sort"

	"github.com/appno/stripe/validator"
	"github.com/davecgh/go-spew/spew"
)

// Compliance : Compliance JSON data
type Compliance struct {
	Compliant    bool     `json:"compliant"`
	Requirements []string `json:"requirements"`
}

// MakeCompliance : Compliance Factory
func MakeCompliance(requirements []string) *Compliance {
	compliant := len(requirements) == 0
	return &Compliance{compliant, requirements}
}

// ComplianceFromData : Compliance Factory From Data
func ComplianceFromData(data interface{}) (*Compliance, error) {
	requirements, err := validator.DocumentValidator.Validate(data)
	if err != nil {
		return nil, err
	}

	compliant := len(requirements) == 0
	return &Compliance{compliant, requirements}, nil
}

// JSONString : JSON Representation of Compliance object
func (c *Compliance) JSONString() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}

	return string(json[:]), nil
}

func (c *Compliance) equal(other *Compliance) bool {
	if c.Compliant != other.Compliant {
		return false
	}

	cr := make([]string, len(c.Requirements))
	copy(cr, c.Requirements)
	or := make([]string, len(other.Requirements))
	copy(or, other.Requirements)

	sort.Strings(cr)
	sort.Strings(or)

	sort.Strings(c.Requirements)
	sort.Strings(other.Requirements)
	return reflect.DeepEqual(cr, or)
}

// DebugString : Debug representation of Compliance object
func (c *Compliance) DebugString() string {
	return spew.Sdump(c)
}

// CompliancePastDue : Compliance with past_due JSON data
type CompliancePastDue struct {
	Compliance
	PastDue []string `json:"past_due"`
}

// MakeCompliancePastDue : CompliancePastDue Factory
func MakeCompliancePastDue(requirements []string, pastDue []string) *CompliancePastDue {
	compliant := len(requirements) == 0
	return &CompliancePastDue{Compliance{compliant, requirements}, pastDue}
}

// JSONString : JSON Representation of CompliancePastDue object
func (c *CompliancePastDue) JSONString() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}

	return string(json[:]), nil
}

func (c *CompliancePastDue) equal(other *CompliancePastDue) bool {
	if !c.Compliance.equal(&other.Compliance) {
		return false
	}

	cp := make([]string, len(c.PastDue))
	copy(cp, c.PastDue)
	op := make([]string, len(other.PastDue))
	copy(op, other.PastDue)

	sort.Strings(cp)
	sort.Strings(op)
	return reflect.DeepEqual(cp, op)
}

// DebugString : Debug representation of CompliancePastDue object
func (c *CompliancePastDue) DebugString() string {
	return spew.Sdump(c)
}
