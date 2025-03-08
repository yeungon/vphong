package vphong

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// trans function
func trans(word string, dialect string, chao bool, eight bool, nosuper bool, glottal bool, phonemic bool) (string, string, string, string, string) {
	var eightTones map[string]string
	if nosuper {
		eightTones = eightLower
	} else {
		eightTones = eightSuper
	}

	var chaoTones map[string]string
	switch dialect {
	case "s":
		if nosuper {
			chaoTones = ChaoSLower
		} else {
			chaoTones = ChaoSSuper
		}
	case "c":
		if nosuper {
			chaoTones = ChaoCLower
		} else {
			chaoTones = ChaoCSuper
		}
	default:
		if nosuper {
			chaoTones = ChaoNLower
		} else {
			chaoTones = ChaoNSuper
		}
	}

	ldNas := "ŋ͡m" // if nosuper else "ŋᵐ"
	ldPlo := "k͡p" // if nosuper else "kᵖ"
	palNas := "ɲ"  // if nosuper else "ʲŋ"
	palPlo := "c"  // if nosuper else "ʲk"
	lvGli := "w"   // if nosuper else "ʷ"

	if nosuper {
		onsets["th"] = "th"
		onsets["qu"] = "kw"
	}

	var ons, gli, nuc, cod, ton string
	oOffset, cOffset := 0, 0
	length := len(word)

	if length > 0 {
		if val, ok := onsets[word[0:3]]; ok && length >= 3 {
			ons = val
			oOffset = 3
		} else if val, ok := onsets[word[0:2]]; ok && length >= 2 {
			ons = val
			oOffset = 2
		} else if val, ok := onsets[word[0:1]]; ok {
			ons = val
			oOffset = 1
		}

		if val, ok := codas[word[length-2:length]]; ok && length >= 2 {
			cod = val
			cOffset = 2
		} else if val, ok := codas[word[length-1:length]]; ok {
			cod = val
			cOffset = 1
		}

		nucl := word[oOffset : length-cOffset]

		if oOffset == 0 {
			ons = "ʔ"
		}

		if val, ok := qu[word]; ok {
			ons = string(val[0])
			nuc = string(val[len(val)-1])
			if len(val) > 2 {
				gli = lvGli
			}
		}

		if _, ok := gi[word[0:2]]; ok && length >= 2 {
			switch word {
			case "giền":
				nucl = "â"
			default:
				if length == 2 || (length == 3 && strings.ContainsAny(word[2:3], "nm")) {
					nucl = "i"
				} else if _, ok := nuclei[nucl]; ok && strings.ContainsAny(word[2:3], "êếềểễệ") {
					nucl = "iê"
				}
			}
			ons = onsets["gi"]
		}

		if val, ok := nuclei[nucl]; ok {
			nuc = val
		} else if val, ok := onglides[nucl]; ok {
			nuc = val
			if ons != "ʔ" {
				ons += lvGli
			} else {
				ons = "w"
			}
		} else if val, ok := onoffglides[nucl]; ok {
			if ons != "ʔ" {
				ons += lvGli
			} else {
				ons = "w"
			}
			nuc = val[0 : len(val)-1]
			cod = string(val[len(val)-1])
		} else if val, ok := offglides[nucl]; ok {
			cod = string(val[len(val)-1])
			nuc = val[0 : len(val)-1]
		} else {
			return "", "", "", "", ""
		}

		if len(ons) >= 2 && string(ons[len(ons)-1]) == lvGli {
			gli = lvGli
			ons = ons[0 : len(ons)-1]
		}

		var toneList []string
		for i := 0; i < length; i++ {
			if val, ok := tones[string(word[i])]; ok {
				toneList = append(toneList, val)
			}
		}
		if len(toneList) > 0 {
			ton = toneList[len(toneList)-1]
		} else {
			ton = "A1"
		}
		if ton == "B1" && strings.ContainsAny(cod, "ptck") {
			ton = "D1"
		}
		if ton == "B2" && strings.ContainsAny(cod, "ptck") {
			ton = "D2"
		}

		if eight {
			ton = eightTones[ton]
		} else if chao {
			ton = chaoTones[ton]
		} else if !nosuper {
			ton = gedneySuper[ton]
		}

		if glottal && ons == "ʔ" {
			ons = ""
		}

		if nuc == "aː" {
			if cod == "c" || cod == "ɲ" {
				nuc = "ɛ"
			}
		}

		if dialect != "o" {
			if strings.ContainsAny(cod, "ŋk") {
				if nuc == "ɛ" {
					nuc = "ɛː"
				}
				if nuc == "e" {
					nuc = "eː"
				}
			}
		} else {
			if word[0:2] == "gi" {
				ons = "ʑ"
			}
			if ons == "j" {
				ons = "z"
			}
		}

		if dialect == "n" || dialect == "o" {
			if cod == "c" {
				cod = "k"
			}
			if cod == "ɲ" {
				cod = "ŋ"
			}
		}

		if dialect == "n" {
			switch ons {
			case "j", "r":
				ons = "z"
			case "c", "ʈ":
				ons = "tɕ"
			case "ʂ":
				ons = "s"
			}

			if !phonemic {
				if strings.ContainsAny(cod, "kŋ") {
					if strings.ContainsAny(nuc, "eɛi") {
						if cod == "k" {
							cod = palPlo
						}
						if cod == "ŋ" {
							cod = palNas
						}
					} else if strings.ContainsAny(nuc, "uɔo") && word != "quốc" {
						if cod == "k" {
							cod = ldPlo
						}
						if cod == "ŋ" {
							cod = ldNas
						}
					}
				}

				if cod == palNas || cod == palPlo {
					if nuc == "ɛ" {
						nuc = "a"
					}
				}

				if len(nuc) == 1 && !strings.ContainsAny(nuc, "aə") {
					if (len(cod) == 1 && nuc != "ɨ") || len(cod) == 0 {
						nuc += "ː"
					}
				}
			} else {
				if len(cod) == 0 && strings.ContainsAny(nuc, "aːəː") {
					if nuc == "aː" {
						nuc = "a"
					}
					if nuc == "əː" {
						nuc = "ə"
					}
				}
			}
		} else if dialect == "s" || dialect == "c" {
			if ons == "z" {
				ons = "j"
			}
			if ons == "k" && gli == lvGli {
				ons = "w"
				gli = ""
			}
			if ons == "ɣ" {
				ons = "ɡ"
			}

			if len(cod) > 0 && strings.ContainsAny(nuc, "iəuəɨə") {
				switch nuc {
				case "iə":
					nuc = "iː"
				case "ɨə":
					nuc = "ɨː"
				case "uə":
					nuc = "uː"
				}
			}

			if nuc == "ɔ" && strings.ContainsAny(cod, "nt") {
				nuc = "ɔː"
			}
			if nuc == "o" && strings.ContainsAny(cod, "ŋk") {
				nuc = "ɔ"
			}

			if nuc == "ɛ" && strings.ContainsAny(cod, "nt") {
				if cod == "n" {
					cod = "ŋ"
				}
				nuc = "ɛː"
			}

			if len(cod) > 0 && len(nuc) == 2 {
				if cod == "n" {
					cod = "ŋ"
				}
				if cod == "t" {
					cod = "k"
				}
			}

			if len(cod) > 0 && strings.ContainsAny(nuc, "ɨəauo") {
				if cod == "n" {
					cod = "ŋ"
				}
				if cod == "t" {
					cod = "k"
				}
			}

			if len(cod) > 0 && strings.ContainsAny(nuc, "ieɛ") {
				if cod == "ŋ" {
					cod = "n"
				}
				if cod == "k" {
					cod = "t"
				}
			}

			if strings.ContainsAny(cod, "ɲc") {
				if cod == "ɲ" {
					cod = "n"
				}
				if cod == "c" {
					cod = "t"
				}
			}

			if !phonemic {
				if ons == "ʂ" {
					ons = "s"
				}

				if strings.ContainsAny(cod, "nt") {
					switch nuc {
					case "i":
						nuc = "ɨ"
					case "ɛ":
						nuc = "a"
					case "e":
						nuc = "əː"
					}
				}

				if nuc == "u" && strings.ContainsAny(cod, "mp") {
					nuc = "ɨ"
				}

				if strings.ContainsAny(nuc, "eɛoɔ") {
					switch nuc {
					case "e":
						nuc = "eː"
					case "ɛ":
						nuc = "ɛː"
					case "o":
						if !strings.ContainsAny(cod, "ŋk") {
							nuc = "oː"
						}
					case "ɔ":
						if !strings.ContainsAny(cod, "ŋk") {
							nuc = "ɔː"
						}
					}
				}

				if strings.ContainsAny(nuc, "uɔoː") && strings.ContainsAny(cod, "ŋk") {
					if cod == "ŋ" {
						cod = ldNas
					}
					if cod == "k" {
						cod = ldPlo
					}
				}
			}
		}

		if dialect != "o" && phonemic && len(cod) == 0 && strings.ContainsAny(nuc, "aːəː") {
			if nuc == "aː" {
				nuc = "a"
			}
			if nuc == "əː" {
				nuc = "ə"
			}
		}

		return ons, gli, nuc, cod, ton
	}

	return "", "", "", "", ""
}

