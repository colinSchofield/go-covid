package daily

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_Model_UnMarshalling(t *testing.T) {

	jsonAsString, err := os.ReadFile("./test_model.json")
	if err != nil {
		t.Errorf("Failed to read test input file: %s", err)
	}

	var daily Daily
	if err := json.Unmarshal([]byte(jsonAsString), &daily); err != nil {
		t.Errorf("Did not DeMarshal structure correctly, error was: %s", err)
	}
	if len(daily.Response) != 2 {
		t.Errorf("Countries read from json file is not equal to 2. Actual number was: %d", len(daily.Response))
	}
}
