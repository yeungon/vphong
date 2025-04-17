package vphong

import "strings"

// contains checks if a string is in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

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
