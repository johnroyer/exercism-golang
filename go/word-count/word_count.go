package wordcount

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
	switch word {
	case ":":
	case "!":
	case "?":
	case "\t":
	case "\n":
	case " ":
		return true
	}
	return false
}
