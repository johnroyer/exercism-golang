package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	index := 1
	lastSplitWordIndex := 0
	var tmpString string

	for {
		if isSaperateWord(phrase[index : index+1]) {
			tmpString = phrase[lastSplitWordIndex+1 : index+1]

			counter, ok := freq[tmpString]
			if ok {
				freq[tmpString] = counter + 1
			} else {
				freq[tmpString] = 1
			}

			lastSplitWordIndex = index
		}

		if index+1 == len(phrase) {
			tmpString = phrase[lastSplitWordIndex:]
			counter, ok := freq[tmpString]
			if ok {
				freq[tmpString] = counter + 1
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
