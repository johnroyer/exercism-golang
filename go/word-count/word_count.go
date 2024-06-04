package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	strLength := len(phrase)
	index := 0
	tmpString := make([]string, 0)
	lastSplitWordIndex := 0
	from := 0
	length := 0

	for {
		if isSaperateWord(phrase[0:1]) {
			from = lastSplitWordIndex + 1
			length = index - lastSplitWordIndex
			tmpString = append(tmpString, phrase[from:length)
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
