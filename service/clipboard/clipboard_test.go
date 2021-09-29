package clipboard

import "testing"

func TestGetLocal(t *testing.T) {
	data, err := GetLocal()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestSetLocal(t *testing.T) {
	data := []byte("123")
	if err := SetLocal(data); err != nil {
		t.Fatal(err)
	}
}
