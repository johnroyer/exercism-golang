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

			value, ok := freq[tmpString]
			if ok {
				freq[tmpString] = value + 1
			} else {
				freq[tmpString] = 1
			}

			tmpString = ""
		} else {
			tmpString = tmpString + char
		}

		if index+1 == len(phrase) {
			value, ok := freq[tmpString]
			if ok {
				freq[tmpString] = value + 1
			} else {
				freq[tmpString] = 1
			}
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
