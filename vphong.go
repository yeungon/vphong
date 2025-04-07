package vphong

import (
	"fmt"
	"strings"
)

// ConvertSentence converts a Vietnamese sentence to IPA with a delimiter for each word
func ConvertSentence(sentence string, glottal, palatals bool, delimit string) string {
	// Split the sentence into words
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return ""
	}

	// Convert each word to IPA
	var converted []string
	for _, word := range words {
		ipa := ConvertCustomize(word, glottal, palatals, delimit)
		converted = append(converted, ipa)
	}

	// Join the converted words with a space
	return strings.Join(converted, " ")
}

// ConvertCustomize converts a Vietnamese word to IPA with a delimiter
func ConvertCustomize(word string, glottal, palatals bool, delimit string) string {
	word = strings.ToLower(word)
	ons, nuc, cod, ton := Trans(word, glottal, palatals)
	if ons == "" && nuc == "" && cod == "" && ton == "" {
		return "[" + word + "]"
	}
	parts := []string{ons, nuc, cod, ton}
	var filtered []string
	for _, p := range parts {
		if p != "" {
			filtered = append(filtered, p)
		}
	}
	return delimit + strings.Join(filtered, delimit) + delimit
}

// Trans converts a Vietnamese word to its phonetic components based on options
func Trans(word string, glottal, palatals bool) (string, string, string, string) {
	// Use custom maps directly
	onsets := CusOnsets
	nuclei := CusNuclei
	codasMap := CusCodasMap
	onglides := CusOnglides
	offglides := CusOffglides
	onoffglides := CusOnoffglides
	//specialCases := CusSpecialRhyme
	qu := CusQu
	gi := CusGi
	tones := CusTonesP

	//fmt.Println("specialCases", specialCases)
	ons_ipa, nuc_ipa, cod_ipa, ton_ipa := "", "", "", "1" // Default tone is "1"
	oOffset, cOffset := 0, 0
	l := len(word)

	if l > 0 {
		ons_ipa, oOffset = DetectOnset(l, word, onsets)
		cod_ipa, cOffset = DetectConsotantCoda(l, word, codasMap)
		ons_ipa, nuc_ipa, cod_ipa = DetectEdgeCases(gi, word, l, ons_ipa, onsets, nuclei, qu, onglides, onoffglides, offglides, oOffset, cOffset, cod_ipa, true)
		ton_ipa = DetecTone(tones, word, l, oOffset, cOffset)

	}
	return ons_ipa, nuc_ipa, cod_ipa, ton_ipa
}

func DetectOnset(l int, word string, onsets map[string]string) (string, int) {
	var ons string
	var oOffset int
	// Onset detection
	if l >= 3 && onsets[word[0:3]] != "" {
		ons = onsets[word[0:3]]
		oOffset = 3
	} else if l >= 2 && onsets[word[0:2]] != "" {
		ons = onsets[word[0:2]]
		oOffset = 2
	} else if onsets[word[0:1]] != "" {
		ons = onsets[word[0:1]]
		oOffset = 1
	}

	//fmt.Printf("line 92: từ: %s || ons: %s || oOffset: %v\n", word, ons, oOffset)
	return ons, oOffset
}

func DetectConsotantCoda(l int, word string, codas map[string]string) (string, int) {
	var cod string
	var cOffset int
	if l >= 2 && codas[word[l-2:l]] != "" {
		cod = codas[word[l-2:l]]
		cOffset = 2
	} else if codas[word[l-1:l]] != "" {
		cod = codas[word[l-1:l]]
		cOffset = 1
	}
	// fmt.Printf("line 110: từ: %s || cod: %s\n", word, cod)
	return cod, cOffset
}

func DetecTone(tones map[string]int, word string, l int, oOffset int, cOffset int) string {
	// Tones detection
	var ton string
	if tones != nil {
		toneChar := ""
		nucl := word[oOffset : l-cOffset]
		for _, r := range nucl {
			s := string(r)
			if _, ok := tones[s]; ok {
				toneChar = s
				break
			}
		}
		if toneChar != "" {
			ton = fmt.Sprintf("%d", tones[toneChar])
			return ton
		}
	}
	return "1"
}
func DetectEdgeCases(
	gi map[string]string,
	word string,
	wordLen int,
	initialOnset string,
	onsets, nuclei, qu, onglides, onoffglides, offglides map[string]string,
	oOffset, cOffset int,
	coda string,
	glottal bool,
) (string, string, string) {
	var nucleus, onset = "", initialOnset

	// Special case: word starts with specific sequence, has a coda, and length is exactly 3
	if gi[word[:2]] != "" && coda != "" && wordLen == 3 {
		return "z", "i", coda
	}

	// Extract nucleus part
	nucleusPart := word[oOffset : wordLen-cOffset]

	switch {
	case nuclei[nucleusPart] != "":
		if oOffset == 0 && glottal && onsets[word[:1]] == "" {
			onset = "ʔ" + nuclei[nucleusPart]
		} else {
			nucleus = nuclei[nucleusPart]
		}

	case onglides[nucleusPart] != "" && onset != "kw":
		nucleus = onglides[nucleusPart]
		if onset != "" {
			onset += "w"
		} else {
			onset = "w"
		}

	case onglides[nucleusPart] != "" && onset == "kw":
		nucleus = onglides[nucleusPart]

	case onoffglides[nucleusPart] != "":
		combined := onoffglides[nucleusPart]
		nucleus = combined[:len(combined)-1]
		coda = string(combined[len(combined)-1])
		if onset != "kw" {
			if onset != "" {
				onset += "w"
			} else {
				onset = "w"
			}
		}

	case offglides[nucleusPart] != "":
		combined := offglides[nucleusPart]
		nucleus = combined[:len(combined)-1]
		coda = string(combined[len(combined)-1])

	case gi[word] != "":
		value := gi[word]
		onset = string(value[0])
		nucleus = string(value[1])

	case qu[word] != "":
		value := qu[word]
		onset = value[:len(value)-1]
		nucleus = string(value[len(value)-1])

	default:
		return "", "", "" // Non-Vietnamese or unrecognized word
	}

	runes := []rune(word)
	if len(runes) >= 3 {
		specialEnding := string(runes[len(runes)-3:])
		flag := contains(SpecialRhyme, specialEnding)
		if flag {
			//fmt.Printf("specialEnding: %s || flag: %t\n", specialEnding, flag)
			nucleus = "ɛ"
		}

	}

	//specialCases := nucleus + coda

	//fmt.Printf("edgecases: line 204: từ: %s||nuc %s || coda: %s ||flag: %t || special : %s\n", word, nucleus, coda, flag, specialCases)

	//fmt.Printf("edgecases: line 204: từ: %s||nuc %s || coda: %s\n", word, nucleus, coda)
	return onset, nucleus, coda
}

// contains checks if a string is in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
