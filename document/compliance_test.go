package document

import (
	"encoding/json"
	"testing"

	"github.com/appno/stripe/internal"
)

type testData map[string]*testInstance

type testInstance struct {
	Data   interface{} `json:"data"`
	Output *Compliance `json:"output"`
}

func TestCompliance(t *testing.T) {
	bytes, err := internal.LoadTestData("testdata/part_1.json")
	if err != nil {
		t.Errorf("FILE READ ERROR: err ==%+v != nil", err)
	}

	var result testData
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("UNMARSHAL ERROR: err == %+v != nil", err)
	}

	for key, value := range result {
		actual, err := ComplianceFromData(value.Data)
		if err != nil {
			t.Errorf("%s: %+v != nil", key, err)
		} else {
			if !actual.equivalent(value.Output) {
				t.Errorf("%s: %+v != %+v", key, actual, value.Output.Requirements)
			}
		}
	}
}
