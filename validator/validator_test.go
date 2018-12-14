package validator

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/appno/stripe/internal"
)

type testData map[string]*testInstance

type output struct {
	Requirements []string `json:"requirements"`
}

type testInstance struct {
	Data   interface{} `json:"data"`
	Output output      `json:"output"`
}

func TestValidator(t *testing.T) {
	validator := DocumentValidator

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
		actual, err := validator.Validate(value.Data)
		if err != nil {
			t.Errorf("%s: %+v != nil", key, err)
		} else {
			sort.Strings(actual)
			sort.Strings(value.Output.Requirements)
			if !reflect.DeepEqual(actual, value.Output.Requirements) {
				t.Errorf("%s: %+v != %+v", key, actual, value.Output.Requirements)
			}
		}
	}
}
