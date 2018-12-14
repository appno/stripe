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
	return NewDocument(data)
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

func (d *Document) getPath() (string, error) {
	home, err := GetAppHome()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(home, os.ModePerm)
	if err != nil {
		return "", nil
	}

	return path.Join(home, fmt.Sprintf("%s.json", d.ID)), nil
}

// GetPastDueCompliance : Update and return new compliance state
func (d *Document) GetPastDueCompliance() *CompliancePastDue {
	return d.getPastDueCompliance(time.Now())
}

func (d *Document) getPastDueCompliance(timestamp time.Time) *CompliancePastDue {
	data, _ := d.ReadCompliance()
	newData, _ := d.computeCompliance(timestamp, data)
	d.SaveCompliance(newData)
	deadline := GetDeadlineDuration()

	requirements := []string{}
	pastDue := []string{}
	for key, value := range newData {
		if timestamp.Sub(*value) >= deadline {
			pastDue = append(pastDue, key)
		}
		requirements = append(requirements, key)
	}
	return MakeCompliancePastDue(requirements, pastDue)
}

// ReadCompliance : Read stored compliance data
func (d *Document) ReadCompliance() (map[string]*time.Time, error) {
	if d.ID == "" {
		return nil, nil
	}

	path, err := d.getPath()
	if err != nil {
		return nil, err
	}

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
	return d.computeCompliance(time.Now(), data)
}

// ComputeCompliance : Compute past due compliance
func (d *Document) computeCompliance(timestamp time.Time, data map[string]*time.Time) (map[string]*time.Time, error) {
	newData := make(map[string]*time.Time)

	compliance := d.GetCompliance()
	for _, key := range compliance.Requirements {
		value, ok := data[key]
		if ok {
			newData[key] = value
		} else {
			newData[key] = &timestamp
		}
	}
	return newData, nil
}

// SaveCompliance : Save compliance data to file
func (d *Document) SaveCompliance(data map[string]*time.Time) error {
	if d.ID == "" {
		return fmt.Errorf("id is not set: cannot save document")
	}

	path, err := d.getPath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCompliance : Delete compliance data file
func (d *Document) DeleteCompliance() error {
	if d.ID == "" {
		return fmt.Errorf("id is not set: cannot delete document")
	}

	path, err := d.getPath()

	if err != nil {
		return err
	}

	return os.Remove(path)
}