// convert function
func convert(word string, dialect string, chao bool, eight bool, nosuper bool, glottal bool, phonemic bool, delimit string) string {
	ons, gli, nuc, cod, ton := trans(word, dialect, chao, eight, nosuper, glottal, phonemic)
	if ons == "" && gli == "" && nuc == "" && cod == "" && ton == "" {
		return "[" + word + "]"
	}
	parts := []string{ons, gli, nuc, cod, ton}
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return delimit + strings.Join(filtered, delimit) + delimit
}

// main function
func main() {
	dialect := flag.String("d", "n", "Specify dialect region (Northern=n, Central=c, Southern=s) or spelling pronunciation (o)")
	chao := flag.Bool("c", false, "Phonetize tones as Chao values")
	glottal := flag.Bool("g", false, "No glottal stops in underlying forms")
	eight := flag.Bool("8", false, "Encode tones as 1-8")
	nosuper := flag.Bool("n", false, "No superscripts anywhere")
	phonemic := flag.Bool("p", false, "Underlying transcriptions after Pham (2006)")
	delimit := flag.String("m", "", "Produce delimited output")
	outputOrtho := flag.String("o", "", "Output orthography as well as IPA")
	tokenize := flag.Bool("t", false, "Preserve underscores or hyphens in tokenized inputs")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		compound := ""
		ortho := ""
		words := strings.Fields(line)
		var filteredWords []string
		for _, word := range words {
			if len(word) > 0 && word != "-" && word != "_" {
				filteredWords = append(filteredWords, word)
			}
		}

		for i, word := range filteredWords {
			ortho += word
			cleanWord := strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			})
			cleanWord = strings.ToLower(cleanWord)

			var seq string
			if *tokenize && (strings.Contains(word, "-") || strings.Contains(word, "_")) {
				parts := strings.FieldsFunc(cleanWord, func(r rune) bool {
					return r == '-' || r == '_'
				})
				delimiters := strings.FieldsFunc(cleanWord, func(r rune) bool {
					return !(r == '-' || r == '_')
				})
				if len(delimiters) < len(parts) {
					delimiters = append(delimiters, "")
				}
				var ipa []string
				for _, part := range parts {
					ipa = append(ipa, strings.TrimSpace(convert(part, *dialect, *chao, *eight, *nosuper, *glottal, *phonemic, *delimit)))
				}
				seq = ""
				for j := range ipa {
					seq += ipa[j] + delimiters[j]
				}
			} else {
				seq = strings.TrimSpace(convert(cleanWord, *dialect, *chao, *eight, *nosuper, *glottal, *phonemic, *delimit))
			}

			if len(filteredWords) >= 2 {
				ortho += " "
			}
			if i < len(filteredWords)-1 {
				seq += " "
			}
			compound += seq
		}

		if ortho != "" {
			ortho = strings.TrimSpace(ortho)
			if *outputOrtho != "" {
				fmt.Print(ortho, *outputOrtho)
			}
			fmt.Println(compound)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
