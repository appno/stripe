package schema

import (
	"encoding/json"
	"reflect"
	"sort"

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

// JSONString : JSON Representation of Compliance object
func (c *Compliance) JSONString() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}

	return string(json[:]), nil
}

func (c *Compliance) equals(other *Compliance) bool {
	if c.Compliant != other.Compliant {
		return false
	}
	sort.Strings(c.Requirements)
	sort.Strings(other.Requirements)

	return reflect.DeepEqual(c.Requirements, other.Requirements)
}

// DebugString : Debug representation of Compliance object
func (c *Compliance) DebugString() string {
	return spew.Sdump(c)
}

// CompliancePastDue : Compliance with past_due JSON data
type CompliancePastDue struct {
	Compliant    bool     `json:"compliant"`
	Requirements []string `json:"requirements"`
	PastDue      []string `json:"past_due"`
}

// MakeCompliancePastDue : CompliancePastDue Factory
func MakeCompliancePastDue(requirements []string, pastDue []string) *CompliancePastDue {
	compliant := len(requirements) == 0
	return &CompliancePastDue{compliant, requirements, pastDue}
}

// JSONString : JSON Representation of CompliancePastDue object
func (c *CompliancePastDue) JSONString() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}

	return string(json[:]), nil
}

func (c *CompliancePastDue) equals(other *CompliancePastDue) bool {
	if c.Compliant != other.Compliant {
		return false
	}
	sort.Strings(c.Requirements)
	sort.Strings(other.Requirements)
	sort.Strings(c.PastDue)
	sort.Strings(other.PastDue)
	return reflect.DeepEqual(c, other)
}

// DebugString : Debug representation of CompliancePastDue object
func (c *CompliancePastDue) DebugString() string {
	return spew.Sdump(c)
}
