package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	index := 0
	tmpString := ""
	char := ""

	for {
		char = phrase[index : index+1]

		if isSeparateWord(char) {
			if 0 < len(tmpString) {
				freq[strings.ToLower(tmpString)]++
			}
			tmpString = ""
		} else {
			tmpString = tmpString + char
		}

		if index+1 == len(phrase) {
			if 0 < len(tmpString) {
				freq[strings.ToLower(tmpString)]++
			}
			break
		}

		index++
	}
	return freq
}

func isSeparateWord(word string) bool {
	valid := "abcdefghijklmnopqrstuvwxyz1234567890"
	char := strings.ToLower(word)

	position := strings.Index(valid, char)

	if -1 == position {
		return true
	}
	return false
}
