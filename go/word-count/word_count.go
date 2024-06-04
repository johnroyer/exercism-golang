package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	strLength := len(phrase)
	index := 0
	tmpString := make([]string, 0)

	for {
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
