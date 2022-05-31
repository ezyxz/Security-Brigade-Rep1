package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/searcher"
	"SearchEngineV2/searcher/model"
)

type Search struct {
	Container *searcher.Container
}

func NewSearch() *Search {
	return &Search{
		Container: global.Container,
	}
}

// WordCut 分词
func (w *Search) WordSearch(keyword string) *model.SearchResult {
	wordsSlices := w.Container.Tokenizer.Cut(keyword)
	return w.Container.Engine.Query(wordsSlices)
}
