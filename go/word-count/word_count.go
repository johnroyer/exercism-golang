package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	index := 0
	tmpString := ""
	char := ""

	for {
		char = phrase[index : index+1]

		if isSaperateWord(char) {
			freq[tmpString]++
			tmpString = ""
		} else {
			tmpString = tmpString + char
		}

		if index+1 == len(phrase) {
			freq[tmpString]++
			break
		}

		index++
	}
	return freq
}

func isSaperateWord(word string) bool {
	valid := "abcdefghijklmnopqrstuvxyz"
	char := strings.ToLower(word)

	position := strings.Index(valid, char)

	if -1 == position {
		return true
	}
	return false
}
