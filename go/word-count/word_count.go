package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	strLength := len(phrase)
	index := 0
	lastSplitWordIndex := 0
	from := 0
	length := 0
	var tmpString string

	for {
		if isSaperateWord(phrase[0:1]) {
			from = lastSplitWordIndex + 1
			length = index - lastSplitWordIndex
			tmpString = phrase[from:length]

			counter, ok := freq[tmpString]
			if ok {
				freq[tmpString] = counter + 1
			} else {
				freq[tmpString] = 1
			}

		} else {
			index++
		}

		if index == strLength {
			break
		}
	}
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
