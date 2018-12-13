package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"
)

// DocumentValidator : global document validator
var DocumentValidator = DefaultValidator()

// Document struct
type Document struct {
	ID         string
	data       interface{}
	compliance *Compliance
}

// DocumentFromBytes : Create Document from byte array
func DocumentFromBytes(bytes []byte) (*Document, error) {
	var data interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	id := ""
	strKeyMap, ok := data.(map[string]interface{})
	if ok {
		value, ok := strKeyMap["id"]
		if ok {
			switch v := value.(type) {
			case int:
				id = strconv.Itoa(v)
			case string:
				id = v
			case bool:
				id = strconv.FormatBool(v)
			default:
			}
		}
	}

	// fmt.Printf("id: %T=%+v", id, id)
	return &Document{id, data, nil}, nil
}

// NewDocument : Create Document from interface
func NewDocument(data interface{}) (*Document, error) {
	id := ""
	strKeyMap, ok := data.(map[string]interface{})
	if ok {
		value, ok := strKeyMap["id"]
		if ok {
			switch v := value.(type) {
			case int:
				id = strconv.Itoa(v)
			case string:
				id = v
			case bool:
				id = strconv.FormatBool(v)
			default:
			}
		}
	}

	// fmt.Printf("id: %T=%+v", id, id)
	return &Document{id, data, nil}, nil
}

// GetCompliance : Get document's compliance object
func (d *Document) GetCompliance() *Compliance {
	if d.compliance == nil {
		compliance, _ := DocumentValidator.IsCompliant(d.data)
		d.compliance = compliance
	}
	return d.compliance
}

// GetPastDueCompliance : Update and return new compliance state
func (d *Document) GetPastDueCompliance() *CompliancePastDue {
	data, _ := d.ReadCompliance()
	newData, _ := d.ComputeCompliance(data)
	d.SaveCompliance(newData)
	now := time.Now()
	deadline := GetDeadline()

	requirements := []string{}
	pastDue := []string{}
	for key, value := range newData {
		if now.Sub(*value) > deadline {
			pastDue = append(pastDue, key)
		}
		requirements = append(requirements, key)
	}
	compliant := len(requirements) == 0
	return &CompliancePastDue{
		compliant,
		requirements,
		pastDue,
	}
}

// ReadCompliance : Read stored compliance data
func (d *Document) ReadCompliance() (map[string]*time.Time, error) {
	if d.ID == "" {
		return make(map[string]*time.Time), nil
		// return nil, fmt.Errorf("id is not set: cannot read document")
	}
	home, err := GetAppHome()
	if err != nil {
		return nil, err
	}

	path := path.Join(home, fmt.Sprintf("%s.json", d.ID))

	if _, err = os.Stat(path); os.IsNotExist(err) {
		return make(map[string]*time.Time), nil
	}
	bytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	var data map[string]*time.Time
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ComputeCompliance : Compute past due compliance
func (d *Document) ComputeCompliance(data map[string]*time.Time) (map[string]*time.Time, error) {
	now := time.Now()
	newData := make(map[string]*time.Time)

	compliance := d.GetCompliance()
	for _, val := range compliance.Requirements {
		timestamp, ok := data[val]
		if ok {
			newData[val] = timestamp
		} else {
			newData[val] = &now
		}
	}
	return newData, nil
}

// SaveCompliance : Save compliance data to file
func (d *Document) SaveCompliance(data map[string]*time.Time) error {
	if d.ID == "" {
		return fmt.Errorf("id is not set: cannot save document")
	}
	home, err := GetAppHome()
	if err != nil {
		return err
	}

	path := path.Join(home, fmt.Sprintf("%s.json", d.ID))

	byteArr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, byteArr, 0644)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCompliance : Delete compliance data file
func (d *Document) DeleteCompliance(data map[string]*time.Time) error {
	if d.ID == "" {
		return fmt.Errorf("id is not set: cannot delete document")
	}
	home, err := GetAppHome()
	if err != nil {
		return err
	}

	path := path.Join(home, fmt.Sprintf("%s.json", d.ID))

	return os.Remove(path)
}
