package vphong

import (
	"fmt"
	"unicode/utf8"
)

// Trans converts a Vietnamese word to its phonetic components based on options
func Trans(word string, glottal, palatals bool) (string, string, string, string) {
	// Use custom maps directly
	onsets := CusOnsets
	nuclei := CusNuclei
	codasMap := CusCodasMapConsonant
	onglides := CusOnglides
	offglides := CusOffglides
	onoffglides := CusOnoffglides
	qu := CusQu
	gi := CusGi
	tones := CusTonesP
	ons_ipa, nuc_ipa, cod_ipa, ton_ipa := "", "", "", "1" // Default tone is "1"
	oOffset, cOffset := 0, 0
	runeLength := utf8.RuneCountInString(word)
	if runeLength > 0 {
		ons_ipa, oOffset = DetectOnset(runeLength, word, onsets)
		cod_ipa, cOffset = DetectCoda(runeLength, word, codasMap)
		ons_ipa, nuc_ipa, cod_ipa = DetectNucleusEdgeCases(gi, word, runeLength, ons_ipa, onsets, nuclei, qu, onglides, onoffglides, offglides, oOffset, cOffset, cod_ipa, true)
		ton_ipa = DetecTone(tones, word, runeLength, oOffset, cOffset)

	}
	//fmt.Println("ons_ipa, nuc_ipa, cod_ipa", ons_ipa, nuc_ipa, cod_ipa)
	return ons_ipa, nuc_ipa, cod_ipa, ton_ipa
}

func DetectOnset(l int, word string, onsets map[string]string) (string, int) {
	var ons string
	var oOffset int
	// Convert text to a slice of rune so that we can manage it.
	runes := []rune(word)
	if l >= 3 && onsets[string(runes[0:3])] != "" {
		ons = onsets[string(runes[0:3])]
		oOffset = 3
	} else if l >= 2 && onsets[string(runes[0:2])] != "" {
		ons = onsets[string(runes[0:2])]
		oOffset = 2
	} else if onsets[string(runes[0:1])] != "" {
		ons = onsets[string(runes[0:1])]
		oOffset = 1
	}
	return ons, oOffset
}

func DetectCoda(l int, word string, codas map[string]string) (string, int) {
	var cod string
	var cOffset int
	runes := []rune(word)
	length := len(runes)

	if l >= 2 {
		twoLetter := string(runes[length-2 : length])
		oneLetter := string(runes[length-1 : length])
		if value, ok := codas[twoLetter]; ok {
			cod = value
			cOffset = 2
		} else if value, ok := codas[oneLetter]; ok {
			cod = value
			cOffset = 1
		}
	}

	//fmt.Printf("line 71: length: %v, từ: %s || cod: %s || twoLetter %v, oneLetter: %v\n", l, word, cod, twoLetter, oneLetter)
	return cod, cOffset
}

func DetecTone(tones map[string]int, word string, l int, oOffset int, cOffset int) string {
	// Tones detection
	runes := []rune(word)
	var ton string
	if tones != nil {
		toneChar := ""
		nucl := string(runes[oOffset : l-cOffset])
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
func DetectNucleusEdgeCases(
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
	runes := []rune(word)

	//fmt.Println("coda, runes[:2], wordLen", coda, string(runes[:2]), wordLen)
	if gi[string(runes[:2])] != "" && coda == "" && wordLen == 2 {
		return "z", "i", coda
	}

	if word == "giếng" || word == "giềng" {
		return "z", "ie", coda
	}

	nucleusPart := string(runes[oOffset : wordLen-cOffset])
	switch {
	case nuclei[nucleusPart] != "":
		nucleus = nuclei[nucleusPart]
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
		return onset, "", "" // Non-Vietnamese or unrecognized word
	}

	if len(runes) >= 3 {
		specialEnding := string(runes[len(runes)-3:])
		flag := contains(SpecialRhyme, specialEnding)
		if flag {
			nucleus = "ɛ"
		}

	}
	return onset, nucleus, coda
}
