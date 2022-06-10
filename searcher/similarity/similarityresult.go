package similarity

import (
	"os"
)

func Similar(s []string) []string {
	// Each raw set must be a slice of unique string tokens.
	// Use frequency order transformation to replace the string tokens
	// with integers.

	file, err := os.Open("data/word1.txt")
	if err != nil {
		panic(err)
	}
	ss, err := ReadFile("data/dec1.txt")
	if err != nil {
		panic(err)
	}
	setIDs, rawSets, err := ReadFlattenedRawSets(file, true)
	if err != nil {
		panic(err)
	}

	sets, dict := FrequencyOrderTransform(rawSets)
	// Build a search index.
	searchIndex, err := NewSearchIndex(sets,
		"jaccard",
		0.1)
	if err != nil {
		panic(err)
	}
	// Use dictionary to transform a query set.
	querySet := dict.Transform(s)
	// Query the search index.
	searchResults := searchIndex.Query(querySet)
	similarword := make([]string, 0)
	for _, result := range searchResults {
		// X is the index to the original rawSets and sets slices.
		similarword = append(similarword, ss[StrToInt(setIDs[result.X])])

	}

	return similarword
}
