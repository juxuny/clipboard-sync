package clipboard

import "testing"

func TestGetLocal(t *testing.T) {
	data, err := GetLocal()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
