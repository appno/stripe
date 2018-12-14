package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/appno/stripe/internal"
)

type documentTestData []*documentTestInstance

type documentTestInstance struct {
	Time   int                `json:"time"`
	Data   interface{}        `json:"data"`
	Output *CompliancePastDue `json:"output"`
}

func (t *documentTestInstance) GetDuration() time.Duration {
	return time.Duration(1000000000 * t.Time)
}

func loadData() (documentTestData, error) {
	bytes, err := internal.LoadTestData("testdata/part_2.json")
	if err != nil {
		return nil, err
	}

	var data documentTestData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestDocumentComputeCompliance(t *testing.T) {
	appHome, err := GetAppHome()
	if err != nil {
		t.Errorf("Error getting app home directory: %s", err)
	}
	cases, err := loadData()
	if err != nil {
		t.Errorf("Error loading test data: %s", err)
	}

	tempDir, err := ioutil.TempDir("", "doctest")
	if err != nil {
		t.Errorf("Error creating temporary test directory: %+v", err)
	}

	os.Setenv("STRIPE_HOME", tempDir)

	index := 0
	testCase := cases[index]
	document, docErr := NewDocument(testCase.Data)
	if docErr != nil {
		t.Errorf("%d: Error creating document: %+v", index, docErr)
	}

	data := make(map[string]*time.Time)
	timestamp := time.Now()
	result, resErr := document.computeCompliance(timestamp, data)
	if resErr != nil {
		t.Errorf("%d: Error creating document: %+v", index, resErr)
	}

	saveErr := document.SaveCompliance(result)
	if saveErr != nil {
		t.Errorf("%d: Error saving document: %+v", index, saveErr)
	}

	fromFile, readErr := document.ReadCompliance()
	if readErr != nil {
		t.Errorf("%d: Error reading document: %+v", index, readErr)
	}

	resultKeys := []string{}
	for key := range result {
		resultKeys = append(resultKeys, key)
	}

	fromFileKeys := []string{}
	for key := range fromFile {
		fromFileKeys = append(fromFileKeys, key)
	}

	sort.Strings(resultKeys)
	sort.Strings(fromFileKeys)

	if !reflect.DeepEqual(resultKeys, fromFileKeys) {
		t.Errorf("Keys are not equal: %+v, %+v", resultKeys, fromFileKeys)
	}

	for key, value := range result {
		otherValue := fromFile[key]
		if !value.Equal(*otherValue) {
			fmt.Printf("\n%s: %+v : %+v\n", key, value, otherValue)
		}
	}

	os.Setenv("STRIPE_HOME", appHome)
}

func TestDocument(t *testing.T) {
	appHome, err := GetAppHome()
	if err != nil {
		t.Errorf("Error getting app home directory: %s", err)
	}
	cases, err := loadData()
	if err != nil {
		t.Errorf("Error loading test data: %s", err)
	}
	now := time.Now()

	for i := 0; i < len(cases); i++ {
		duration := cases[i].GetDuration()

		tempDir, err := ioutil.TempDir("", "doctest")
		if err != nil {
			t.Errorf("Error creating temporary test directory: %+v", err)
		}

		os.Setenv("STRIPE_HOME", tempDir)

		for j, testCase := range cases[:i+1] {

			document, docErr := NewDocument(testCase.Data)
			if err != nil {
				t.Errorf("%d: Error creating document: %+v", j, docErr)
			}

			currDuration := testCase.GetDuration()
			timestamp := now.Add(currDuration - duration)
			compliance := document.getPastDueCompliance(timestamp)

			if !compliance.equals(testCase.Output) {
				t.Errorf("%d:%d: %+v != %+v", i, j, compliance, testCase.Output)
			}
		}

		err = os.RemoveAll(tempDir)
		if err != nil {
			t.Logf("Error removing temporary directory: %+v", err)
		}
	}

	os.Setenv("STRIPE_HOME", appHome)
}
