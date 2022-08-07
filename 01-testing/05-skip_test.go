package testingdemo

import (
	"encoding/json"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func FuzzJSONMarshalling(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var v interface{}
		if err := json.Unmarshal(b, &v); err != nil {
			t.Skip()
		}
		if _, err := json.Marshal(v); err != nil {
			t.Error("Marshal: ", err)
		}
	})
}
