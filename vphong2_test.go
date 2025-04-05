package vphong

import (
	"testing"
)

func TestVphong2(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "xin chào việt nam cải cay cao cau",
			expected: "sin1 caw2 viet6 nam1 kaj4 kăj1 kaw1 kăw1",
		},
		{
			input: "na nga ca cá cha bánh canh cá lọc",
			// bánh canh need to fixed
			expected: "na1 ŋa1 ka1 ka5 ca1 ɓaŋ5 kaŋ1 ka5 lɔk6",
		},
		{
			input:    "gió mùa màu đông ",
			expected: "zɔ5 muo2 măw2 doŋ1",
		},
		{
			//Ref: http://www.lel.ed.ac.uk/~jkirby/docs/kirby2011vietnamese.pdf
			input:    "Gió bấc và mặt trời cãi nhau xem ai mạnh",
			expected: "zɔ5 ɓɤ̆k5 va2 măt6 ʈɤj2 kaj3 ɲăw1 sɛm1 aj1 maŋ6",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()                                        // Mark this subtest as parallel
			result := ConvertSentence(tc.input, true, true, "") // Use ConvertSentence
			if result != tc.expected {
				t.Errorf("wanted: %s, got: %s", tc.expected, result)
			}
		})
	}
}
