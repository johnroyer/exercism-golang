package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	scoreMap := make(map[string]int)

	for score, chars := range in {
		for _, char := range chars {
			scoreMap[strings.ToLower(char)] = score
		}
	}

	return scoreMap
}
