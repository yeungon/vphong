package vphong

import (
	"fmt"
	"strings"
)

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
	return ons, oOffset
}

func DetectCoda(l int, word string, codas map[string]string) (string, int) {
	var cod string
	var cOffset int
	if l >= 2 && codas[word[l-2:l]] != "" {
		cod = codas[word[l-2:l]]
		cOffset = 2
	} else if codas[word[l-1:l]] != "" {
		cod = codas[word[l-1:l]]
		cOffset = 1
	}
	return cod, cOffset
}

// Trans converts a Vietnamese word to its phonetic components based on options
func Trans(word string, glottal, pham, cao, palatals bool) (string, string, string, string) {
	// Use custom maps directly
	onsets := CusOnsets
	nuclei := CusNuclei
	codas := CusCodas
	onglides := CusOnglides
	offglides := CusOffglides
	onoffglides := CusOnoffglides
	specialCases := CusSpecialVan
	qu := CusQu
	gi := CusGi

	var tones map[string]int
	if pham || cao {
		tones = CusTonesP
	}

	fmt.Println("specialCases", specialCases)
	ons, nuc, cod, ton := "", "", "", "1" // Default tone is "1"
	oOffset, cOffset := 0, 0
	l := len(word)

	if l > 0 {
		// Onset detection
		ons, oOffset = DetectOnset(l, word, onsets)
		// Coda detection
		cod, cOffset = DetectCoda(l, word, codas)

		// Nucleus and special cases
		if gi[word[0:2]] != "" && cod != "" && l == 3 {
			nuc = "i"
			ons = "z"
		} else {
			nucl := word[oOffset : l-cOffset]

			fmt.Printf("word và nucl %s và %s\n", word, nucl)

			switch {
			case nuclei[nucl] != "":
				if oOffset == 0 {
					if glottal && onsets[word[0:1]] == "" {
						ons = "ʔ" + nuclei[nucl]
					} else {
						nuc = nuclei[nucl]
					}
				} else {
					nuc = nuclei[nucl]
				}
			case onglides[nucl] != "" && ons != "kw":
				nuc = onglides[nucl]
				if ons != "" {
					ons += "w"
				} else {
					ons = "w"
				}
			case onglides[nucl] != "" && ons == "kw":
				nuc = onglides[nucl]
			case onoffglides[nucl] != "":
				cod = string(onoffglides[nucl][len(onoffglides[nucl])-1])
				nuc = onoffglides[nucl][:len(onoffglides[nucl])-1]
				if ons != "kw" {
					if ons != "" {
						ons += "w"
					} else {
						ons = "w"
					}
				}
			case offglides[nucl] != "":
				cod = string(offglides[nucl][len(offglides[nucl])-1])
				nuc = offglides[nucl][:len(offglides[nucl])-1]
			case gi[word] != "":
				ons = string(gi[word][0])
				nuc = string(gi[word][1])
			case qu[word] != "":
				ons = qu[word][:len(qu[word])-1]
				nuc = string(qu[word][len(qu[word])-1])
			default:
				return "", "", "", "" // Non-Vietnamese word
			}
		}

		// Palatals logic
		if palatals && contains([]string{"i", "e", "ɛ"}, nuc) && cod == "k" {
			cod = "c"
		}

		// Tones detection
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
			}
		}

		// Modifications for closed syllables
		if cOffset != 0 {
			if cao {
				if ton == "5" && contains([]string{"p", "t", "k"}, cod) {
					ton = "5b"
				}
				if ton == "6" && contains([]string{"p", "t", "k"}, cod) {
					ton = "6b"
				}
			}
		}
	}

	return ons, nuc, cod, ton
}

// ConvertCustomize converts a Vietnamese word to IPA with a delimiter
func ConvertCustomize(word string, glottal, pham, cao, palatals bool, delimit string) string {
	word = strings.ToLower(word)
	ons, nuc, cod, ton := Trans(word, glottal, pham, cao, palatals)
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

// ConvertSentence converts a Vietnamese sentence to IPA with a delimiter for each word
func ConvertSentence(sentence string, glottal, pham, cao, palatals bool, delimit string) string {
	// Split the sentence into words
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return ""
	}

	// Convert each word to IPA
	var converted []string
	for _, word := range words {
		ipa := ConvertCustomize(word, glottal, pham, cao, palatals, delimit)
		converted = append(converted, ipa)
	}

	// Join the converted words with a space
	return strings.Join(converted, " ")
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
