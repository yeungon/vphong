package vphong

import (
	"testing"
)

//onsets[word[0:3]]

func TestVphong2CusOnsets(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    CusOnsets["bánh"[0:1]],
			expected: "ɓ",
		},
		{
			input:    CusOnsets["kiến"[0:1]],
			expected: "k",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			result := tc.input
			if result != tc.expected {
				t.Errorf("wanted: %s, got: %s", tc.expected, result)
			}
		})
	}
}

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
			input:    "na nga ca cá cha bánh canh cá lọc",
			expected: "na1 ŋa1 ka1 ka5 ca1 ɓɛŋ5 kɛŋ1 ka5 lɔk6",
		},
		{
			input:    "cải cay sạch sành sanh đau tay",
			expected: "kaj4 kăj1 ʂɛk6 ʂɛŋ2 ʂɛŋ1 dăw1 tăj1",
		},
		{
			//Retrieved: http://www.lel.ed.ac.uk/~jkirby/docs/kirby2011vietnamese.pdf
			input:    "Gió bấc và mặt trời cãi nhau xem ai mạnh mạng",
			expected: "zɔ5 ɓɤ̆k5 va2 măt6 ʈɤj2 kaj3 ɲăw1 sɛm1 aj1 mɛŋ6 maŋ6",
		},

		{
			//Retrieved from https://github.com/v-nhandt21/Viphoneme
			input:    "có thể xử lý những trường hợp chứa",
			expected: "kɔ5 tʰe4 sɯ4 li5 ɲɯŋ3 ʈɯəŋ2 hɤp6 cɯə5",
		},

		{
			input:    "an ang anh ảnh au ay á an ó",
			expected: "an1 aŋ1 ɛŋ1 ɛŋ4 ăw1 ăj1 a5 an1 ɔ5",
		},

		{
			input:    "anh em như thể tay chân",
			expected: "ɛŋ1 ɛm1 ɲɯ1 tʰe4 tăj1 cɤ̆n1",
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
