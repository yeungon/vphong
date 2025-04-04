package vphong

import "testing"

func TestRunWithInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "xin chào việt nam",
			expected: "siːnᴬ¹ tɕaːwᴬ¹ viətᴰ¹ naːmᴬ¹",
		},
		{
			input:    "đầu lòng hai ả Tố Nga",
			expected: "ɗəwᴮ¹ lɔŋ͡mᴬ¹ haːjᴬ¹ ʔaːᴮ¹ toːᴮ¹ ŋaːᴬ¹",
		},
		{
			input:    "na nga ca cá cha",
			expected: "naːᴬ¹ ŋaːᴬ¹ kaːᴬ¹ kaːᴬ¹ tɕaːᴬ¹",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel() // Mark this subtest as parallel
			result := RunWithInput(tc.input, "n", false, false, false, false, false, "", "", false)
			if result != tc.expected {
				t.Errorf("wanted: %s, got: %s", tc.expected, result)
			}
		})
	}
}
