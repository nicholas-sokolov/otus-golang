package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

type WordMap map[string]int

func Top10(inputStr string) []string {
	freqMap := make(WordMap)
	for _, word := range strings.Fields(inputStr) {
		freqMap[word]++
	}
	if len(freqMap) == 0 {
		return nil
	}

	result := make([]string, 0, len(freqMap))
	for key := range freqMap {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		return freqMap[result[i]] > freqMap[result[j]]
	})

	threshold := 10
	if len(result) < threshold {
		threshold = len(result)
	}
	return result[:threshold]
}
