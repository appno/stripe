package schema

import (
	"encoding/json"
	"reflect"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

// Compliance : Compliance JSON data
type Compliance struct {
	Compliance    bool     `json:"compliant"`
	Requirements []string `json:"requirements"`
}

func ComplianceFail() *Compliance {
	requirements := []string{
		"id",
		"tax_id",
		"first_name",
		"last_name",
		"address.street",
		"address.city",
		"address.postal_code",
		"address.country",
	}
	return &Compliance{false, requirements}
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
	if c.Compliance != other.Compliance {
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
	Compliance    bool     `json:"compliant"`
	Requirements []string `json:"requirements"`
	PastDue      []string `json:"past_due"`
}

// CompliancePastDueFail : CompliancePastDue with all fields failing validation
func CompliancePastDueFail() *CompliancePastDue {
	c := ComplianceFail()
	return &CompliancePastDue{
		c.Compliance,
		c.Requirements,
		[]string{},
	}
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
	if c.Compliance != other.Compliance {
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
