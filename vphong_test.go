package vphong

import "testing"

func TestSayHello(t *testing.T) {
	result := SayHello("Vượng")
	t.Log(result)
	wanted := "Hello, Vượng!"

	if result != wanted {
		t.Errorf("wanted: %s, got: %s", wanted, result)
	} else {
		t.Logf("wanted: %s, got: %s", wanted, result)
	}
}
