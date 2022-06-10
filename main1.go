package main
/*
import (
	"SearchEngineV2/searcher/utils"

	"fmt"
	"os"
	"strings"

	"github.com/wangbin/jiebago"
)

type Tokenizer struct {
	seg jiebago.Segmenter
}

func main() {
	// Each raw set must be a slice of unique string tokens.
	// Use frequency order transformation to replace the string tokens
	// with integers.

	file, err := os.Open("E:/query/study/Security-Brigade-Rep1/Security-Brigade-Rep1/data/word1.txt")
	if err != nil {
		panic(err)
	}
	ss, err := SetSimilaritySearch.ReadFile("E:/query/study/Security-Brigade-Rep1/Security-Brigade-Rep1/data/dec1.txt")
	if err != nil {
		panic(err)
	}
	setIDs, rawSets, err := SetSimilaritySearch.ReadFlattenedRawSets(file, true)
	if err != nil {
		panic(err)
	}

	sets, dict := SetSimilaritySearch.FrequencyOrderTransform(rawSets)
	// Build a search index.
	searchIndex, err := SetSimilaritySearch.NewSearchIndex(sets,
		"jaccard",
		0.1)
	if err != nil {
		panic(err)
	}
	// Use dictionary to transform a query set.
	var word string
	_, _ = fmt.Scanln(&word)
	s := Cut(word)
	querySet := dict.Transform(s)
	// Query the search index.
	searchResults := searchIndex.Query(querySet)
	for _, result := range searchResults {
		// X is the index to the original rawSets and sets slices.
		fmt.Println(ss[SetSimilaritySearch.StrToInt(setIDs[result.X])])

	}
}

//1662
func Cut(text string) []string {
	tokenizer := &Tokenizer{}
	dictionaryPath := "E:/query/study/Security-Brigade-Rep1/Security-Brigade-Rep1/data/dictionary.txt"
	err := tokenizer.seg.LoadDictionary(dictionaryPath) //不区分大小写
	if err != nil {
		panic(err)
	}
	text = strings.ToLower(text)
	//移除所有的标点符号
	text = utils.RemovePunctuation(text)
	//移除所有的空格
	text = utils.RemoveSpace(text)

	var wordMap = make(map[string]int)

	resultChan := tokenizer.seg.CutForSearch(text, true)
	for {
		w, ok := <-resultChan
		if !ok {
			break
		}
		_, found := wordMap[w]
		if !found {
			//去除重复的词
			wordMap[w] = 1
		}
	}

	var wordsSlice []string
	for k := range wordMap {
		wordsSlice = append(wordsSlice, k)
	}

	return wordsSlice
}
*/