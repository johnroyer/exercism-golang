package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	strLength := len(phrase)
	index := 0
	lastSplitWordIndex := 0
	from := 0
	var tmpString string

	for {
		if isSaperateWord(phrase[index : index+1]) {
			from = lastSplitWordIndex + 1
			tmpString = phrase[from : index-1]

			counter, ok := freq[tmpString]
			if ok {
				freq[tmpString] = counter + 1
			} else {
				freq[tmpString] = 1
			}

			lastSplitWordIndex = index
		} else {
			index++
		}

		if index == strLength {
			tmpString = phrase[lastSplitWordIndex:]
			counter, ok := freq[tmpString]
			if ok {
				freq[tmpString] = counter + 1
			} else {
				freq[tmpString] = 1
			}
			break
		}
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
