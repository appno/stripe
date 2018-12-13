package schema

import (
	"encoding/json"
	"testing"

	"github.com/appno/stripe/internal"
)

type documentTestData []*documentTestInstance

type documentTestInstance struct {
	Time   int                `json:"time"`
	Data   interface{}        `json:"data"`
	Output *CompliancePastDue `json:"output"`
}

func TestDocument(t *testing.T) {
	bytes, err := internal.LoadTestData("testdata/part_4.json")
	if err != nil {
		t.Errorf("FILE READ ERROR: err ==%+v != nil", err)
	}

	var result documentTestData
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("UNMARSHAL ERROR: err == %+v != nil", err)
	}

	for i, testCase := range result {
		document, err := NewDocument(testCase.Data)
		if err != nil {
			t.Errorf("%d: Error creating document: %+v", i, err)
		}

		compliance := document.GetPastDueCompliance()
		if !compliance.equals(testCase.Output) {
			t.Errorf("%d: %+v != %+v", i, compliance, testCase.Output)
		}
	}
}
