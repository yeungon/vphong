package vphong

import "testing"

func TestRunWithInput(t *testing.T) {
	input := "xin chào việt nam"
	result := RunWithInput(input, "n", false, false, false, false, false, "", "", false)
	wanted := "siːnᴬ¹ tɕaːwᴬ¹ viətᴰ¹ naːmᴬ¹"
	if result != wanted {
		t.Errorf("wanted: %s, got: %s", wanted, result)
	} else {
		t.Logf("wanted: %s, got: %s", wanted, result)
	}
}
