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
	input2 := "đầu lòng hai ả Tố Nga"
	result2 := RunWithInput(input2, "n", false, false, false, false, false, "", "", false)
	wanted2 := "ɗəwᴮ¹ lɔŋ͡mᴬ¹ haːjᴬ¹ ʔaːᴮ¹ toːᴮ¹ ŋaːᴬ¹"
	if result2 != wanted2 {
		t.Errorf("wanted: %s, got: %s", wanted, result)
	} else {
		t.Logf("wanted: %s, got: %s", wanted, result)
	}

}
